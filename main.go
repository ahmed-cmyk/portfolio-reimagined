package main

import (
	"context"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const port = 8080

func main() {
	// router
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	// create the server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// channel to listen for OS signals
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	log.Printf("🚀 Server starting on http://localhost:%d", port)

	// run server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// block until we get SIGINT (Ctrl+C)
	<-done
	log.Println("🛑 Server stopping …")

	// shutdown with timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed: %+v", err)
	}

	log.Println("✅ Server exited properly")
}
