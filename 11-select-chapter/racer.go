package selectchapter

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(slowUrl, fastUrl string) (winner string) {
	slowUrlDuration := measureResponseTime(slowUrl)
	fastUrlDuration := measureResponseTime(fastUrl)

	fmt.Println(slowUrlDuration, fastUrlDuration)

	if slowUrlDuration < fastUrlDuration {
		return slowUrl
	}

	return fastUrl
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
