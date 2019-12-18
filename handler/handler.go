package handler

import "RosterAutomaticDeliverySystem/domain"

type Handler struct {
	ApiHandler ApiHandler
}
func NewHandler(apiHandler ApiHandler) (*Handler, error) {
	return &Handler{
		ApiHandler: apiHandler,
	}, nil
}
type ApiHandler interface {
	GetFileList() (fileList []domain.FileInfo, err error)
	TransferDrive(fileName string)(webViewLink string,err error)
	DownloadFile(file *domain.FileInfo)(err error)
}