package main

import (
	"fmt"
	"log"
	"net/http"
)

// Hello returns hello with the name you pass in :)
func Hello(name string) string {
	result := "Hello " + name
	fmt.Println(result)
	return result
}

func handler(w http.ResponseWriter, r *http.Request) {
	var input = "World"
	if len(r.URL.Path) > 1 {
		input = r.URL.Path[1:]
	}
	fmt.Println(w, Hello(input))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
