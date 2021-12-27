package main

import (
	"encoding/base64"
	"os"
)

// TODO: get saved location from settings
func SaveImg(raw string) error {
	b, e := base64.StdEncoding.DecodeString(raw)
	if e != nil {
		return e
	}

	f, e := os.Create("test.jpg")
	if e != nil {
		return e
	}
	defer f.Close()

	if _, e := f.Write(b); e != nil {
		return e
	}

	f.Sync()
	return nil
}
