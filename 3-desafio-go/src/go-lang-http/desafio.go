package main

import (
  "fmt"
  "log"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, greeting("Code.education Rocks!"))
}

func greeting(greet string) string {
	return "<b>" + greet + "</b>"
}

func main() {
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":8000", nil))
}