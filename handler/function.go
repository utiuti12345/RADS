package handler

import (
	"RosterAutomaticDeliverySystem/domain"
	"RosterAutomaticDeliverySystem/handler/request"
	"RosterAutomaticDeliverySystem/handler/response"
	"net/http"
)

func (h Handler) GetFileList(ctx Context) (err error) {
	di := ctx.Param("driveId")
	fl, err := h.ApiHandler.GetFileList(di)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	var fir []response.FileInfoResponse
	for _,v := range fl{
		fir = append(fir,response.NewFileInfoResponse(v.FileName,v.FileId,v.MineType,v.DriveId,v.TeamDriveId))
	}

	return ctx.JSON(http.StatusOK, fir)
}

func (h Handler) GetTeamDriveList(ctx Context) (err error) {
	tdl, err := h.ApiHandler.GetTeamDriveList()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	var tdir []response.TeamDriveInfoResponse
	for _,v := range tdl{
		tdir = append(tdir,response.NewTeamDriveInfoResponse(v.TeamDriveName,v.TeamDriveId,v.Kind))
	}

	return ctx.JSON(http.StatusOK, tdir)
}

func (h Handler) CopyFile(ctx Context) (err error) {
	cfr := new(request.CopyFileRequest)
	if err := ctx.Bind(cfr); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	var fil []domain.FileInfo
	for _,cf := range cfr.DistFiles {
		fil = append(fil,domain.NewFileInfo("",cf.FileId,"",cf.DriveId,"",0,nil))
	}

	if err := h.ApiHandler.TransferDrive(fil,cfr.SrcDriveIds); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}