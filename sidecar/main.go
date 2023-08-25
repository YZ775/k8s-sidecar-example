package main

import (
	"io"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Sidecar application\n")
	}
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
