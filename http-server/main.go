package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

type Server struct {
	totalRequests int64
	startTime     time.Time
}

func (s *Server) healthcheck(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	atomic.AddInt64(&s.totalRequests, 1)

	w.Header().Set("Content-Type", "application/json")

	status := map[string]interface{}{
		"status": "ok",
		"uptime": int(time.Since(s.startTime).Seconds()),
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

func (s *Server) metrics(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	atomic.AddInt64(&s.totalRequests, 1)

	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprintf(w, "requests_total %d\n\n", atomic.LoadInt64(&s.totalRequests))
}

func main() {

	srv := &Server{
		totalRequests: 0,
		startTime:     time.Now(),
	}

	http.HandleFunc("/health", srv.healthcheck)
	http.HandleFunc("/metrics", srv.metrics)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
