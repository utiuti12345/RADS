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