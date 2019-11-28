package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"guin/example-go/https/common"
	"io/ioutil"
	"net"
	"net/http"
)

type Server struct {
	serve *http.Server
	port  string

	isTls      bool
	clientAuth tls.ClientAuthType

	caCert     string
	serverCert string
	serverKey  string
}

func new() *Server {
	return &Server{
		serve: &http.Server{},
		port:  ":9091",

		isTls:      true,
		clientAuth: tls.RequireAndVerifyClientCert,

		caCert:     "./etc/ca/ca.crt",
		serverCert: "./etc/ca/server.crt",
		serverKey:  "./etc/ca/server.key",
	}
}

func (s *Server) Run() error {
	s.loadRouter()

	listener, err := s.genListener()
	if err != nil {
		return err
	}

	fmt.Println("http server running, listen on", s.port)

	return s.serve.Serve(listener)
}

func (s *Server) Shutdown() {
	if err := s.serve.Shutdown(context.Background()); err != nil {
		fmt.Printf("Shutdown err, %v", err)
		return
	}
}

func main() {
	s := new()

	if err := s.Run(); err != nil {
		fmt.Printf("Run err, %v", err)
		return
	}

	s.Shutdown()
}

func (s *Server) loadRouter() {
	mux := http.NewServeMux()

	// todo: add your routers
	mux.Handle("/", handlerFunc(&myHandler{}))
	mux.Handle("/test", handlerFunc(&myHandler{}))
	// ...

	s.serve.Handler = mux
}

func (s *Server) genListener() (net.Listener, error) {
	switch s.isTls {
	case true:
		tlsConfig, err := s.tlsConfig()
		if err != nil {
			return nil, err
		}
		return tls.Listen("tcp", s.port, tlsConfig)
	default:
		return net.Listen("tcp", s.port)
	}
}

func (s *Server) tlsConfig() (*tls.Config, error) {
	pool := common.LoadCacert(s.caCert)

	tlsConfig := &tls.Config{
		ClientCAs:  pool,
		ClientAuth: s.clientAuth,
	}

	if err := common.LoadCertificate(tlsConfig, s.serverCert, s.serverKey); err != nil {
		return nil, err
	}

	return tlsConfig, nil
}

// -------------------- handler --------------------------

func handlerFunc(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}

type myHandler struct {
}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Println("request:", string(b))

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("hello client!"))
	fmt.Println("response:", w)
}
