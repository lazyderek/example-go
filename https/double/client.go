package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	caCrt     string
	clientCrt string
	clientKey string

	url    string
	client *http.Client
}

func loadCA(caCertPath string) *x509.CertPool {
	pool := x509.NewCertPool()

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return nil
	}
	pool.AppendCertsFromPEM(caCrt)

	return pool
}

func newClient() *Client {
	client := Client{
		url:       "https://localhost:9091",
		caCrt:     "../etc/ca/ca.crt",
		clientCrt: "../etc/ca/client.crt",
		clientKey: "../etc/ca/client.key",
	}

	pool := loadCA(client.caCrt)
	cliCrt, err := tls.LoadX509KeyPair(client.clientCrt, client.clientKey)
	if err != nil {
		fmt.Println("Loadx509keypair err:", err)
		return nil
	}

	tr := http.Transport{
		TLSClientConfig: &tls.Config{
			//InsecureSkipVerify: true,
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}

	client.client = &http.Client{Transport: &tr}

	return &client
}

func (c *Client) Do() (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, c.url, nil)
	if err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func main() {
	cli := newClient()

	rsp, err := cli.Do()
	if err != nil {
		fmt.Printf("cli.Do, err: %v\n", err)
		return
	}
	defer rsp.Body.Close()

	fmt.Println("statusCode", rsp.StatusCode)

	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll, err: %v\n", err)
		return
	}
	fmt.Println("body", string(b))

}
