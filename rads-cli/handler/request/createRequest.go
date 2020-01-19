package request

type CreateRequest struct {
	Name string `json:"Name"`
	MimeType string `json:"MimeType"`
	Parents []string `json:"Parents"`
}

func NewCreateRequest(name string,mimeType string,parents []string)CreateRequest{
	return CreateRequest{
		Name:     name,
		MimeType: mimeType,
		Parents:  parents,
	}
}