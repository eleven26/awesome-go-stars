package core

import (
	"io"
	"net/http"

	"github.com/eleven26/awesome-go-stars/contract"
)

var _ contract.Puller = &puller{}

type puller struct {
	token string
}

func NewPuller(token string) contract.Puller {
	return &puller{
		token: token,
	}
}

func (p *puller) Pull(url string) contract.PullResult {
	statusCode, content := p.pull(url)

	return NewPullResult(statusCode, content)
}

func (p *puller) pull(url string) (statusCode int, content []byte) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return -1, nil
	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+p.token)
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	res, err := client.Do(req)
	if err != nil {
		return res.StatusCode, nil
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return -1, nil
	}

	return res.StatusCode, data
}
