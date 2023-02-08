package main

import (
    "bytes"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "net/url"
    "net/http"
    "strings"

    "github.com/Euvaz/Backstage-Hive/logger"
    "github.com/Euvaz/Backstage-Hive/models"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

func main() {
    viper := viper.New()
    viper.SetConfigFile("config.toml")
    err :=  viper.ReadInConfig()
    if err != nil {
        logger.Fatal(err.Error())
    }

    viper.SetDefault("host", "localhost")
    viper.SetDefault("port", 3894)
    viper.SetDefault("name", "drone-1")

    // Add root command
    cmd := &cobra.Command {
        Use:   "Backstage-Drone",
        Short: "Short Desc",
        Long:  `Long Desc`,
        PersistentPreRun: func(cmd *cobra.Command, args []string) {
        },
        Run: func(cmd *cobra.Command, args []string) {
            logger.Info("Starting server")
        },
    }

    // Add command
    enrollCmd := &cobra.Command {
        Use:   "enroll",
        Short: "Short Desc",
        Long:  `Long Desc`,
        Args: cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            var recvToken models.Token

            tokenEncoded := strings.Join(args, " ")

            // Decode Base64 token string
            tokenBytes, err := base64.StdEncoding.DecodeString(tokenEncoded)
            if err != nil {
                logger.Fatal(err.Error())
            }

            // Convert []Byte type to Token
            err = json.Unmarshal(tokenBytes, &recvToken)
            if err != nil {
                logger.Fatal(err.Error())
            }

            // POST request
            // JSON body
            dlvrToken, err := json.Marshal(models.Token{Addr: viper.GetString("host"), Port: viper.GetInt("port"), Key: recvToken.Key})
            
            // Convert String type to URL
            postUrl, err := url.Parse(fmt.Sprintf("http://%s:%d/drones/%s", recvToken.Addr, recvToken.Port, viper.GetString("name")))
            if err != nil {
                logger.Fatal(err.Error())
            }
            fmt.Println(postUrl)

            // Create an HTTP post request
            r, err := http.NewRequest("POST", postUrl.String(), bytes.NewBuffer(dlvrToken))
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
        },
    }

    // Add commands
    cmd.AddCommand(enrollCmd)

    err = cmd.Execute()
    if err != nil {
        logger.Fatal(err.Error())
    }
}

