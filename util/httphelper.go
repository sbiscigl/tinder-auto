package util

import (
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

var headerMap = map[string]string{
	"user-agent":   "Tinder/4.0.9 (iPhone; iOS 8.0.2; Scale/2.00)",
	"content-type": "application/json",
}

/*HTTPHelper helper for request functions*/
type HTTPHelper struct {
	client http.Client
}

/*HTTPHeader type for a key value for a header*/
type HTTPHeader struct {
	Key   string
	Value string
}

/*New constructor for helper function*/
func New() *HTTPHelper {
	return &HTTPHelper{
		client: http.Client{},
	}
}

/*MakeReq makes a http request*/
func (h *HTTPHelper) MakeReq(method string, urlStr string,
	body io.Reader, headerList []HTTPHeader) []byte {

	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		panic("could not create request")
	}

	for k, v := range headerMap {
		req.Header.Add(k, v)
	}

	for _, header := range headerList {
		req.Header.Add(header.Key, header.Value)
	}

	resp, err := h.client.Do(req)
	if err != nil {
		panic("could not execute response")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("got a bad request with: " + strconv.Itoa(resp.StatusCode))
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("could not read body")
	}
	return respBody
}
