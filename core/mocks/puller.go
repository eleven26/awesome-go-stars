package mocks

import (
	"github.com/eleven26/awesome-go-stars/contract"
	"github.com/stretchr/testify/mock"
)

var _ contract.Puller = &Puller{}

type Puller struct {
	mock.Mock
}

func (p *Puller) Pull(url string) contract.PullResult {
	args := p.Called(url)
	return args.Get(0).(contract.PullResult)
}

var _ contract.PullResult = &PullResult{}

type PullResult struct {
	mock.Mock
}

func (p *PullResult) Stars() int {
	args := p.Called()
	return args.Int(0)
}

func (p *PullResult) Ok() bool {
	args := p.Called()
	return args.Bool(0)
}
