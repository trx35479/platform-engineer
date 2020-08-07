package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Return struct {
	Hostname string
	Date int64
	Path string
	Method string
}

func main() {
	var h Return
	http.HandleFunc("/", h.handler)
	log.Println("Starting app")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (r Return) handler(w http.ResponseWriter, req *http.Request) {
	r.Date = time.Now().Unix()
	name, err := os.Hostname()
	if err != nil {
		log.Fatalf("Error getting hostname: %v", err)
	}
	r.Hostname = name
	r.Path = req.RequestURI
	r.Method = req.Method
	data, err := json.Marshal(r)
	if err != nil {
		log.Fatal("Error structuring data")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}