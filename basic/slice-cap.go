package main

import "fmt"

func main()  {
	s:= []int{1,2,3,4,5,6}
	s = s[:1]
	printSlice(s)
}
func printSlice(s []int)  {
	fmt.Println("len = %d cap =%d %v\n",len(s),cap(s),s)
}