package response

type FileInfoResponse struct {
	FileName string
	FileId string
	DriveId string
	TeamDriveId string
}

func NewFileInfoResponse(fileName string,fileId string,driveId string,teamDriveId string) FileInfoResponse {
	return FileInfoResponse{
		FileName:fileName,
		FileId:fileId,
		DriveId:driveId,
		TeamDriveId:teamDriveId,
	}
}

type TeamDriveInfoResponse struct {
	TeamDriveName string
	TeamDriveId string
	Kind string
}

func NewTeamDriveInfoResponse(teamDriveName string,teamDriveId string, kind string) TeamDriveInfoResponse {
	return TeamDriveInfoResponse{
		TeamDriveName:teamDriveName,
		TeamDriveId:teamDriveId,
		Kind:kind,
	}
}