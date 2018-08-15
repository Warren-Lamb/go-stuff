package future

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFutureToReturnChannels(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, "Golang Concurrency")

	}))
	defer ts.Close()

	futureChan := futureData(ts.URL)

	expectedChan := data{Body: []byte("Golang Concurrency"), Error: nil}

	assert.Equal(t, expectedChan, <-futureChan)
}
