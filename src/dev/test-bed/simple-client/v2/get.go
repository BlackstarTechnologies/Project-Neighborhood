package v2

import (
	// "fmt"
	"fmt"
	"io"
	"net/http"
)

type HttpClient struct {
	Address string
	resp    *http.Response
	body    []byte
	length  int
	err     error
}

func NewClient(address string) HttpClient {
	return HttpClient{
		Address: address,
	}
}

func (H *HttpClient) Get(url string) *HttpClient {
	H.resp, H.err = http.Get(fmt.Sprint(H.Address, url))
	return H
}
func (H *HttpClient) ReadString() string {
	if (H.length < 1) && H.err == nil {
		defer H.resp.Body.Close()
		H.body, H.err = io.ReadAll(H.resp.Body)

	}
	return string(H.body)
}

func Get(url string) (res string, err error) {
	// "https://api.github.com/users/KrunalLathiya"
	myclient := NewClient(url)
	myclient.Get("")
	return myclient.ReadString(), myclient.err
}
