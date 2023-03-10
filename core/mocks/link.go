package mocks

import (
	"github.com/eleven26/awesome-go-stars/contract"
	"github.com/stretchr/testify/mock"
)

var _ contract.Link = &Link{}

type Link struct {
	mock.Mock
}

func (l *Link) OldTitleUrl() string {
	args := l.Called()
	return args.String(0)
}

func (l *Link) IsRepoUrl() bool {
	args := l.Called()
	return args.Bool(0)
}

func (l *Link) Title() string {
	args := l.Called()
	return args.String(0)
}

func (l *Link) Url() string {
	args := l.Called()
	return args.String(0)
}

func (l *Link) Raw() string {
	args := l.Called()
	return args.String(0)
}

func (l *Link) ApiEndpoint() string {
	args := l.Called()
	return args.String(0)
}
