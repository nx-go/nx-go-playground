package main

import (
  "fmt"
  "log"
  "net/http"
)

func Hello(name string) string {
	result := "Hello " + name
	return result
}


func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, Hello(r.URL.Path[1:]))
}

func main() {
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

