package main

import (
	"fmt"
	"github.com/panasoniclam/code4func_go1/day-2/child"
)
type (
	User struct{
		Id int
		Name string
		Password string
	}
)

func (u User) Getname() string{
	return "hello My name" + u.Name
}
func main()  {
    dog:= child.Dog{
    	id:0
    	Name:"ddd"
	}
}