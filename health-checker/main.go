package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	url        string
	StatusCode int
	err        error
}

func main() {

	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://httpbin.org/status/200",
		"https://httpbin.org/delay/10",
		"https://invalid-url",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	messages := SendRequest(ctx, urls)

	for msg := range messages {

		if msg.err != nil {
			fmt.Printf("ERROR [%s]: %v\n\n", msg.url, msg.err)
			continue
		}

		fmt.Printf("%s is healthy. Status code : %d\n\n", msg.url, msg.StatusCode)
	}
}

func SendRequest(ctx context.Context, urls []string) <-chan Result {
	out := make(chan Result)

	var wg sync.WaitGroup

	wg.Add(len(urls))

	client := &http.Client{}

	for _, url := range urls {
		go func(u string) {
			defer wg.Done()
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
			if err != nil {
				out <- Result{url: u, err: err}
				return
			}
			res, err := client.Do(req)
			if err != nil {
				out <- Result{
					url: u,
					err: err,
				}
				return
			}
			defer res.Body.Close()

			out <- Result{
				url:        u,
				StatusCode: res.StatusCode,
				err:        nil,
			}
		}(url)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
