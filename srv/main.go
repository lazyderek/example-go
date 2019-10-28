package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"syscall"

	"github.com/judwhite/go-svc/svc"
	"guin/example-go/srv/service"
)

// this is a simple service
// 参考nsq的代码

type Service struct {
	once sync.Once
	main *service.Mainer
}

func main() {
	s := &Service{}
	if err := svc.Run(s, syscall.SIGINT, syscall.SIGTERM); err != nil {
		fmt.Printf("%v\n", err)
	}
}

func (s *Service) Init(env svc.Environment) error {
	fmt.Println("initing ...")

	// init log/db/service ...
	s.main = &service.Mainer{}

	return nil
}

func (s *Service) Start() error {
	// parse options
	flag.Parse()

	// start your main service
	go func() {
		if err := s.main.Run(); err != nil {
			s.Stop()
			os.Exit(1)
		}
	}()

	return nil
}

func (s *Service) Stop() error {
	s.once.Do(func() {
		// exit your s.main
		s.main.Exit()
	})
	return nil
}
