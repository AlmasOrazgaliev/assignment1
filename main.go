package main

import (
	"fmt"
	"github.com/AlmasOrazgaliev/assignment1/apiserver"
	"github.com/AlmasOrazgaliev/assignment1/model"
	"log"
)

func main() {
	u := model.User{
		Email:    "email@example.org",
		Password: "password",
	}
	fmt.Println(u.BeforeCreate(), u)

	config := apiserver.NewConfig()
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
