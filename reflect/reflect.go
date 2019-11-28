package main

import (
	"fmt"
	"reflect"
)

type data struct {
	Name  string `json:"name"`
	Age   uint   `json:"age"`
	Phone string `json:"phone"`
}

type config struct {
	Name   string      `json:"name"`
	Atr    interface{} `json:"atr"`
	Remark string      `json:"remark"`
}

type Person struct {
	Atr *data `json:"atr"`
}

func assign() {
	var man interface{}

	c := &config{
		Name: "123124",
		Atr: &data{
			Name:  "derek",
			Age:   23,
			Phone: "10086",
		},
	}
	manConfig := &Person{
		Atr: &data{
			Name:  "guin",
			Age:   235,
			Phone: "10010",
		},
	}
	man = manConfig

	ctype := reflect.TypeOf(c)
	cvalue := reflect.ValueOf(c)

	mtype := reflect.TypeOf(man)
	mvalue := reflect.ValueOf(man)

	for i := 0; i < ctype.Elem().NumField(); i++ {

		fmt.Println("c", ctype.Elem().Field(i), cvalue.Elem().Field(i))

		for j := 0; j < mtype.Elem().NumField(); j++ {

			fmt.Println("m", mtype.Elem().Field(j), mvalue.Elem().Field(j))

			if ctype.Elem().Field(i).Name == mtype.Elem().Field(j).Name {
				//c := ctype.Elem().Field(i)
				//m := mtype.Elem().Field(j)

				reflect.ValueOf(cvalue.Elem().Field(i))
				cvalue.Elem().Field(i).Set(mvalue.Elem().Field(j))
				//mvalue.Elem().Field(j).Set()
			}
		}
	}

	fmt.Println(c.Atr)
	fmt.Println(manConfig.Atr)
}

func main() {
	assign()

	return

	var man interface{}

	c := &config{
		Atr: &data{
			Name:  "derek",
			Age:   23,
			Phone: "10086",
		},
	}

	manConfig := &Person{
		Atr: &data{
			Name:  "guin",
			Age:   235,
			Phone: "10010",
		},
	}
	man = manConfig

	// todo: assign data -> person.atr

	fmt.Println(c.Atr)
	fmt.Println(manConfig.Atr)
	fmt.Println(reflect.TypeOf(c).Elem())

	etype := reflect.TypeOf(c).Elem()
	evalue := reflect.ValueOf(c.Atr)

	mtype := reflect.TypeOf(man).Elem()
	mvalue := reflect.ValueOf(man).Elem()

	// find m atr
	for i := 0; i < mtype.NumField(); i++ {
		for j := 0; j < etype.NumField(); j++ {

			fmt.Println(mtype.Field(i).Name, etype.Field(j).Name)

			if mtype.Field(i).Name == etype.Field(j).Name {
				// c.atr assign to m.atr
				mvalue.Field(i).Set(evalue)
				break
			}
		}

	}

	fmt.Println(c.Atr)
	fmt.Println(manConfig.Atr)
}
