// server.go
// brigham toskin 2015

package main

import (
	"log"
	"net/http"
	"io"
	"strconv"
)

const (
	BASE10 = 10
	PORT = ":8080"
	REQUEST = "/counter/"
)

var (
	counts map[string]int
)

func init() {
	counts = make(map[string]int)
}

func fmtint(val int) (string) {
	return strconv.FormatInt(int64(val), BASE10)
}

func serveCounter(w http.ResponseWriter, count int) {
	scount := fmtint(count)
	log.Printf("  Count: %s", scount)
	header := w.Header()
	header.Set("Content-Length", fmtint(1000000))
	log.Printf("Set response content length: %s", header.Get("Content-Length"))
	
	_, err := io.WriteString(w, scount)
	if err != nil {
		log.Fatal("Writing Response: ", err)
	}
}

func HandleRequest(w http.ResponseWriter, req *http.Request) {
	var count int
	switch req.Method {
	case "GET":
		count = counts[req.URL.Path] + 1
	case "DELETE":
		count = 0
	default:
		log.Printf("Unknown request method: %s", req.Method)
		return
	}
	counts[req.URL.Path] = count
	log.Printf("Request: %s %s ", req.Method, req.URL.Path)
	go serveCounter(w, count)
}

func main() {
	log.Print("Brig and Misham's Amazing Counter Server, ver. 1.0!!!")

	http.HandleFunc(REQUEST, HandleRequest)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
