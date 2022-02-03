package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

// SolutionArticles contains Collection an array of SolutionArticle
type SolutionArticles struct {
	Collection []SolutionArticle `json:"articles"`
}

// solutionArticleWrapper contains Details of one SolutionArticle
type solutionArticleWrapper struct {
	Details SolutionArticle `json:"article"`
}

// SolutionArticle represents a FreshService SolutionArticle
type SolutionArticle struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Position       int       `json:"position"`
	ArticleType    int       `json:"article_type"`
	FolderID       int       `json:"folder_id"`
	CategoryID     int       `json:"category_id"`
	Status         int       `json:"status"`
	ApprovalStatus int       `json:"approval_status"`
	ThumbsUp       int       `json:"thumbs_up"`
	ThumbsDown     int       `json:"thumbs_down"`
	AgentID        int       `json:"agent_id"`
	Views          int       `json:"views"`
	Tags           []string  `json:"tags"`
	Keywords       []string  `json:"keywords"`
	Url            string    `json:"url"`
	ReviewDate     time.Time `json:"review_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	// Attachments
}

// CreateSolutionArticleModel is a data struct for creating a new SolutionArticle
type CreateSolutionArticleModel struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ArticleType int       `json:"article_type"`
	FolderID    int       `json:"folder_id"`
	Status      int       `json:"status"`
	Tags        []string  `json:"tags"`
	Keywords    []string  `json:"keywords"`
	ReviewDate  time.Time `json:"review_date"`
}

// UpdateSolutionArticleModel is a data struct for updating a SolutionArticle
type UpdateSolutionArticleModel struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	ArticleType int       `json:"article_type,omitempty"`
	FolderID    int       `json:"folder_id,omitempty"`
	Status      int       `json:"status,omitempty"`
	Tags        []string  `json:"tags,omitempty"`
	Keywords    []string  `json:"keywords,omitempty"`
	ReviewDate  time.Time `json:"review_date,omitempty"`
}

// ListSolutionArticlesOptions represents filters/pagination for SolutionArticles
type ListSolutionArticlesOptions struct {
	ListOptions
	FolderID int `json:"folder_id"`
}

// GetSolutionArticle will return a SolutionArticle by id
func (s *SolutionService) GetSolutionArticle(id int) (*SolutionArticle, *http.Response, error) {
	o := new(solutionArticleWrapper)
	res, err := s.client.Get(fmt.Sprintf(solutionArticleIdUrl, id), &o)
	return &o.Details, res, err
}

// ListSolutionArticles will return paginated/filtered SolutionArticles using ListSolutionArticlesOptions
func (s *SolutionService) ListSolutionArticles(opt *ListSolutionArticlesOptions) (*SolutionArticles, *http.Response, error) {
	o := new(SolutionArticles)
	res, err := s.client.List(solutionArticlesUrl, opt, &o)
	return o, res, err
}

// CreateSolutionArticle will create and return a new SolutionArticle based on CreateSolutionArticleModel
func (s *SolutionService) CreateSolutionArticle(solutionArticle *CreateSolutionArticleModel) (*SolutionArticle, *http.Response, error) {
	o := new(solutionArticleWrapper)
	res, err := s.client.Post(solutionArticlesUrl, solutionArticle, &o)
	return &o.Details, res, err
}

// UpdateSolutionArticle will update and return a SolutionArticle matching id based UpdateSolutionArticleModel
func (s *SolutionService) UpdateSolutionArticle(id int, solutionArticle *UpdateSolutionArticleModel) (*SolutionArticle, *http.Response, error) {
	o := new(solutionArticleWrapper)
	res, err := s.client.Put(fmt.Sprintf(solutionArticleIdUrl, id), solutionArticle, &o)
	return &o.Details, res, err
}

// DeleteSolutionArticle will completely remove a SolutionArticle from FreshService matching id
func (s *SolutionService) DeleteSolutionArticle(id int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(solutionArticleIdUrl, id))
	return success, res, err
}

// SendSolutionArticleForApproval sends the SolutionArticle matching id for approval
func (s *SolutionService) SendSolutionArticleForApproval(id int) (*SolutionArticle, *http.Response, error) {
	o := new(solutionArticleWrapper)
	res, err := s.client.Put(fmt.Sprintf(solutionArticleApprovalUrl, id), nil, &o)
	return &o.Details, res, err
}
