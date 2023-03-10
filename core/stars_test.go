package core

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eleven26/awesome-go-stars/contract"
	"github.com/eleven26/awesome-go-stars/core/mocks"
)

func TestReadmeNotRepoUrl(t *testing.T) {
	l := new(mocks.Link)
	l.On("IsRepoUrl").Return(false)

	puller := new(mocks.Puller)

	GetStars([]contract.Link{l}, puller)

	l.AssertExpectations(t)
}

func TestReadmePullFails(t *testing.T) {
	l := new(mocks.Link)
	l.On("IsRepoUrl").Return(true)
	l.On("Url").Return("")
	l.On("ApiEndpoint").Return("")

	result := new(mocks.PullResult)
	result.On("Ok").Return(false)

	puller := new(mocks.Puller)
	puller.On("Pull", l.Url()).Return(result)

	GetStars([]contract.Link{l}, puller)

	l.AssertExpectations(t)
	result.AssertExpectations(t)
	puller.AssertExpectations(t)
}

func TestReadmePullSuccess(t *testing.T) {
	l := new(mocks.Link)
	l.On("IsRepoUrl").Return(true)
	l.On("Url").Return("foo")
	l.On("ApiEndpoint").Return("foo")

	result := new(mocks.PullResult)
	result.On("Ok").Return(true)
	result.On("Stars").Return(1)

	puller := new(mocks.Puller)
	puller.On("Pull", l.Url()).Return(result)

	res := GetStars([]contract.Link{l}, puller)
	assert.Len(t, res, 1)
	assert.Equal(t, 1, res["foo"])

	l.AssertExpectations(t)
	result.AssertExpectations(t)
	puller.AssertExpectations(t)
}
