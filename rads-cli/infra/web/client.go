package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"rads-cli/domain"
	"rads-cli/handler/request"
	"rads-cli/handler/response"
	"regexp"
)

type Client struct {
	URL           string
	DefaultHeader http.Header
	HTTPClient    *http.Client
}
func NewClient(host string) (client *Client, err error) {
	u, err := url.ParseRequestURI(host)
	if err != nil {
		return nil, err
	}
	return &Client{
		URL:           u.String(),
		DefaultHeader: make(http.Header),
		HTTPClient:    &http.Client{},
	}, nil
}
func (c Client) Get(path string) (body string, statusCode int, err error) {
	resp, err := c.HTTPClient.Get(fmt.Sprintf("%s%s", c.URL, path))
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, statusCode, err
	}
	defer resp.Body.Close()
	return string(b), resp.StatusCode, err
}

func (c Client) GetTeamDriveByName(path string,driveName string) (teamDriveInfo domain.TeamDriveInfo, err error){
	// /teamDriveList
	resp, err := c.HTTPClient.Get(fmt.Sprintf("%s%s", c.URL, path))
	if err != nil {
		return teamDriveInfo, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	var res []response.TeamDriveInfoResponse
	if err = json.Unmarshal(b, &res); err != nil {
		return teamDriveInfo, err
	}

	for _, v := range res {
		if v.TeamDriveName == driveName {
			return domain.NewTeamDriveInfo(v.TeamDriveName,v.TeamDriveId,v.Kind) , nil
		}
	}

	return teamDriveInfo,err
}

func (c Client) GetRosterFileList(path string,driveId string,dateTime string) (fileInfoList []domain.FileInfo, err error){
	// /files/:driveId
	resp, err := c.HTTPClient.Get(fmt.Sprintf("%s%s/%s", c.URL, path,driveId))
	if err != nil {
		return fileInfoList, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	var res []response.FileInfoResponse
	if err = json.Unmarshal(b, &res); err != nil {
		return fileInfoList, err
	}

	r := regexp.MustCompile(dateTime)
	var fil []domain.FileInfo
	for _, v := range res {
		if(v.MineType == "application/vnd.google-apps.spreadsheet" ||
			v.MineType == "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" ||
			v.MineType == "application/vnd.ms-excel"){
			if r.MatchString(v.FileName){
				fil = append(fil,domain.NewFileInfo(v.FileName,v.FileId,v.MineType,v.DriveId,v.TeamDriveId,0,nil))
			}
		}
	}

	return fil,err
}

func (c Client) CreateFolder(path string,content domain.ContentInfo ,driveIds []string) (createContent domain.ContentInfo,err error) {
	cr := request.NewCreateRequest(content.Name,content.MimeType,driveIds)

	body,err := json.Marshal(cr)

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s%s", c.URL, path),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return createContent,err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return createContent,err
	}

	if resp.StatusCode != http.StatusOK {
		return createContent,errors.New(resp.Status)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	var res response.ContentInfoResponse
	if err = json.Unmarshal(b, &res); err != nil {
		return createContent, err
	}

	return domain.NewContentInfo(res.Name,res.Id,res.MimeType,res.DriveId,res.DriveId,0,nil),nil
}

func (c Client) CopyFiles(path string,fileInfoList []domain.FileInfo ,driveIds []string) (err error) {
	// /copy
	var copyreq request.CopyFileRequest
	var fileInfoListReq []request.FileInfoRequest

	for _,v := range fileInfoList{
		fileInfoListReq = append(fileInfoListReq,request.NewFileInfoRequest(v.DriveId,v.FileId))
	}

	copyreq.DistFiles = fileInfoListReq
	copyreq.SrcDriveIds = driveIds

	body,err := json.Marshal(copyreq)

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s%s", c.URL, path),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return nil
}