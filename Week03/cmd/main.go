package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func startServer(ctx context.Context, port int) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "ok")
	})
	addr := fmt.Sprintf("127.0.0.1:%d", port)
	srv := http.Server{Addr: addr, Handler: mux}
	log.Println("Start server ... ")
	go func() {
		<-ctx.Done()
		err := srv.Shutdown(ctx)
		log.Printf("Shutdown server:%d succ, %v", port, err)

	}()
	return srv.ListenAndServe()
}

func captureSignal(ctx context.Context) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
			log.Printf("get signal: %s\n", s.String())
			return fmt.Errorf("kill by signal: %s", s.String())
		case syscall.SIGHUP:
		default:
			return nil
		}
	}
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	// Server port 8000
	g.Go(func() error {
		return startServer(ctx, 8000)
	})
	// Server port 8001
	g.Go(func() error {
		return startServer(ctx, 8001)
	})
	// Server port 8002
	g.Go(func() error {
		return startServer(ctx, 8002)
	})
	// Capture and handle signal
	g.Go(func() error {
		return captureSignal(ctx)
	})

	if err := g.Wait(); err != nil {
		log.Printf("Error occurs: %v ", err)
	}
	os.Exit(1)
}
