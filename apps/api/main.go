package main

import (
	"api/libs/api/core"
	"fmt"
	"log"
	"net/http"
	"os"
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
	fmt.Fprintf(w, "%s\n", Hello(input))
}

func main() {
	var defaultHost = "localhost"
	host := os.Getenv("HOST")
	if host == "" {
		host = defaultHost
	}

	var defaultPort = "3000"
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	address := fmt.Sprintf("%s:%s", host, port)

	http.HandleFunc("/", handler)
	http.HandleFunc("/uptime", core.UptimeHandler)
	fmt.Printf("Listening on port %s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
