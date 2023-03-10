package core

import (
	"encoding/json"
	"net/http"

	"github.com/eleven26/awesome-go-stars/contract"
)

var _ contract.PullResult = &pullResult{}

type GithubResponse struct {
	StargazersCount int `json:"stargazers_count"`
}

type pullResult struct {
	statusCode int
	content    []byte
}

func NewPullResult(statusCode int, content []byte) contract.PullResult {
	return &pullResult{
		statusCode: statusCode,
		content:    content,
	}
}

func (p *pullResult) Stars() int {
	if !p.Ok() {
		return -1
	}

	res := GithubResponse{}
	err := json.Unmarshal(p.content, &res)
	if err != nil {
		return -1
	}

	return res.StargazersCount
}

func (p *pullResult) Ok() bool {
	return p.statusCode == http.StatusOK
}
