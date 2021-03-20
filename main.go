package main

import (
	"fmt"

	srv "github.com/josiahsams/demo1/service"
	u "github.com/josiahsams/demo1/utils"
)


func main() {
	sv := srv.NewService(u.Utils{})
	v, _ := sv.Method1(5)

	fmt.Println("hello" , v)
}

