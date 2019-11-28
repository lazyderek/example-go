package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

type myHandler struct {
}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "http server ...........................................")
}

type Server struct {
	//tlsConfig *tls.Config
	caCert    string
	serverCrt string
	serverKey string

	server *http.Server

	port string
}

func _loadCA(caCertPath string) *x509.CertPool {
	pool := x509.NewCertPool()

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return nil
	}
	pool.AppendCertsFromPEM(caCrt)

	return pool
}

func newServer() *Server {
	return &Server{
		caCert:    "../etc/ca/ca.crt",
		serverCrt: "../etc/ca/server.crt",
		serverKey: "../etc/ca/server.key",
		port:      ":9091",
	}
}

func (s *Server) Run() error {
	pool := _loadCA(s.caCert)

	server := http.Server{
		Addr:    s.port,
		Handler: &myHandler{},
		TLSConfig: &tls.Config{
			ClientCAs: pool,
			//ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	s.server = &server

	return s.server.ListenAndServeTLS(s.serverCrt, s.serverKey)
}

func (s *Server) Close() {
	if err := s.server.Shutdown(context.Background()); err != nil {
		fmt.Printf("http server close, err: %v\n", err)
		return
	}
}

func main() {
	s := newServer()

	if err := s.Run(); err != nil {
		fmt.Printf("http servr run, err: %v", err)
		return
	}
}
