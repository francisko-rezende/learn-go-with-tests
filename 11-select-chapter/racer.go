package selectchapter

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(slowUrl, fastUrl string) (winner string, err error) {
	select {
	case <-ping(slowUrl):
		return slowUrl, nil
	case <-ping(fastUrl):
		return fastUrl, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", slowUrl, fastUrl)
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
