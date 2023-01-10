package test

import (
	"blackstar.go/simple-client/original"
	"blackstar.go/simple-client/utils"
	"blackstar.go/simple-client/v1"
	"blackstar.go/simple-client/v2"
)

func main() {
	{
		body, err := original.Get("https://api.github.com/users/KrunalLathiya")
		openstring := "OriginalGet(https://api.github.com/users/KrunalLathiya)\n"
		utils.ShowResponse(body, err, openstring)
	}

	{
		body, err := v1.Get("https://api.github.com/users/KrunalLathiya")
		openstring := "v1.Get(https://api.github.com/users/KrunalLathiya)\n"
		utils.ShowResponse(body, err, openstring)
	}

	{
		body, err := v2.Get("https://api.github.com/users/KrunalLathiya")
		openstring := "v2.Get(https://api.github.com/users/KrunalLathiya)\n"
		utils.ShowResponse(body, err, openstring)
	}
}
