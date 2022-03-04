package main

import (
	"encoding/json"
	"fmt"

	smodel "github.com/adhimaswaskita/go-shared-lib/models"
)

func main() {
	user := register("Adhimas", 21)
	u, _ := json.Marshal(user)
	fmt.Println(string(u))
}

func register(name string, age uint8) smodel.User {
	user := smodel.User{
		Name: name,
		Age:  age,
	}
	return user
}
