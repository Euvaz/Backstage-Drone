package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
    "net/url"
	"net/http"

	"github.com/Euvaz/Backstage-Hive/logger"
	"github.com/spf13/viper"
)

type Token struct {
    Addr    string  `json:"addr"`
    Key     string  `json:"key"`
}

func enroll(tokenEncoded string) {
    var token Token

    // Decode Base64 token string
    tokenBytes, err := base64.StdEncoding.DecodeString(tokenEncoded)
    if err != nil {
        logger.Fatal(err.Error())
    }

    // Convert []Byte type to Token
    err = json.Unmarshal(tokenBytes, &token)
    if err != nil {
        logger.Fatal(err.Error())
    }

    // POST request
    // JSON body
	body := []byte(fmt.Sprintf(`{"key":"%s"}`, token.Key))
    
    // Convert String type to URL
    postUrl, err := url.Parse(fmt.Sprintf("http://%s", token.Addr))
    if err != nil {
        logger.Fatal(err.Error())
    }

    // Create an HTTP post request
    r, err := http.NewRequest("POST", postUrl.String(), bytes.NewBuffer(body))
	if err != nil {
		logger.Fatal(err.Error())
	}

    r.Header.Add("Content-Type", "application/json")
    client := &http.Client{}
    res, err := client.Do(r)
    if err != nil {
    	logger.Fatal(err.Error())
    }
    defer res.Body.Close()
}

func main() {
    vi := viper.New()
    vi.SetConfigFile("config.toml")
    vi.ReadInConfig()

    //enroll("eyJhZGRyIjoiMTI3LjAuMC4xOjY5NjkiLCJrZXkiOiJ0ZXN0In0=")
}
