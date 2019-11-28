package common

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
)

func LoadCacert(caCert string) *x509.CertPool {
	pool := x509.NewCertPool()

	cert, err := ioutil.ReadFile(caCert)
	if err != nil {
		return nil
	}

	pool.AppendCertsFromPEM(cert)

	return pool
}

func LoadCertificate(config *tls.Config, certFile, keyFile string) error {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return fmt.Errorf(
			"could not load keypair %s:%s: %v", certFile, keyFile, err)
	}
	config.Certificates = []tls.Certificate{cert}

	return nil
}
