package selectchapter

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares the speed of servers, returning the url with the fastest one", func(t *testing.T) {
		slowSever := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		slowUrl := slowSever.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, err := Racer(slowUrl, fastUrl)
		if err != nil {
			t.Error("Didn't expect an error but got one")
		}

		if want != got {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't return within 10 seconds", func(t *testing.T) {
		slowServer := makeDelayedServer(12 * time.Second)
		fastServer := makeDelayedServer(13 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		_, err := Racer(slowServer.URL, fastServer.URL)

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(sleepDuration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(sleepDuration)
		w.WriteHeader(http.StatusOK)
	}))
}
