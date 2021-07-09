package url

import (
	"io"
	"net/http"
)

func Open(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	return resp.Body, err
}
