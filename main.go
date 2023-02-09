package main

import (
	"fmt"
	"github.com/AlmasOrazgaliev/assignment1/model"
)

func main() {
	u := model.User{
		Email:    "email@example.org",
		Password: "password",
	}
	fmt.Println(u.BeforeCreate(), u)
}
