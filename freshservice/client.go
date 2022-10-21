package freshservice

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	cleanHttp "github.com/hashicorp/go-cleanhttp"
	retryHttp "github.com/hashicorp/go-retryablehttp"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	userAgent = "go-freshservice"
)

type Client struct {
	client    *retryHttp.Client
	baseUrl   *url.URL
	token     string
	UserAgent string
	// Add services
	Agents                 *AgentService
	Announcements          *AnnouncementService
	Assets                 *AssetService
	BusinessHours          *BusinessHoursService
	Changes                *ChangeService
	Contracts              *ContractService
	Departments            *DepartmentService
	Locations              *LocationService
	Problems               *ProblemService
	Products               *ProductService
	PurchaseOrders         *PurchaseOrderService
	Releases               *ReleaseService
	Requesters             *RequesterService
	Services               *ServiceCatalogService
	ServiceLevelAgreements *SLAPoliciesService
	Software               *SoftwareService
	Solutions              *SolutionService
	Tickets                *TicketService
	Vendors                *VendorService
	CustomObject           *CustomObjectService
}

// NewClient generates a new API client, requires the subdomain of your FreshService instance as well as an API key
// ctx is optional and will default to context.Background() if nil is passed.
func NewClient(ctx context.Context, subDomain string, apiKey string) (*Client, error) {

	if ctx == nil {
		ctx = context.Background()
	}

	if subDomain == "" {
		return nil, fmt.Errorf("sub-domain was not provided but is required")
	}

	if apiKey == "" {
		return nil, fmt.Errorf("api key was not provided but is required")
	}

	baseUrl, err := buildBaseUrl(subDomain)
	if err != nil {
		return nil, err
	}

	fs := &Client{
		baseUrl:   baseUrl,
		token:     apiKey,
		UserAgent: userAgent,
	}

	fs.client = &retryHttp.Client{
		Backoff:      fs.setBackoff,
		CheckRetry:   fs.canRetry,
		ErrorHandler: retryHttp.PassthroughErrorHandler,
		HTTPClient:   cleanHttp.DefaultPooledClient(),
		RetryMax:     5,
		RetryWaitMin: 500 * time.Millisecond,
		RetryWaitMax: time.Second,
	}

	// Add services
	fs.Agents = &AgentService{client: fs}
	fs.Announcements = &AnnouncementService{client: fs}
	fs.Assets = &AssetService{client: fs}
	fs.BusinessHours = &BusinessHoursService{client: fs}
	fs.Changes = &ChangeService{client: fs}
	fs.Contracts = &ContractService{client: fs}
	fs.Departments = &DepartmentService{client: fs}
	fs.Locations = &LocationService{client: fs}
	fs.Problems = &ProblemService{client: fs}
	fs.Products = &ProductService{client: fs}
	fs.PurchaseOrders = &PurchaseOrderService{client: fs}
	fs.Releases = &ReleaseService{client: fs}
	fs.Requesters = &RequesterService{client: fs}
	fs.Services = &ServiceCatalogService{client: fs}
	fs.ServiceLevelAgreements = &SLAPoliciesService{client: fs}
	fs.Software = &SoftwareService{client: fs}
	fs.Solutions = &SolutionService{client: fs}
	fs.Tickets = &TicketService{client: fs}
	fs.Vendors = &VendorService{client: fs}
	fs.CustomObject = &CustomObjectService{client: fs}

	return fs, nil
}

// canRetry determines if retrying should happen (rate limit, temporary server issue)
func (c *Client) canRetry(ctx context.Context, res *http.Response, err error) (bool, error) {
	if err != nil {
		return false, err
	}

	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	if res.StatusCode == 429 || res.StatusCode >= 500 {
		return true, nil
	}

	return false, nil
}

// setBackoff determines backoff for re-attempts due to rate-limits (take time from header) or transient failures.
func (c *Client) setBackoff(min time.Duration, max time.Duration, attemptNum int, res *http.Response) time.Duration {
	if res != nil && res.StatusCode == 429 {
		if secs := res.Header.Get("Retry-After"); secs != "" {
			delay, err := strconv.Atoi(secs)
			if err != nil {
				delay = 60
			}
			return time.Duration(delay) * time.Second
		}
	}

	return retryHttp.LinearJitterBackoff(min, max, attemptNum, res)
}

func (c *Client) buildRequest(method string, path string, opt interface{}) (*retryHttp.Request, error) {
	dest := *c.baseUrl

	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, fmt.Errorf("error in path: %s\n%v", path, err)
	}

	dest.RawPath = c.baseUrl.Path + path
	dest.Path = c.baseUrl.Path + unescaped

	var content interface{}
	setContentType := false

	switch method {
	case http.MethodGet:
		if opt != nil {
			q, err := query.Values(opt)
			if err != nil {
				return nil, fmt.Errorf("error creating query string for request: %v", err)
			}
			dest.RawQuery = q.Encode()
		}
	case http.MethodPost, http.MethodPut:
		setContentType = true
		if opt != nil {
			content, err = json.Marshal(opt)
			if err != nil {
				return nil, fmt.Errorf("error formatting body for request: %v", err)
			}
		}
	case http.MethodDelete:
		break // no special behaviour required
	default:
		return nil, fmt.Errorf("method %s is not supported", method)
	}

	req, err := retryHttp.NewRequest(method, dest.String(), content)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %v", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0, post-check=0, pre-check=0")
	req.Header.Set("Strict-Transport-Security", "max-age=31536000 ; includeSubDomains")
	req.Header.Set("User-Agent", c.UserAgent)
	if setContentType {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (c *Client) sendRequest(req *retryHttp.Request, o interface{}) (*http.Response, error) {
	fmt.Printf("%s\n", req.URL.String())
	req.SetBasicAuth(c.token, "X")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer res.Body.Close()

	success, msg := isSuccessful(res)
	if !success {
		return res, fmt.Errorf(msg)
	}

	if o != nil {
		if w, ok := o.(io.Writer); ok {
			_, err = io.Copy(w, res.Body)
		} else {
			err = json.NewDecoder(res.Body).Decode(o)
		}
	}

	return res, err
}

func (c *Client) Get(path string, out interface{}) (*http.Response, error) {
	req, err := c.buildRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating GET request for path '%s': %v", path, err)
	}

	res, err := c.sendRequest(req, &out)
	if b, s := isSuccessful(res); !b {
		return res, fmt.Errorf("%s: %v", s, err)
	}

	return res, nil
}

func (c *Client) List(path string, opt interface{}, out interface{}) (*http.Response, error) {
	req, err := c.buildRequest(http.MethodGet, path, opt)
	if err != nil {
		return nil, fmt.Errorf("error creating GET request for path '%s': %v", path, err)
	}

	res, err := c.sendRequest(req, &out)
	if b, s := isSuccessful(res); !b {
		return res, fmt.Errorf("%s: %v", s, err)
	}

	return res, nil
}

func (c *Client) Post(path string, body interface{}, out interface{}) (*http.Response, error) {
	req, err := c.buildRequest(http.MethodPost, path, body)
	if err != nil {
		return nil, fmt.Errorf("error creating POST request for path '%s': %v", path, err)
	}

	res, err := c.sendRequest(req, &out)
	if b, s := isSuccessful(res); !b {
		return res, fmt.Errorf("%s: %v", s, err)
	}

	return res, nil
}

func (c *Client) Put(path string, body interface{}, out interface{}) (*http.Response, error) {
	req, err := c.buildRequest(http.MethodPut, path, body)
	if err != nil {
		return nil, fmt.Errorf("error creating PUT request for path '%s': %v", path, err)
	}

	res, err := c.sendRequest(req, &out)
	if b, s := isSuccessful(res); !b {
		return res, fmt.Errorf("%s: %v", s, err)
	}

	return res, nil
}

func (c *Client) Delete(path string) (bool, *http.Response, error) {
	req, err := c.buildRequest(http.MethodDelete, path, nil)
	if err != nil {
		return false, nil, fmt.Errorf("error creating DELETE request for path '%s': %v", path, err)
	}

	res, err := c.sendRequest(req, nil)
	if b, s := isSuccessful(res); !b {
		return false, res, fmt.Errorf("%s: %v", s, err)
	}

	return true, res, nil
}

// buildBaseUrl sets the baseUrl based on provided subdomain
func buildBaseUrl(subDomain string) (*url.URL, error) {
	return url.Parse(fmt.Sprintf("https://%s.freshservice.com/api/v2/", subDomain))
}

// isSuccessful is a function to determine a http call executed successfully
func isSuccessful(res *http.Response) (bool, string) {
	if res.StatusCode >= 200 && res.StatusCode <= 204 {
		return true, ""
	}

	return false, fmt.Sprintf("request returned non-success status %d: %s", res.StatusCode, res.Status)
}
