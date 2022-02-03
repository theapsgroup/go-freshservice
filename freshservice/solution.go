package freshservice

const (
	solutionCategoriesUrl      = "solutions/categories"
	solutionCategoryIdUrl      = "solutions/categories/%d"
	solutionFoldersUrl         = "solutions/folders"
	solutionFolderIdUrl        = "solutions/folders/%d"
	solutionArticlesUrl        = "solutions/articles"
	solutionArticleIdUrl       = "solutions/articles/%d"
	solutionArticleApprovalUrl = "solutions/articles/%d/send_for_approval"
)

// SolutionService API Docs: https://api.freshservice.com/#solution-category https://api.freshservice.com/#solution-folder https://api.freshservice.com/#solution-article
type SolutionService struct {
	client *Client
}
