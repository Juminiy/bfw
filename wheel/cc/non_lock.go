package cc

import (
	"io/ioutil"
	"net/http"
	"sync"
)

type ResCache struct {
	FUNC  func(string) (interface{}, error)
	mu    *sync.Mutex
	cache map[string]*resEntry
}

func MakeResCache() *ResCache {
	resCache := &ResCache{}
	resCache.make()
	return resCache
}

func (rc *ResCache) make() {
	rc.FUNC = httpGetBody
	rc.mu = new(sync.Mutex)
	rc.cache = make(map[string]*resEntry)
}

func (rc *ResCache) Get(url string) (interface{}, error) {
	rc.mu.Lock()
	urlEntry := rc.cache[url]
	if urlEntry == nil {
		urlEntry = &resEntry{}
		rc.cache[url] = urlEntry
		urlEntry.ready = make(chan bool)

		rc.mu.Unlock()

		urlEntry.body, urlEntry.err = rc.FUNC(url)

		close(urlEntry.ready)
	} else {
		rc.mu.Unlock()
		<-urlEntry.ready
	}
	return urlEntry.body, urlEntry.err
}

type resEntry struct {
	ready chan bool
	res
}

type res struct {
	body interface{}
	err  error
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
