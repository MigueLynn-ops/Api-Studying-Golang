package handler

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(2, 5)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is the request:")
	defer r.Body.Close()

	if err := limiter.Wait(context.Background()); err != nil {
		http.Error(w, "Too many requests", http.StatusTooManyRequests)
		return
	}

	b, _ := io.ReadAll(r.Body)

	fmt.Println(string(b))

	query := strings.TrimSpace(string(b))
	apiKey := os.Getenv("GOOGLE_API_KEY")
	googleURL := "https://www.googleapis.com/books/v1/volumes?q=" + url.QueryEscape(query) + "&key=" + apiKey

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, googleURL, nil)
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Error calling external API", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)

	fmt.Printf("Awaiting new request:\n")
}
