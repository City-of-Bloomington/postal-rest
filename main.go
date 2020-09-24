package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	expand "github.com/openvenues/gopostal/expand"
	parser "github.com/openvenues/gopostal/parser"
)

func main() {
	host := os.Getenv("LISTEN_HOST")
	if host == "" {
       host = "0.0.0.0"
	}
	port := os.Getenv("LISTEN_PORT")
	if port == "" {
       port = "8080"
	}
	listenSpec := fmt.Sprintf("%s:%s", host, port)
	certFile   := os.Getenv("SSL_CERT_FILE")
	keyFile    := os.Getenv("SSL_KEY_FILE")

	router := mux.NewRouter()
	router.HandleFunc("/health", HealthHandler).Methods("GET")
	router.HandleFunc("/expand", ExpandHandler).Methods("GET")
	router.HandleFunc("/parse",  ParserHandler).Methods("GET")

	s := &http.Server{Addr: listenSpec, Handler: router}
	go func() {
		if certFile != "" && keyFile != "" {
			fmt.Printf("listening on https://%s\n", listenSpec)
			s.ListenAndServeTLS(certFile, keyFile)
		} else {
			fmt.Printf("listening on http://%s\n", listenSpec)
			s.ListenAndServe()
		}
	}()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	<-stop
	fmt.Println("\nShutting down the server...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	s.Shutdown(ctx)
	fmt.Println("Server stopped")
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func ExpandHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

    query, ok  := r.URL.Query()["address"]
	expansions := expand.ExpandAddress(query)
	output, _  := json.Marshal(expansions)
	w.Write(output)
}

func ParserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

    query, ok := r.URL.Query()["address"]
	parsed    := parser.ParseAddress(query)
	output, _ := json.Marshal(parsed)
	w.Write(output)
}
