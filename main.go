package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Visitor!<br/>")
	t := time.Now()
	fmt.Fprintf(w, "The current time is %d:%d:%d", t.Hour(), t.Minute(), t.Second())
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
