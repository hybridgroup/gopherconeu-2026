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
	levelJSON              [20]byte // reused buffer for /alarmlevel JSON response
)

func startWebServer() {
	h, _ := link.Addr()
	host := h.String() + port
	println("HTTP server listening on http://" + host)

	http.HandleFunc("/", root)
	http.HandleFunc("/mincss.min.css", css)
	http.HandleFunc("/6", sixlines)
	http.HandleFunc("/on", systemActivate)
	http.HandleFunc("/off", systemDeactivate)
	http.HandleFunc("/status", systemStatus)
	http.HandleFunc("/alarmlevel", alarmlevel)

	err := http.ListenAndServe(host, nil)
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
	w.WriteHeader(http.StatusOK)
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

func alarmlevel(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var buf [32]byte
		n, _ := r.Body.Read(buf[:])
		b := buf[:n]
		for i := 0; i < n; i++ {
			if b[i] == '=' {
				end := i + 1
				for end < n && b[end] != '&' {
					end++
				}
				if end > i+1 {
					alarmLevel = uint16(bytesToInt(b[i+1 : end]))
				}
				break
			}
		}
	}

	w.Header().Set(`Content-Type`, `application/json`)
	const prefix = `{"level": `
	n := copy(levelJSON[:], prefix)
	n += uintToBytes(levelJSON[n:], uint32(alarmLevel))
	levelJSON[n] = '}'
	w.Write(levelJSON[:n+1])
}
