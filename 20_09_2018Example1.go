package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var str ="GTTGFTGDFGCVFGFC"
	fmt.Fprintf(w, "<html><h1>Hi there,"+str+", I love %s! </html></h1>", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":787", nil))
}