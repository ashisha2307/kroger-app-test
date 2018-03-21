package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "New Pipeline - Test Notifications")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":80", nil)
}
