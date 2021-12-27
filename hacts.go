package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FileRequest(rw http.ResponseWriter, r *http.Request) {
	if b, e := io.ReadAll(r.Body); e == nil {
		if r.Method != http.MethodPost {
			return
		}

		var payload FileReq
		json.Unmarshal(b, &payload)
		fmt.Println(payload.Content)

		rw.Write([]byte("{'res':'success'}"))
	}
}
