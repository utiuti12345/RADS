package handler

import (
	"RosterAutomaticDeliverySystem/domain"
	"RosterAutomaticDeliverySystem/handler/request"
	"RosterAutomaticDeliverySystem/handler/response"
	"net/http"
)

func (h Handler) HealthCheck(ctx Context) (err error) {
	return ctx.String(http.StatusOK, "HealthCheck OK")
}

func (h Handler) GetFileList(ctx Context) (err error) {
	di := ctx.Param("driveId")
	fl, err := h.ApiHandler.GetFileList(di)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	var fir []response.FileInfoResponse
	for _,v := range fl{
		fir = append(fir,response.NewFileInfoResponse(v.FileName,v.FileId,v.MimeType,v.DriveId,v.TeamDriveId))
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

func (h Handler) GetDriveList(ctx Context) (err error) {
	tdl, err := h.ApiHandler.GetDriveList()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	var tdir []response.DriveInfoResponse
	for _,v := range tdl{
		tdir = append(tdir,response.NewDriveInfoResponse(v.DriveName,v.DriveId,v.Kind))
	}

	return ctx.JSON(http.StatusOK, tdir)
}

func (h Handler) CopyFile(ctx Context) (err error) {
	cfr := new(request.CopyFileRequest)
	if err := ctx.Bind(cfr); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	var fil []domain.FileInfo
	for _,cf := range cfr.DistFiles {
		fil = append(fil,domain.NewFileInfo("",cf.FileId,"",cf.DriveId,"",0,nil))
	}

	if err := h.ApiHandler.TransferDrive(fil,cfr.SrcDriveIds); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, "OK")
}

func (h Handler) Create(ctx Context) (err error) {
	cr := new(request.CreateRequest)
	if err := ctx.Bind(cr); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	parents := ""
	if len(cr.Parents) != 0 {
		parents = cr.Parents[0]
	}

	cc, err := h.ApiHandler.GetContent(domain.NewContentInfo(cr.Name,"",cr.MimeType,parents,"",0,nil))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	if cc.Id == "" {
		cc, err = h.ApiHandler.CreateContent(domain.NewContentInfo(cr.Name,"",cr.MimeType,parents,"",0,nil))
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}
	}

	return ctx.JSON(http.StatusOK, response.NewContentInfoResponse(cc.Name,cc.Id,cc.MimeType,cc.DriveId,cc.TeamDriveId))
}

func (h Handler) PostSlack(ctx Context) (err error){
	psr := new(request.PostSlackRequest)
	if err := ctx.Bind(psr); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	if err := h.Slack.PostMessage(psr.Title,psr.Message); err != nil{
		return ctx.String(http.StatusInternalServerError, err[0].Error())
	}
	return nil
}