package core

import (
	"sync"

	"github.com/eleven26/awesome-go-stars/contract"
)

type readme struct {
	content string

	links  []contract.Link
	puller contract.Puller
}

func NewReadme(content string, links []contract.Link, puller contract.Puller) contract.Readme {
	return &readme{
		content: content,
		links:   links,
		puller:  puller,
	}
}

func (r *readme) Links() []contract.Link {
	return r.links
}

func (r *readme) Raw() string {
	return r.content
}

func (r *readme) GetStars() map[string]int {
	result := make(map[string]int)

	tickets := 10
	ch := make(chan struct{}, tickets)
	var wg sync.WaitGroup

	for _, l := range r.links {
		ch <- struct{}{}
		wg.Add(1)

		go func(link contract.Link) {
			defer func() {
				<-ch
				wg.Done()
			}()

			if !link.IsRepoUrl() {
				return
			}

			res := r.puller.Pull(link.ApiEndpoint())
			if !res.Ok() {
				return
			}

			result[link.Url()] = res.Stars()
		}(l)
	}

	wg.Wait()

	return result
}
