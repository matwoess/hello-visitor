package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Visitor!<br/>")
	fmt.Fprintf(w, "The current date and time is: %s\n", getDateTimeString(time.Now()))
}

func getDateTimeString(t time.Time) string {
	return fmt.Sprintf(
		"%d.%d.%d %d:%d:%d",
		t.Day(),
		t.Month(),
		t.Year(),
		t.Hour(),
		t.Minute(),
		t.Second(),
	)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
