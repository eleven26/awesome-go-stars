package core

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		assert.Equal(t, "application/vnd.github+json", r.Header.Get("Accept"))
		assert.Equal(t, "2022-11-28", r.Header.Get("X-GitHub-Api-Version"))

		if r.Header.Get("Authorization") == "Bearer foo" {
			_, _ = w.Write([]byte(`{"rate":{"limit":5000,"remaining":3000,"reset":1637740000,"used":0}}`))
		} else {
			_, _ = w.Write([]byte(`{"rate":{"limit":5000,"remaining":4000,"reset":1637740000,"used":0}}`))
		}
	}))

	err := CheckRateLimit("foo", ts.URL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "less than 4000")
	assert.Contains(t, err.Error(), "remaining: 3000")

	err = CheckRateLimit("bar", ts.URL)
	assert.Nil(t, err)
}
