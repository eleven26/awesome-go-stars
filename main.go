package main

import (
	"bufio"
	"io"
	"net/http"
	"os"

	"github.com/eleven26/awesome-go-stars/core"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const ReadmeUrl = "https://raw.githubusercontent.com/avelino/awesome-go/master/README.md"

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
			debug, _ := cmd.Flags().GetBool("debug")

			handle(token, debug)
		},
		Args: cobra.ExactArgs(0),
	}

	cmd.Flags().StringP("token", "t", "", "github token")
	cmd.Flags().BoolP("debug", "d", false, "debug mode")
	_ = cmd.MarkFlagRequired("token")

	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func handle(token string, debug bool) {
	var content string
	var scanner *bufio.Scanner
	var file string

	if debug {
		file = "test_README.md"
		content = local(file)
	} else {
		content = get()
		file = "README.md"
	}

	links := core.Parse(content)

	puller := core.NewPuller(token)

	stars := core.GetStars(links, puller)

	scanner = getScanner(file)
	res, err := core.Replace(scanner, stars)
	if err != nil {
		log.Error("replace readme error")
		log.Fatal(err)
	}

	err = os.WriteFile(file, []byte(res), 0o644)
	if err != nil {
		log.Fatal(err)
	}
}

func get() string {
	resp, err := http.Get(ReadmeUrl)
	if err != nil {
		log.Error("get readme error")
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Error("get readme error")
		log.Fatal(resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("read readme error")
		log.Fatal(err)
	}

	return string(data)
}

func getScanner(f string) *bufio.Scanner {
	file, err := os.Open(f)
	if err != nil {
		log.Error("open readme error")
		log.Fatal(err)
	}

	return bufio.NewScanner(file)
}

func local(f string) string {
	file, err := os.Open(f)
	if err != nil {
		log.Error("open readme error")
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Error("read readme error")
		log.Fatal(err)
	}

	return string(data)
}
