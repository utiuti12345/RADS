package request

type CopyFileRequest struct {
	DistFiles []FileInfoRequest `json:"DistFiles"`
	SrcDriveIds []string `json:"SrcDriveIds"`
}

type FileInfoRequest struct {
	DriveId string `json:"DriveId"`
	FileId string `json:"FileId"`
}

func NewFileInfoRequest(driveId string,fileId string)(fileInfoRequest FileInfoRequest){
	return FileInfoRequest{
		DriveId: driveId,
		FileId:  fileId,
	}
}