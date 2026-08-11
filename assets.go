package main

import (
	_ "embed"
	"encoding/base64"
)

//go:embed assets/colony.png
var bgImageBytes []byte

func bgImageDataURI() string {
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(bgImageBytes)
}