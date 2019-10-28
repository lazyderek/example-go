package service

import "fmt"

type Mainer struct {
}

func (m *Mainer) Run() error {
	fmt.Println("running main service ...")
	return nil
}

func (m *Mainer) Exit() {
	fmt.Println("exiting main service ...")
}
