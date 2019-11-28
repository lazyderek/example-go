package main

import (
	"fmt"
	"net"
	"sync"
)

type NatPeer interface {
	Name() string
	Network() string
}

type Peer struct {
	Name string
	Addr string
}

type Server struct {
	sync.WaitGroup

	Network string
	Port    int

	//Peers []*Peer
	Peers []*net.UDPAddr
}

func New() *Server {
	return &Server{
		Network: "udp",
		Port:    9092,
	}
}

func (s *Server) Run() error {
	sAddr := net.UDPAddr{IP: net.IPv4zero, Port: s.Port}
	fmt.Println("udp server running: ", sAddr.String())
	conn, err := net.ListenUDP(s.Network, &sAddr)
	if err != nil {
		fmt.Printf("ListenUDP err:%v, string:%v, network:%v ", err, sAddr.String(), sAddr.Network())
		return err
	}
	//defer conn.Close()

	data := make([]byte, 1024)
	// 监听，记录, 打洞
	s.Wrap(func() {
		for {
			//var b []byte
			n, peer, err := conn.ReadFromUDP(data)
			if err != nil {
				fmt.Printf("ReadFromUDP err:%v, n=%v\n", err, n)
				continue
			}

			fmt.Println("peer:", peer.String())
			// 检测 peer
			// add/update peer
			if err := s.AddPeer(peer); err != nil {
				fmt.Printf("AddPeer err:%v, ip=%v, port=%v\n", err, peer.IP, peer.Port)
				continue
			}

			// todo 打洞
			if len(s.Peers) == 2 {
				conn.WriteToUDP([]byte(s.Peers[0].String()), s.Peers[1])
				conn.WriteToUDP([]byte(s.Peers[1].String()), s.Peers[0])
				fmt.Println("make hold", s.Peers[0].String(), s.Peers[1].String())
				s.Peers = s.Peers[0:0]
			}
		}
	})

	return nil
}

func (s *Server) AddPeer(peer *net.UDPAddr) error {

	s.Peers = append(s.Peers, peer)

	return nil
}

func (s *Server) MakeHold() error {
	return nil
}

func (s *Server) Wrap(fc func()) {
	s.Add(1)
	go func() {
		defer s.Done()
		fc()
	}()
}

func main() {

	s := New()

	s.Run()

	s.Wait()
}
