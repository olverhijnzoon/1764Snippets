package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("# 1764Snippets")
	fmt.Println("## Golang clientCertificate")

	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatal("Error loading server certificate and key:", err)
	}

	// Create a pool with the client certificates
	clientCertPool := x509.NewCertPool()
	clientCert, err := os.ReadFile("client.crt")
	if err != nil {
		log.Fatal("Error reading client certificate:", err)
	}
	clientCertPool.AppendCertsFromPEM(clientCert)

	// Create a TLS config with the client pool and server certificate
	tlsConfig := &tls.Config{
		ClientCAs:    clientCertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{serverCert},
	}

	// Create a server instance with the TLS config
	server := &http.Server{
		Addr:      ":8071",
		TLSConfig: tlsConfig,
	}

	// Define a handler for the start page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.TLS != nil && len(r.TLS.PeerCertificates) > 0 {
			w.Write([]byte("1764Snippets welcome"))
		} else {
			http.Error(w, "Client certificate required", http.StatusUnauthorized)
		}
	})

	// Start the server with TLS
	log.Println("Starting secure server on port 8071...")
	log.Fatal(server.ListenAndServeTLS("", ""))
}
