package web

import (
	"RosterAutomaticDeliverySystem/domain"
	"encoding/json"
	"fmt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

type DriveClient struct {
	HTTPClient    *http.Client
	DriveService drive.Service
}

func NewClient(config *oauth2.Config) (driveClient DriveClient,err error) {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}

	ds, err := drive.NewService(context.TODO(), option.WithTokenSource(config.TokenSource(context.TODO(), tok)))

	return DriveClient{
		HTTPClient:config.Client(context.Background(), tok),
		DriveService:*ds,
	}, nil
}

func (dr DriveClient) GetFileList(driveId string)(fileList []domain.FileInfo, err error) {
	fl, err := dr.DriveService.Files.List().SupportsAllDrives(true).IncludeItemsFromAllDrives(true).Corpora("drive").DriveId(driveId).Do()
	if err != nil {
		return nil,err
	}
	var fil []domain.FileInfo
	for _,f := range fl.Files{
		//if f.MimeType == "application/vnd.google-apps.spreadsheet" {
		fil = append(fil,domain.NewFileInfo(f.Name,f.Id,f.MimeType,f.DriveId,f.TeamDriveId,f.Size,nil))
	}

	return fil, nil
}

func (dr DriveClient) GetTeamDriveList()(fileList []domain.TeamDriveInfo, err error) {
	tdl,err := dr.DriveService.Teamdrives.List().UseDomainAdminAccess(true).Do()
	if err != nil {
		return nil,err
	}

	var fdil []domain.TeamDriveInfo
	for _,td := range tdl.TeamDrives{
		fdil = append(fdil,domain.NewTeamDriveInfo(td.Name,td.Id,td.Kind))
	}

	return fdil, nil
}

func (dr DriveClient) TransferDrive(distFileInfoList []domain.FileInfo, srcDriveIds []string) (err error){
	for _ , dfi := range distFileInfoList{
		df, err := dr.DriveService.Files.Get(dfi.FileId).SupportsAllDrives(true).Do()
		if err != nil {
			return err
		}
		cf := &drive.File{
			Name:"copy" + df.Name,
			Parents:srcDriveIds,
		}
		f, err := dr.DriveService.Files.Copy(df.Id,cf).SupportsAllDrives(true).Do(); if err != nil {
			return err
		}
		fmt.Println(f.Name)
	}
	return err
}

func (dr DriveClient) DownloadFile(file *domain.FileInfo)(err error){
	srv, err := drive.New(dr.HTTPClient)
	if err != nil {
		return err
	}

	r,err := srv.Files.Export(file.FileId,file.MineType).Download()
	b, err := ioutil.ReadAll(r.Body)
	file.Data = b

	return nil
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.TODO(), "4/uwEXj2HIc1As0KCpmzhiPYbVsWIYTDG0EKnE77MCAaYphvYzCzehGGI")
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}