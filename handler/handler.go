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
	GetFileList(driveId string) (fileList []domain.FileInfo, err error)
	TransferDrive(distFileInfoList []domain.FileInfo, srcDriveIds []string)(err error)
	DownloadFile(file *domain.FileInfo)(err error)
	GetTeamDriveList() (teamDriveList []domain.TeamDriveInfo, err error)
}