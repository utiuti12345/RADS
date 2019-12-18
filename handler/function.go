package handler

import (
	"RosterAutomaticDeliverySystem/handler/response"
	"net/http"
)

func (h Handler) GetFileList(ctx Context) (err error) {
	fl, err := h.ApiHandler.GetFileList()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	var fir []response.FileInfoResponse
	for _,v := range fl{
		fir = append(fir,response.NewFileInfoResponse(v.FileName,v.FileId,v.DriveId,v.TeamDriveId))
	}

	return ctx.JSON(http.StatusOK, fir)
}

func (h Handler) CopyFile(ctx Context) (err error) {

	return nil
}