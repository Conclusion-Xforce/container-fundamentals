package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Printf("Server started at %s. Listening at :8080...", time.Now())

	http.HandleFunc("/fibonacci", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		w.Header().Set("Content-Type", "application/json")

		n, err := strconv.Atoi(r.URL.Query().Get("n"))
		if err != nil {
			http.Error(w, "{\"message\": \"Invalid query parameter\"}", http.StatusBadRequest)
			return
		}
		
		start_time := time.Now()
		result := fibonacci(n)
		time_taken := time.Since(start_time)
		
		fmt.Fprintf(w, "{\"input\": %d, \"result\": %d, \"time_ns\": %d}\n", n, result, time_taken.Nanoseconds())
	})

	http.ListenAndServe(":8080", nil)
}
