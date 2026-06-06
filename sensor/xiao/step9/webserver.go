package main

import (
	"net/http"

	_ "embed"
)

//go:embed index.html
var page string

// Uses Min - a tiny framework that makes websites pretty.
// See https://mincss.com/
//
//go:embed mincss.min.css
var mincss string

//go:embed tetromino.html
var tetromino string

var (
	responseActive         = []byte("system active")
	responseInactive       = []byte("system inactive")
	responseStatusActive   = []byte(`{"status": "active"}`)
	responseStatusInactive = []byte(`{"status": "inactive"}`)
)

func startWebServer() {
	h, _ := link.Addr()
	host := h.String()
	println("HTTP server listening on http://" + host + port)

	http.HandleFunc("/", root)
	http.HandleFunc("/mincss.min.css", css)
	http.HandleFunc("/6", sixlines)
	http.HandleFunc("/on", systemActivate)
	http.HandleFunc("/off", systemDeactivate)
	http.HandleFunc("/status", systemStatus)

	err := http.ListenAndServe(host+port, nil)
	for err != nil {
		failMessage("error: " + err.Error())
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(page))
}

func css(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(mincss))
}

// https://fukuno.jig.jp/3267
func sixlines(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(tetromino))
}

func systemActivate(w http.ResponseWriter, r *http.Request) {
	systemActive = true
	w.Header().Set(`Content-Type`, `text/plain; charset=UTF-8`)
	w.Write(responseActive)
}

func systemDeactivate(w http.ResponseWriter, r *http.Request) {
	systemActive = false
	w.Header().Set(`Content-Type`, `text/plain; charset=UTF-8`)
	w.Write(responseInactive)
}

func systemStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(`Content-Type`, `application/json`)
	if systemActive {
		w.Write(responseStatusActive)
	} else {
		w.Write(responseStatusInactive)
	}
}
