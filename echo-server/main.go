package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var (
	mode = flag.String("mode", "http", "Run the server in following mode - HTTP, HTTPS, gRPC [default: http]")
	port = flag.String("port", "8080", "Server listen port [default: 8080]")
	cert = flag.String("cert", "/etc/certs/cert.pem", "HTTPS Cert location [default: /etc/certs/cert.pem")
	key  = flag.String("key", "/etc/certs/key.pem", "HTTPS Key location [default: /etc/certs/key.pem")
)

func main() {
	flag.Parse()
	var err error
	switch {
	case strings.EqualFold(*mode, "http"):
		err = httpServer(":" + *port)
	case strings.EqualFold(*mode, "https"):
		err = httpsServer(":" + *port)
	default:
		log.Fatal("Unknown server mode")
	}
	if err != nil {
		log.Fatal("failed to start server err", err)
	}
}

// HTTPS
func httpsServer(addr string) error {
	http.HandleFunc("/", echo)
	return http.ListenAndServeTLS(addr, *cert, *key, nil)
}

func httpServer(addr string) error {

	http.HandleFunc("/", echo)
	return http.ListenAndServe(addr, nil)
}

func echo(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	response := strings.TrimPrefix(req.RequestURI, "/")
	fmt.Println("responding to request with ", response)
	fmt.Fprintf(w, response)
}
