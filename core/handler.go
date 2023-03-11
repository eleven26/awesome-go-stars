package core

import (
	"bufio"
	"os"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/eleven26/awesome-go-stars/contract"
)

var _ contract.Handler = &handler{}

type handler struct {
	readmePath string
	url        string
	puller     contract.Puller
}

func NewHandler(puller contract.Puller, readmePath string, url string) contract.Handler {
	return &handler{
		puller:     puller,
		readmePath: readmePath,
		url:        url,
	}
}

// Handle will fetch the readme file from the url, parse the links,
// fetch the stars of each link, and replace the readme file.
func (h *handler) Handle() error {
	content, err := getUrlContent(h.url)
	log.Debugf("fetch readme from %s.", h.url)
	if err != nil {
		return errors.Wrap(err, "Handle: could not get readme")
	}

	stars := h.stars(content)
	log.Debugf("fetch %d repos stars.", len(stars))

	log.Debugf("begin to replace file %s.", h.readmePath)
	return h.replaceFile(content, stars)
}

func (h *handler) stars(content string) map[string]int {
	links := Parse(content)

	return GetStars(links, h.puller)
}

func (h *handler) replaceFile(content string, stars map[string]int) error {
	scanner := bufio.NewScanner(strings.NewReader(content))

	res, err := Replace(scanner, stars)
	if err != nil {
		return errors.Wrap(err, "replaceFile: could not replace file")
	}

	return os.WriteFile(h.readmePath, []byte(res), 0o644)
}
