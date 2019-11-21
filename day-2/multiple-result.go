package main

import "fmt"
// function rea ve nhieu ket qua
func swap(a ,b string) (string,string)  {
	return b,a
}
func main()  {
	a,b := swap("hello","word")
	fmt.Println(a,b)
}
