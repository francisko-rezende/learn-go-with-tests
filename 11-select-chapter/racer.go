package selectchapter

import (
	"net/http"
)

func Racer(slowUrl, fastUrl string) (winner string) {
	select {
	case <-ping(slowUrl):
		return slowUrl
	case <-ping(fastUrl):
		return fastUrl
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
