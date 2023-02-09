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
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
