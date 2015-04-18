// server.go
// brigham toskin 2015

package main

import (
	"log"
	"net/http"
)

const (
	PORT    = ":8080"
	REQUEST = "/counter/"
)

var (
	counts map[string]int64
)

func init() {
	counts = make(map[string]int64)
}

func serveCounter(w http.ResponseWriter, count int64) {
	result, err := GetCountPng(count)
	w.Header().Set("Content-Type", "image/png")

	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		result = []byte("Error showing counter")
	}
	w.Write(result)
}

func HandleRequest(w http.ResponseWriter, req *http.Request) {
	log.Printf("Request: %s %s\n", req.Method, req.URL.Path)
	count := counts[req.URL.Path] + 1
	counts[req.URL.Path] = count
	log.Printf("%s -> %d\n", req.URL.Path, count)

	serveCounter(w, count)
}

func main() {
	log.Print("Brig and Misham's Amazing Counter Server, ver. 1.0!!!")

	http.HandleFunc(REQUEST, HandleRequest)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
