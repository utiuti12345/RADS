package web

import (
	"RosterAutomaticDeliverySystem/handler"
	"github.com/labstack/echo"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"

	"google.golang.org/api/drive/v3"
)

func LoadRouter() *echo.Echo {
	e := echo.New()
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.jsn.
	config, err := google.ConfigFromJSON(b, drive.DriveScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client,err := NewClient(config)
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	h, _ := handler.NewHandler(client)
	e.GET("/", func(ctx echo.Context) error { return h.GetFileList(ctx) })
	e.GET("/teamDriveList", func(ctx echo.Context) error { return h.GetTeamDriveList(ctx) })


	return e
}