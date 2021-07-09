package openany

import (
	"io"
	"net/url"
	"os"
	"sync"

	urlOpener "github.com/Reisender/openany/url"
)

type Opener func(string) (io.ReadCloser, error)

var mut sync.RWMutex
var backends map[string]Opener = map[string]Opener{
	"http":  urlOpener.Open,
	"https": urlOpener.Open,
}

func Register(scheme string, backend Opener) {
	mut.Lock()
	defer mut.Unlock()

	backends[scheme] = backend
}

func Open(uri string) (io.ReadCloser, error) {
	// identify the scheme
	parsed, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	mut.RLock()
	backend, found := backends[parsed.Scheme]
	mut.RUnlock()

	if !found {
		// default to local file if scheme was not found
		return os.Open(uri)
	}

	return backend(uri)
}
