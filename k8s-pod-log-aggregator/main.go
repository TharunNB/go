package main

import (
	"fmt"
	"strings"
)

func main() {
	namespace := "test-logs"
	maxWorkers := 10

	fmt.Printf("Fetching logs from pods in namespace: %s (max %d concurrent)\n\n", namespace, maxWorkers)

	logsMap, errors, _ := FetchLogsConcurrently(namespace, maxWorkers)

	fmt.Printf("Success: %d pods\n", len(logsMap))
	fmt.Printf("Failed: %d pods\n\n", len(errors))

	for pod, logs := range logsMap {
		fmt.Printf("=== %s ===\n", pod)

		lines := strings.Split(logs, "\n")
		for i, line := range lines {
			if i >= 8 {
				fmt.Println("... (truncated)")
				break
			}
			if line != "" {
				fmt.Println(line)
			}
		}
		fmt.Println("---")
	}

	if len(errors) > 0 {
		fmt.Println("Errors:")
		for _, e := range errors {
			fmt.Println(e)
		}
	}
}
