package ubereats

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type httpClientError struct {
	msg  string
	code int
}

func (e *httpClientError) Error() string {
	return e.msg
}

func makeHTTPClientError(url string, resp *http.Response) error {

	body, _ := ioutil.ReadAll(resp.Body)
	msg := fmt.Sprintf("HTTP request failure on %s:\n%d: %s", url, resp.StatusCode, string(body))

	return &httpClientError{
		msg:  msg,
		code: resp.StatusCode,
	}
}
