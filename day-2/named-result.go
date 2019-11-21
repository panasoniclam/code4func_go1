package main

import "fmt"
//result fucntion tra ve mot cai ten da dc dinh ngia truoc
func split(sum int)(x,y int)  {
	x = sum+1
	y= x-2
	return
}
func main()  {
	fmt.Println(split(10))
}