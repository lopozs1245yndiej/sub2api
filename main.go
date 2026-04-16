// sub2api - A subscription to API converter service
// Fork of Wei-Shaw/sub2api
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/sub2api/sub2api/handler"
)

const (
	defaultPort    = 8080
	defaultHost    = "127.0.0.1" // changed from 0.0.0.0 - prefer localhost-only by default
	appName        = "sub2api"
	appVersion     = "dev"
)

func main() {
	var (
		host    string
		port    int
		version bool
	)

	flag.StringVar(&host, "host", getEnvStr("HOST", defaultHost), "Host address to listen on")
	flag.IntVar(&port, "port", getEnvInt("PORT", defaultPort), "Port to listen on")
	flag.BoolVar(&version, "version", false, "Print version information and exit")
	flag.Parse()

	if version {
		fmt.Printf("%s version %s\n", appName, appVersion)
		os.Exit(0)
	}

	addr := fmt.Sprintf("%s:%d", host, port)

	mux := http.health", handler.HealthCheck)
	mux.HandleFunc("/sub", handler.SubHandler)
	mux.HandleFunc("/", handler.IndexHandler)

	log.Printf("Starting %s %s on %s", appName, appVersion, addr)

	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  30 * time.Second, // add timeouts to avoid hanging connections
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// getEnvStr retrieves a string environment variable or returns the default value.
func getEnvStr(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

// getEnvInt retrieves an integer environment variable or returns the default value.
func getEnvInt(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); ok {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
		log.Printf("Warning: invalid value for %s, using default %d", key, defaultVal)
	}
	return defaultVal
}
