package request

type CreateRequest struct {
	Name string `json:"Name"`
	MimeType string `json:"MimeType"`
	Parents []string `json:"Parents"`
}