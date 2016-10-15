package sense

import (
	"encoding/json"
	"os"
	"path"
)

type Client struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type Token struct {
	TokenType    string `json:"token_type"`
	Name         string `json:"name"`
	AccountID    string `json:"account_id"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type Sense struct {
	client     Client
	token      Token
	userAgent  string
	serverBase string
}

func NewFromDir(dir string) (*Sense, error) {
	client := Client{}
	token := Token{}
	for fPath, obj := range map[string]interface{}{
		"client_keys": &client,
		"token":       &token,
	} {
		fd, err := os.Open(path.Join(dir, fPath))
		if err != nil {
			return nil, err
		}
		defer fd.Close()
		if err := json.NewDecoder(fd).Decode(obj); err != nil {
			return nil, err
		}
	}
	return New(client, token)
}

func New(client Client, token Token) (*Sense, error) {
	return &Sense{
		client:     client,
		token:      token,
		userAgent:  "pault.ag/go/sense",
		serverBase: "https://api.hello.is/",
	}, nil
}
