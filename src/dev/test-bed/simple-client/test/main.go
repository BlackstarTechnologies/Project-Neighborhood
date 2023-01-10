package test

import (
	"fmt"

	"blackstar.go/simple-client/original"
	"blackstar.go/simple-client/utils"
	"blackstar.go/simple-client/v1"
	"blackstar.go/simple-client/v2"
)

func Main(url string) {
	// "https://api.github.com/users/KrunalLathiya"
	{
		body, err := original.Get(url)
		openstring := fmt.Sprintf("OriginalGet(%s)\n", url)
		utils.ShowResponse(body, err, openstring)
	}

	{
		body, err := v1.Get(url)
		openstring := fmt.Sprintf("v1.Get(%s)\n", url)
		utils.ShowResponse(body, err, openstring)
	}

	{
		body, err := v2.Get(url)
		openstring := fmt.Sprintf("v2.Get(%s)\n", url)
		utils.ShowResponse(body, err, openstring)
	}
}
