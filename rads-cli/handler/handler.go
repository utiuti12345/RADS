package handler

import "rads-cli/domain"

type Handler struct {
	ApiHandler ApiHandler
}

func NewHandler(apiHandler ApiHandler) (*Handler, error) {
	return &Handler{
		ApiHandler: apiHandler,
	}, nil
}

type ApiHandler interface {
	Get(path string) (body string,statusCode int,err error)
	GetTeamDriveByName(path string,driveName string) (teamDriveInfo domain.TeamDriveInfo, err error)
	GetRosterFileList(path string,driveId string,dateTime string) (fileInfoList []domain.FileInfo, err error)
	CopyFiles(path string,fileInfoList []domain.FileInfo ,driveIds []string) (err error)
}
