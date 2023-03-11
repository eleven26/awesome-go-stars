package main

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/eleven26/awesome-go-stars/core"
)

const ReadmeUrl = "https://raw.githubusercontent.com/avelino/awesome-go/master/README.md"

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	cmd := &cobra.Command{
		Use:   "",
		Short: "Awesome Go Stars",
		Long:  "Awesome Go Stars",
		Run:   updateReadme,
		Args:  cobra.ExactArgs(0),
	}

	cmd.Flags().StringP("token", "t", "", "github token")

	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func updateReadme(cmd *cobra.Command, args []string) {
	token, _ := cmd.Flags().GetString("token")
	if token == "" {
		token = os.Getenv("API_TOKEN")
		if token == "" {
			log.Fatal("Token is empty. please set API_TOKEN env or use -t flag to set token.")
		}
	}

	err := core.CheckRateLimit(token)
	if err != nil {
		log.Fatal(err)
	}

	path, err := filepath.Abs("./README.md")
	if err != nil {
		log.Fatal(err)
	}

	puller := core.NewPuller(token)
	handler := core.NewHandler(puller, path, ReadmeUrl)
	err = handler.Handle()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("success")
	}
}
