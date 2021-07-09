package openany_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/matryer/is"

	"github.com/Reisender/openany"
)

// func Register(scheme string, backend Opener)
func TestRegister(t *testing.T) {
	is := is.New(t)
	scheme := "fooscheme"
	payload := "hello world"
	url := scheme + "://some/random/path"
	r := ioutil.NopCloser(strings.NewReader(payload))

	// register the testing backend
	openany.Register(scheme, func(path string) (io.ReadCloser, error) {
		is.Equal(path, url)
		return r, nil
	})

	reader, err := openany.Open(url)
	is.NoErr(err)

	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	reader.Close()

	is.Equal(buf.String(), payload)
}
