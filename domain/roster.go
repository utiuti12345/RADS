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

type DriveInfo struct {
	DriveName string
	DriveId string
	Kind string
}

type ContentInfo struct {
	Name string
	Id string
	MineType string
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

func NewTeamDriveInfo(driveName string,driveId string, kind string) TeamDriveInfo {
	return TeamDriveInfo{
		TeamDriveName:driveName,
		TeamDriveId:driveId,
		Kind:kind,
	}
}

func NewDriveInfo(teamDriveName string,teamDriveId string, kind string) DriveInfo {
	return DriveInfo{
		DriveName:teamDriveName,
		DriveId:teamDriveId,
		Kind:kind,
	}
}

func NewContentInfo(name string,id string,mineType string,driveId string,teamDriveId string,size int64,data []byte) ContentInfo{
	return ContentInfo{
		Name:name,
		Id:id,
		MineType:mineType,
		DriveId:driveId,
		TeamDriveId:teamDriveId,
		Size:size,
		Data:data,
	}
}