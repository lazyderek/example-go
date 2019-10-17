package singleton

import (
	"fmt"
	"testing"
)

func TestGetUser(t *testing.T) {
	user := GetUser()
	fmt.Println(user.Username)
}

func TestGetSingleUser(t *testing.T) {
	user := GetSingleUser()
	fmt.Printf("%x\n", user.Username)

	user = GetSingleUser()
	fmt.Printf("%x\n", user.Username)
}
