package selectchapter

import (
	"net/http"
	"time"
)

func Racer(slowUrl, fastUrl string) (winner string) {
	startSlowUrl := time.Now()
	http.Get(slowUrl)
	slowUrlDuration := time.Since(startSlowUrl)

	startFastUrl := time.Now()
	http.Get(fastUrl)
	fastUrlDuration := time.Since(startFastUrl)

	if slowUrlDuration < fastUrlDuration {
		return slowUrl
	}

	return fastUrl
}
