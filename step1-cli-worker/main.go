package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	workers := flag.Int("workes", 5, "Number of worker go routines")
	flag.Parse()

	scheduler := NewScheduler(*workers)
	scheduler.pool.Start()

	fmt.Println("Scheduler started (5 workers)")
	fmt.Println("Commands: provision small | medium | large | exit")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\nShutting Down.....")
		scheduler.pool.Stop()
		os.Exit(0)
	}()

	for {
		var cmd, size string
		fmt.Printf("> ")
		fmt.Scanln(&cmd, &size)

		switch cmd {
		case "provision":
			switch size {
			case "small", "medium", "large":
				scheduler.Provision(VMSize(size))
			default:
				fmt.Println("Invalid Size Input")
				fmt.Println("Usage: provision small/medium/large")
			}
		case "exit":
			scheduler.pool.Stop()
			return
		}
	}

}
