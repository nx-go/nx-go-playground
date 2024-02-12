package main

import (
	"fmt"
	"net/http"
	"nx-go-playground/logger"
	"nx-go-playground/math"
)

func clamp(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "clamped: %.2f\n", math.Clamp(84, 10, 55))
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world!\n")
}

func main() {
	http.HandleFunc("/clamp", clamp)
	http.HandleFunc("/hello", hello)
	logger.Info("Listening on port 8090")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		logger.Error(fmt.Sprintf("HTTP server error: %v", err))
	}
}
