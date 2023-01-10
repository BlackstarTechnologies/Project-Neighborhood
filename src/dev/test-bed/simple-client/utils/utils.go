package utils

import (
	"errors"
	"strings"
)

func SplitOnComma(text string) (res string) {
	res = strings.Join(strings.Split(text, ","), ",\n")
	return
}

var trackBody string

func ShowResponse(body string, err error, openstring string) {
	// openstring := "Version1Get(https://api.github.com/users/KrunalLathiya)\n"
	if trackBody != "" {
		if body == trackBody {
			println(openstring, "Response: same \n\n")
			return
		} else {
			err = errors.New("incorrect body format")
		}
	}

	if err != nil {
		println(openstring, "Error:\n", err.Error())
		return
	}
	println(openstring, "Response:\n", SplitOnComma(body), "\n\n")
	trackBody = body
}
