package core

import (
	"fmt"
	"net/http"
	"time"
)

var startTime = time.Now().Add(time.Second)

func Uptime() time.Duration {
	return time.Since(startTime)
}

func UptimeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", Uptime())
}
