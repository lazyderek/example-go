package main

import (
	"fmt"
	"github.com/influxdata/toml/ast"
	"io/ioutil"
	"os"

	btoml "github.com/BurntSushi/toml"
	"github.com/influxdata/toml"
)

type Person struct {
	Name string `toml:name`
	Age  int    `toml:age`
}

type Game struct {
	Name    string "toml:name"
	Size    uint   "toml:size"
	Company string "toml:company"
}

type Product struct {
	Name   string   "toml:name"
	Phones []string "toml:phones"
	//Sub    *subProduct "toml:sub"
}

type subProduct struct {
	Name string
}

type tomlConfig struct {
	Person  *Person
	Game    *Game
	Product *Product
}

func getConfigFile() string {
	dir, _ := os.Getwd()
	return dir + "/example.toml"
}

func loadConfig(file string) ([]byte, error) {
	fmt.Println(file)

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}

func main() {

	data, err := loadConfig(getConfigFile())
	if err != nil {
		panic("loadConfig: " + err.Error())
	}

	// case 1
	var tc tomlConfig
	if err := toml.Unmarshal(data, &tc); err != nil {
		panic("Unmarshal: " + err.Error())
	}

	fmt.Printf("person: %v\n", *tc.Person)
	fmt.Printf("game: %v\n", *tc.Game)
	fmt.Printf("product: %v\n", *tc.Product)

	// case 2
	t, err := toml.Parse(data)
	v := make(map[string]interface{})

	for _, field := range t.Fields {
		//if tbl, ok := t.Fields[name]; ok {
		if err := toml.UnmarshalTable(field.(*ast.Table), v); err != nil {
			panic("UnmarshalTable 1: " + err.Error())
		}
		//}
	}
	fmt.Println(v)

	// case 3
	var config tomlConfig
	fmt.Println(config.Person, config.Product, config.Game)
	var vc interface{}
	vc = &config
	if err := toml.UnmarshalTable(t, vc); err != nil {
		panic("UnmarshalTable 2: " + err.Error())
	}
	fmt.Println(config.Person, config.Product, config.Game)
}

// --------- BurntSushi toml --------------
func ParseToml() error {
	// config file -> parse -> config struct

	data, err := loadConfig(getConfigFile())
	if err != nil {
		panic("loadConfig: " + err.Error())
	}

	var (
		config tomlConfig
		vc     interface{}
	)

	vc = &config
	if err := btoml.Unmarshal(data, vc); err != nil {
		return err
	}
	fmt.Println(config.Person, config.Product, config.Game)
	return nil
}
