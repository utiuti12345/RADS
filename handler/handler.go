package handler

import "RosterAutomaticDeliverySystem/domain"

type Handler struct {
	ApiHandler ApiHandler
	Slack      Slack
}

func NewHandler(apiHandler ApiHandler, slack Slack) (*Handler, error) {
	return &Handler{
		ApiHandler: apiHandler,
		Slack:      slack,
	}, nil
}

type ApiHandler interface {
	GetFileList(driveId string) (fileList []domain.FileInfo, err error)
	TransferDrive(distFileInfoList []domain.FileInfo, srcDriveIds []string) (err error)
	DownloadFile(file *domain.FileInfo) (err error)
	GetTeamDriveList() (teamDriveList []domain.TeamDriveInfo, err error)
	GetDriveList() (driveList []domain.DriveInfo, err error)
	CreateContent(content domain.ContentInfo) (createInfo domain.ContentInfo, err error)
}

type Slack interface {
	PostMessage(title string,message string)(errs []error)
}
