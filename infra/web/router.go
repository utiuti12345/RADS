package web

import (
	"RosterAutomaticDeliverySystem/config"
	"RosterAutomaticDeliverySystem/handler"
	"RosterAutomaticDeliverySystem/infra/slack"
	"github.com/labstack/echo"
)

func LoadRouter(config config.Config) *echo.Echo {
	e := echo.New()

	s := slack.NewSlack(config.SlackConfig.WebHookConfig.WebhookUrl,config.SlackConfig.WebHookConfig.Channel,config.SlackConfig.WebHookConfig.UserName)
	c, _ := NewClient(config.GoogleConfig.Config,config.GoogleConfig.Token)
	h, _ := handler.NewHandler(c,s)
	e.GET("/healthCheck", func(ctx echo.Context) error { return h.HealthCheck(ctx) })
	e.GET("/files/:driveId", func(ctx echo.Context) error { return h.GetFileList(ctx) })
	e.GET("/teamDriveList", func(ctx echo.Context) error { return h.GetTeamDriveList(ctx) })
	e.GET("/DriveList", func(ctx echo.Context) error { return h.GetDriveList(ctx) })
	e.POST("/copy", func(ctx echo.Context) error { return h.CopyFile(ctx) })
	e.POST("/create", func(ctx echo.Context) error { return h.Create(ctx) })
	e.POST("/postSlack", func(ctx echo.Context) error { return h.PostSlack(ctx) })

	return e
}