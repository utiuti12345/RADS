package handler

import "fmt"

func (h Handler) CopyFile(driveName string, dateTime string) (err error) {
	td, err := h.ApiHandler.GetTeamDriveByName("/teamDriveList", driveName)
	if err != nil {
		return err
	}

	fil, err := h.ApiHandler.GetRosterFileList("/files", td.TeamDriveId, dateTime)
	if err != nil {
		return err
	}

	fmt.Print(fil)
	//driveIds := []string{"0AM9O1T03S7baUk9PVA","1eZr6En4CjDGFhT1fkjbrwXNT9Q9aXKu5"}
	//driveIds := []string{"12bjAIa8We6B2RP0oMmXA3NkO4Sy866Qg"}
	driveIds := []string{"1eZr6En4CjDGFhT1fkjbrwXNT9Q9aXKu5"}
	err = h.ApiHandler.CopyFiles("/copy",fil,driveIds)
	if err != nil {
		return err
	}

	return nil
}
