package core

import (
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

// https://docs.github.com/en/rest/rate-limit?apiVersion=2022-11-28

const RateLimitUrl = "https://api.github.com/rate_limit"

type RateLimitResponse struct {
	Rate struct {
		Limit     int `json:"limit"`
		Remaining int `json:"remaining"`
		Reset     int `json:"reset"`
		Used      int `json:"used"`
	} `json:"rate"`
}

func CheckRateLimit(token string, args ...any) error {
	url := RateLimitUrl
	if len(args) == 1 {
		url = args[0].(string)
	}

	resp, err := getRateLimitResponse(token, url)
	if err != nil {
		return errors.Wrap(err, "CheckRateLimit: could not get rate limit")
	}

	if resp.Rate.Remaining < 4000 {
		return errors.New("CheckRateLimit: rate limit is not enough (less than 4000), remaining: " + cast.ToString(resp.Rate.Remaining))
	}

	log.Info("CheckRateLimit: rate limit is enough (more than 4000), remaining: " + cast.ToString(resp.Rate.Remaining))

	return nil
}

func getRateLimitResponse(token, url string) (*RateLimitResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "github new request error.")
	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	res, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "github request error.")
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "github read response. error")
	}

	resp := RateLimitResponse{}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "github unmarshal response. error")
	}

	return &resp, nil
}
