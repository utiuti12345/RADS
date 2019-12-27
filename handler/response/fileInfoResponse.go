package response

type FileInfoResponse struct {
	FileName string
	FileId string
	MineType string
	DriveId string
	TeamDriveId string
}

func NewFileInfoResponse(fileName string,fileId string,mineType string,driveId string,teamDriveId string) FileInfoResponse {
	return FileInfoResponse{
		FileName:fileName,
		FileId:fileId,
		MineType:mineType,
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