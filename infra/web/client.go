package web

import (
	"RosterAutomaticDeliverySystem/domain"
	"fmt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

type DriveClient struct {
	HTTPClient   *http.Client
	DriveService drive.Service
}

func NewClient(config *oauth2.Config, token *oauth2.Token) (driveClient DriveClient, err error) {
	ds, err := drive.NewService(context.TODO(), option.WithTokenSource(config.TokenSource(context.TODO(), token)))

	return DriveClient{
		HTTPClient:   config.Client(context.Background(), token),
		DriveService: *ds,
	}, nil
}

func (dr DriveClient) GetFileList(driveId string) (fileList []domain.FileInfo, err error) {
	fl, err := dr.DriveService.Files.List().SupportsAllDrives(true).IncludeItemsFromAllDrives(true).Corpora("drive").DriveId(driveId).Do()
	if err != nil {
		return nil, err
	}
	var fil []domain.FileInfo
	for _, f := range fl.Files {
		//if f.MimeType == "application/vnd.google-apps.spreadsheet" {
		fil = append(fil, domain.NewFileInfo(f.Name, f.Id, f.MimeType, f.DriveId, f.TeamDriveId, f.Size, nil))
	}

	return fil, nil
}

func (dr DriveClient) GetTeamDriveList() (fileList []domain.TeamDriveInfo, err error) {
	tdl, err := dr.DriveService.Teamdrives.List().UseDomainAdminAccess(true).Do()
	if err != nil {
		return nil, err
	}

	var fdil []domain.TeamDriveInfo
	for _, td := range tdl.TeamDrives {
		fdil = append(fdil, domain.NewTeamDriveInfo(td.Name, td.Id, td.Kind))
	}

	return fdil, nil
}

func (dr DriveClient) GetDriveList() (fileList []domain.DriveInfo, err error) {
	tdl, err := dr.DriveService.Drives.List().UseDomainAdminAccess(true).Do()
	if err != nil {
		return nil, err
	}

	var fdil []domain.DriveInfo
	for _, td := range tdl.Drives {
		fdil = append(fdil, domain.NewDriveInfo(td.Name, td.Id, td.Kind))
	}

	return fdil, nil
}

func (dr DriveClient) TransferDrive(distFileInfoList []domain.FileInfo, srcDriveIds []string) (err error) {
	for _, dfi := range distFileInfoList {
		df, err := dr.DriveService.Files.Get(dfi.FileId).SupportsAllDrives(true).Do()
		if err != nil {
			return err
		}
		cf := &drive.File{
			Name:    df.Name,
			Parents: srcDriveIds,
		}

		f, err := dr.DriveService.Files.Copy(df.Id, cf).SupportsAllDrives(true).Do();
		if err != nil {
			return err
		}
		fmt.Println(f.Name)
	}
	return nil
}

func (dr DriveClient) CreateContent(content domain.ContentInfo) (createInfo domain.ContentInfo, err error) {
	cf := &drive.File{
		Name:     content.Name,
		MimeType: content.MineType,
		Parents:  []string{content.DriveId},
	}

	f, err := dr.DriveService.Files.Create(cf).SupportsAllDrives(true).Do()
	if err != nil {
		return createInfo, err
	}

	return domain.NewContentInfo(f.Name, f.Id, f.MimeType, f.DriveId, f.TeamDriveId, f.Size, nil), nil
}

func (dr DriveClient) DownloadFile(file *domain.FileInfo) (err error) {
	srv, err := drive.New(dr.HTTPClient)
	if err != nil {
		return err
	}

	r, err := srv.Files.Export(file.FileId, file.MineType).Download()
	b, err := ioutil.ReadAll(r.Body)
	file.Data = b

	return nil
}
