package handler

import (
	"rads-cli/domain"
)

const mimeTypeFolder string = "application/vnd.google-apps.folder"

func (h Handler) CopyFile(driveName string, year string,month string) (err error) {
	td, err := h.ApiHandler.GetTeamDriveByName("/teamDriveList", driveName)
	if err != nil {
		return err
	}

	fil, err := h.ApiHandler.GetRosterFileList("/files", td.TeamDriveId, year + month)
	if err != nil {
		return err
	}

	// 勤務表のフォルダを作成(あれば取得)
	cf,err := h.ApiHandler.CreateFolder("/create",domain.NewContentInfo("勤務表","",mimeTypeFolder,"","",0,nil),nil)
	if err != nil {
		return err
	}

	// 年度のフォルダを作成(あれば取得)
	cf,err = h.ApiHandler.CreateFolder("/create",domain.NewContentInfo(year+"年","",mimeTypeFolder,"","",0,nil),[]string{cf.Id})
	if err != nil {
		return err
	}
	// 月のフォルダを作成(あれば取得)
	cf,err = h.ApiHandler.CreateFolder("/create",domain.NewContentInfo(month+"月分","",mimeTypeFolder,"","",0,nil),[]string{cf.Id})
	if err != nil {
		return err
	}
	//driveIds := []string{"0AM9O1T03S7baUk9PVA","1eZr6En4CjDGFhT1fkjbrwXNT9Q9aXKu5"}
	//driveIds := []string{"12bjAIa8We6B2RP0oMmXA3NkO4Sy866Qg"}
	//driveIds := []string{"1eZr6En4CjDGFhT1fkjbrwXNT9Q9aXKu5"}
	driveIds := []string{cf.Id}
	err = h.ApiHandler.CopyFiles("/copy",fil,driveIds)
	if err != nil {
		return err
	}

	return nil
}
