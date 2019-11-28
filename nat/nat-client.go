package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Client struct {
	sync.WaitGroup

	clients []*net.UDPAddr

	Server *net.UDPAddr
}

func new() *Client {
	return &Client{
		clients: []*net.UDPAddr{
			{IP: net.IPv4zero, Port: 9093},
			//{IP: net.IPv4zero, Port: 9094},
		},
		Server: &net.UDPAddr{IP: net.ParseIP("52.194.245.198") /*net.IPv4()*/, Port: 9092},
	}
}

func (c *Client) Connect() {
	fmt.Println("clients", len(c.clients))
	for _, cli := range c.clients {
		cli := cli
		c.Wrap(func() {
			conn, err := net.DialUDP("udp", cli, c.Server)
			if err != nil {
				fmt.Printf("DialUDP err: %v", err)
				return
			}
			var data = make([]byte, 1024)
			n, err := conn.Write(data)
			if err != nil {
				fmt.Printf("Write err: %v, n=%v", err, n)
				return
			}

			for {
				n, addr, err := conn.ReadFromUDP(data)
				if err != nil {
					return
				}

				// todo: parse data
				fmt.Println(cli.String(), addr.String(), n, string(data))

				// write to another udp
				if _, err := conn.WriteToUDP([]byte("hello"), parseIp(string(data))); err != nil {
					fmt.Printf("WriteToUDP err=%v, from=%v, to=%v\n", err, cli.String(), string(data))
					return
				}

				time.Sleep(time.Second * 10)
			}

		})
	}
}

func (c *Client) Wrap(fc func()) {
	c.Add(1)
	go func() {
		defer c.Done()
		fc()
	}()
}

func main() {
	client := new()

	client.Connect()

	client.Wait()

}

func parseIp(addr string) *net.UDPAddr {
	ip := strings.Split(addr, ":")
	port, err := strconv.Atoi(ip[1][:4])
	if err != nil {
		fmt.Printf("parseIp err=%v", err)
	}
	//fmt.Println(ip[0], ip[1], port)
	return &net.UDPAddr{
		IP:   net.ParseIP(ip[0]),
		Port: port,
	}
}
