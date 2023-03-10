package main

import (
	"io"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome!")
}

func Server() {

	mux := http.NewServeMux()

	mux.HandleFunc("/welcome", welcome)

	http.ListenAndServe(":5050", mux)
}
