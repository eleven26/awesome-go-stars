package core

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPullResultOk(t *testing.T) {
	pr := NewPullResult(-1, nil)
	assert.False(t, pr.Ok())

	pr = NewPullResult(http.StatusOK, nil)
	assert.True(t, pr.Ok())
}

func TestPullResultStars(t *testing.T) {
	pr := NewPullResult(http.StatusOK, []byte(`{"stargazers_count": 42}`))
	assert.Equal(t, 42, pr.Stars())

	pr = NewPullResult(http.StatusOK, []byte(`{"stargazers_count": "foo"}`))
	assert.Equal(t, -1, pr.Stars())

	pr = NewPullResult(http.StatusOK, []byte(`foo`))
	assert.Equal(t, -1, pr.Stars())

	pr = NewPullResult(http.StatusNotFound, []byte(`foo`))
	assert.Equal(t, -1, pr.Stars())
}
