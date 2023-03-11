package core

import (
	"io"
	"net/http"

	"github.com/pkg/errors"
)

func getUrlContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", errors.Wrap(err, "getUrlContent: could not get readme")
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("getUrlContent: could not get readme")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "getUrlContent: could not read readme")
	}

	return string(data), nil
}
