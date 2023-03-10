package core

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var pullResponse = `{"id":1296269,"stargazers_count":10}`

func TestPuller(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(pullResponse))

		assert.Equal(t, "application/vnd.github+json", r.Header.Get("Accept"))
		assert.Equal(t, "Bearer foo", r.Header.Get("Authorization"))
		assert.Equal(t, "2022-11-28", r.Header.Get("X-GitHub-Api-Version"))
	}))

	p := &puller{token: "foo"}
	statusCode, content := p.pull(ts.URL)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, pullResponse, string(content))

	p1 := NewPuller("foo")
	result := p1.Pull(ts.URL)
	assert.Equal(t, true, result.Ok())
	assert.Equal(t, 10, result.Stars())
}
