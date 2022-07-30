package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-micorservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120 *time.Second,
		ReadTimeout: 1 *time.Second,
		WriteTimeout: 1 *time.Second,
	}

	// Start the server
	go func ()  {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// block until recive the message from the chanel
	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30 *time.Second)
	s.Shutdown(tc)
}