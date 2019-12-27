package request

type CopyFileRequest struct {
	DistFiles []FileInfoRequest `json:"DistFiles"`
	SrcDriveIds []string `json:"SrcDriveIds"`
}

type FileInfoRequest struct {
	DriveId string `json:"DriveId"`
	FileId string `json:"FileId"`
}