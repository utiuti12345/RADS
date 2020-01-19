package domain

type FileInfo struct {
	FileName string
	FileId string
	MineType string
	DriveId string
	TeamDriveId string
	Size int64
	Data []byte
}

type TeamDriveInfo struct {
	TeamDriveName string
	TeamDriveId string
	Kind string
}

type ContentInfo struct {
	Name string
	Id string
	MimeType string
	DriveId string
	TeamDriveId string
	Size int64
	Data []byte
}

func NewFileInfo(fileName string,fileId string,mineType string,driveId string,teamDriveId string,size int64,data []byte) FileInfo{
	return FileInfo{
		FileName:fileName,
		FileId:fileId,
		MineType:mineType,
		DriveId:driveId,
		TeamDriveId:teamDriveId,
		Size:size,
		Data:data,
	}
}

func NewTeamDriveInfo(teamDriveName string,teamDriveId string, kind string) TeamDriveInfo {
	return TeamDriveInfo{
		TeamDriveName:teamDriveName,
		TeamDriveId:teamDriveId,
		Kind:kind,
	}
}

func NewContentInfo(name string,id string,mimeType string,driveId string,teamDriveId string,size int64,data []byte) ContentInfo{
	return ContentInfo{
		Name:name,
		Id:id,
		MimeType:mimeType,
		DriveId:driveId,
		TeamDriveId:teamDriveId,
		Size:size,
		Data:data,
	}
}