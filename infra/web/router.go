package web

import (
	"RosterAutomaticDeliverySystem/config"
	"RosterAutomaticDeliverySystem/handler"
	"github.com/labstack/echo"
)

func LoadRouter(googleConfig config.GoogleConfig) *echo.Echo {
	e := echo.New()

	client, _ := NewClient(googleConfig.Config,googleConfig.Token)
	h, _ := handler.NewHandler(client)
	e.GET("/files/:driveId", func(ctx echo.Context) error { return h.GetFileList(ctx) })
	e.GET("/teamDriveList", func(ctx echo.Context) error { return h.GetTeamDriveList(ctx) })
	e.POST("/copy", func(ctx echo.Context) error { return h.CopyFile(ctx) })

	return e
}