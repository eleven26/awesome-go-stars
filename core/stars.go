package core

import (
	"sync"

	"github.com/eleven26/awesome-go-stars/contract"
)

var mu sync.Mutex

func GetStars(links []contract.Link, puller contract.Puller) map[string]int {
	result := make(map[string]int)

	tickets := 10
	ch := make(chan struct{}, tickets)
	var wg sync.WaitGroup

	for _, l := range links {
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

			res := puller.Pull(link.ApiEndpoint())
			if !res.Ok() {
				return
			}

			mu.Lock()
			result[link.Url()] = res.Stars()
			mu.Unlock()
		}(l)
	}

	wg.Wait()

	return result
}
