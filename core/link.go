package core

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/eleven26/awesome-go-stars/contract"
)

var _ contract.Link = &link{}

type link struct {
	title string
	url   string
	raw   string
}

func NewLink(title, url, raw string) contract.Link {
	return &link{
		title: title,
		url:   url,
		raw:   raw,
	}
}

func (l *link) Title() string {
	return l.title
}

func (l *link) Url() string {
	return l.url
}

func (l *link) OldTitleUrl() string {
	return fmt.Sprintf("[%s](%s)", l.title, l.url)
}

func (l *link) NewTitleUrl(star int) string {
	return fmt.Sprintf("[%s(stars: %d)](%s)", l.title, star, l.url)
}

func (l *link) Raw() string {
	return l.raw
}

func (l *link) IsRepoUrl() bool {
	if !l.isGithub() {
		return false
	}

	re := regexp.MustCompile(`https://github.com/[\w-]+/[\w-]+/?$`)

	return re.MatchString(l.url)
}

func (l *link) isGithub() bool {
	return strings.HasPrefix(l.url, "https://github.com")
}

func (l *link) ApiEndpoint() string {
	return strings.Replace(l.url, "https://github.com", "https://api.github.com/repos", 1)
}
