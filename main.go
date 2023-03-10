package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

func main() {
	cmd := &cobra.Command{
		Use:   "",
		Short: "Awesome Go Stars",
		Long:  "Awesome Go Stars",
		Run: func(cmd *cobra.Command, args []string) {
			token, _ := cmd.Flags().GetString("token")

			_ = token
		},
		Args: cobra.ExactArgs(0),
	}

	cmd.Flags().StringP("token", "t", "", "github token")

	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func handle() {
	// 1. parse readme
	// 2. get stars
	// 3. update readme
}
