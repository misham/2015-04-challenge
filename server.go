// server.go
// brigham toskin 2015

package main

import (
	"log"
	"net/http"
	"strconv"
)

const (
	PORT = ":8080"
	REQUEST = "/counter/"
)

var (
	counts map[string]int32
)

func init() {
	counts = make(map[string]int32)
}

func serveCounter(w http.ResponseWriter, count int32) {
	// w.Write(([]byte)strconv.FormatInt(count, 10))
}

func HandleRequest(w http.ResponseWriter, req *http.Request) {
	log.Printf("Request: %s %s\n", req.Method, req.URL.Path)
	count := counts[req.URL.Path] + 1
	counts[req.URL.Path] = count
	go serveCounter(w, count)
}

func main() {
	log.Print("Brig and Misham's Amazing Counter Server, ver. 1.0!!!")

	http.HandleFunc(REQUEST, HandleRequest)
	for {
		err := http.ListenAndServe(PORT, nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}
}
