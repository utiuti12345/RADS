package config

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v2"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type Config struct {
	GoogleConfig GoogleConfig
	SlackConfig  SlackConfig
}

type GoogleConfig struct {
	Config *oauth2.Config
	Token  *oauth2.Token
}

type SlackConfig struct {
	WebHookConfig WebHookConfig `toml:"slack"`
}
type WebHookConfig struct {
	WebhookUrl string `toml:"webhookurl"`
	Channel    string `toml:"channel"`
	UserName   string `toml:"username"`
}

const CONFIG_TOML = "config.toml"
const CREDENTIALS_JSON = "credentials.json"
const TOKEN_JSON = "token.json"

func NewConfig(googleConfig GoogleConfig, slackConfig SlackConfig)Config{
	return Config{
		GoogleConfig: googleConfig,
		SlackConfig:  slackConfig,
	}
}

func NewGoogleConfig(credentialJsonPath string, tokenJsonPath string) GoogleConfig {
	b, err := ioutil.ReadFile(path.Join(credentialJsonPath, CREDENTIALS_JSON))
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.jsn.
	config, err := google.ConfigFromJSON(b, drive.DriveScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	tokFile := path.Join(credentialJsonPath, TOKEN_JSON)
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}

	return GoogleConfig{
		Config: config,
		Token:  tok,
	}
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
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

func LoadFile(configTomlPath string) (slackConfig SlackConfig,error error) {
	var sc SlackConfig
	_, err := toml.DecodeFile(path.Join(configTomlPath,CONFIG_TOML), &sc)
	if err != nil {
		return slackConfig, err
	}
	return sc, err
}