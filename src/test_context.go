package main

import (
	"net/http"
	"time"
)

func test(w http.ResponseWriter, r *http.Request) {
	time.Sleep(20 * time.Second)
	w.Write([]byte("test"))
}

func main() {
	http.HandleFunc("/", test)
	timeoutHandler := http.TimeoutHandler(http.DefaultServeMux, 5*time.Second, "timeout")
	http.ListenAndServe(":8080", timeoutHandler)
}
