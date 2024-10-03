package main

import (
	"fmt"
	"os"
	"os/signal"
	"stock/lib/web/server"
	"sync"
	"syscall"
)

func runWebServer(wg *sync.WaitGroup, sigs chan os.Signal) {
	defer wg.Done()

	s := server.New(server.Option{Port: 8080})
	go func() {
		if err := s.Run(); err != nil {
			fmt.Println("Error starting server:", err)
		}
		fmt.Println("Server start")
	}()

	sig := <-sigs
	fmt.Printf("Received signal: %s. Shutting down...\n", sig)

	if err := s.Stop(); err != nil {
		fmt.Println("Error shutting down server:", err)
	}
}

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		runWebServer(&wg, sigs)
		fmt.Println("Server shut down successfully.")
	}()

	wg.Wait()
	fmt.Println("Done!")
}
