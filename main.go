package main

import (
	"encoding/base64"
	"encoding/json"
	"github.com/spf13/viper"
	"log"
	_ "net/http"
	_ "os"

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
        log.Fatalln(err)
    }
    
    // Convert []Byte type to Token
    err = json.Unmarshal(tokenBytes, &token)
    if err != nil {
        log.Fatalln(err)
    }
}

func main() {
    log.SetFlags(log.Lshortfile)
    log.SetPrefix("Backstage-Drone: ")

    vi := viper.New()
    vi.SetConfigFile("config.yaml")
    vi.ReadInConfig()
}
