package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"guin/example-go/https/common"
	"io/ioutil"
	"net/http"
)

type Client struct {
	client *http.Client
	urls   []string

	isTls      bool
	caCert     string
	clientCert string
	clientKey  string
}

func newClient() *Client {
	return &Client{
		urls:       []string{"https://localhost:9091"},
		isTls:      true,
		caCert:     "./etc/ca/ca.crt",
		clientCert: "./etc/ca/client.crt",
		clientKey:  "./etc/ca/client.key",
	}
}

func (c *Client) load() {

	client := &http.Client{
		Transport: c.getTransport(),
	}

	c.client = client
}

func (c *Client) Do(url string) (string, error) {
	req, _ := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte("hello server!")))
	rsp, err := c.client.Do(req)
	if err != nil {
		fmt.Printf("url:%v, err:%v", url, err)
		return "", err
	}
	defer rsp.Body.Close()

	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func main() {
	client := newClient()

	client.load()

	for _, url := range client.urls {
		rsp, err := client.Do(url)
		if err != nil {
			return
		}
		fmt.Println(url, rsp)
	}
}

func (c *Client) getTransport() *http.Transport {
	if !c.isTls {
		return &http.Transport{}
	}

	tr := http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: common.LoadCacert(c.caCert),
		},
	}

	if err := common.LoadCertificate(tr.TLSClientConfig, c.clientCert, c.clientKey); err != nil {
		fmt.Printf("LoadCertificate err, %v", err)
		return nil
	}
	return &tr
}
