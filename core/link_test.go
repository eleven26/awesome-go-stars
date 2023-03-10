package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkIsRepoUrl(t *testing.T) {
	l := link{url: "https://github.com/eleven26/awesome-go-stars"}
	assert.True(t, l.IsRepoUrl())

	l = link{url: "https://github.com/eleven26/awesome-go-stars/x"}
	assert.False(t, l.IsRepoUrl())

	l = link{url: "https://github.com/eleven26"}
	assert.False(t, l.IsRepoUrl())

	l = link{url: "https://google.com/eleven26/awesome-go-stars"}
	assert.False(t, l.IsRepoUrl())
}

func TestLinkApiEndpoint(t *testing.T) {
	l := link{url: "https://github.com/eleven26/awesome-go-stars"}
	assert.Equal(t, "https://api.github.com/repos/eleven26/awesome-go-stars", l.ApiEndpoint())
}
