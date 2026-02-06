package data

type ProjectPayload struct {
	ID          string `json:id`
	Name        string `json:"name"`
	Description string `json:"description"`
	GithubLink  string `json:"github_link"`
	Image       string `json:"image"`
	BlogLink    string `json:"blog_link"`
	Type        string `json:"type"`
}

type ProjectResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	GithubLink  string `json:"github_link"`
	ImageLink   string `json:"image"`
	BlogLink    string `json:"blog_link"`
	Type        string `json:"type"`
	CreatedAt   string `json:created_at`
	UpdatedAt   string `json:updated_at`
}
