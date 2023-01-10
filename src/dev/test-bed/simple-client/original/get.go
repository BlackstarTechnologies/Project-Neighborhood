package original

import (
	"io/ioutil"
	"net/http"
)

func Get(url string) (body string, err error) {
	// "https://api.github.com/users/KrunalLathiya"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body_, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body_), nil
}
