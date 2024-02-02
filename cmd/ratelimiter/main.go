package main

import (
	"encoding/json"
	"fmt"
	"github.com/odrianoaliveira/ratelimiter/pkg/ratelimit"
	"net/http"
	"time"
)

type pingResponse struct {
	Message string `json:"message"`
}

var TokenBucketRateLimiter *ratelimit.TokenBucket

func main() {
	// 10 tokens with a refill rate of 10 tokens per minute
	TokenBucketRateLimiter = ratelimit.NewTokenBucket(10, 10, time.Minute)

	http.HandleFunc("/ping", pingHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Process aborted. %s\n", err.Error())
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	// TODO compute tokens per user id
	if TokenBucketRateLimiter.Allow() {
		response := pingResponse{Message: "pong"}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			fmt.Println("Error encoding response.")
		}
	} else {
		http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
	}
}
