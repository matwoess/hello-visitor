package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello, Visitor!\n")
	_, _ = fmt.Fprintf(w, "The current date and time is: %s\n", getDateTimeString(time.Now()))
}

func getDateTimeString(t time.Time) string {
	return fmt.Sprintf(
		"%02d.%02d.%d %02d:%02d:%02d",
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
