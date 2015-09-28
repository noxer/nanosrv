package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/bradfitz/http2"
)

var (
	dir     string
	port    string
	crtFile string
	keyFile string
)

func init() {
	d, err := os.Getwd()
	if err != nil {
		fmt.Errorf("Error: %s\n", err)
	}

	flag.StringVar(&dir, "d", d, "The directory that should be served.")
	flag.StringVar(&port, "p", ":8888", "The port/interface the server should work on.")
	flag.StringVar(&crtFile, "c", "", "The certificate file for HTTPS connections.")
	flag.StringVar(&keyFile, "k", "", "The private key file for HTTPS connections.")
	flag.Parse()
}

func main() {
	fmt.Println("Starting nano server...")
	fmt.Printf("Serving directory: %s\n", dir)

	if crtFile != "" && keyFile != "" {
		fmt.Printf("Using certificate: %s\nUsing private key: %s\n", crtFile, keyFile)
	} else {
		fmt.Println("No certificate and/or key provided!")
	}

	srv := &http.Server{
		Addr:    port,
		Handler: &requestLogger{http.FileServer(http.Dir(dir))},
	}
	http2.ConfigureServer(srv, nil)

	if crtFile != "" && keyFile != "" {
		fmt.Printf("Starting HTTPS server on port %s... Kill with [Ctrl]-[C]!\n", port)
		if err := srv.ListenAndServeTLS(crtFile, keyFile); err != nil {
			fmt.Errorf("Error: %s\n", err)
		}
	} else {
		fmt.Printf("Starting HTTP server on port %s... Kill with [Ctrl]-[C]!\n", port)
		if err := srv.ListenAndServe(); err != nil {
			fmt.Errorf("Error: %s\n", err)
		}
	}
}

type requestLogger struct {
	h http.Handler
}

func (l *requestLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\"%s\" --> \"%s\" started\n", r.RemoteAddr, r.URL.Path)
	l.h.ServeHTTP(w, r)
	fmt.Printf("\"%s\" --> \"%s\" finished\n", r.RemoteAddr, r.URL.Path)
}
