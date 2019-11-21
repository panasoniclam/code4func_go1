package main

import "fmt"

func Tong(a int, b int) int  {
	return  a+b

}

func Mul(a int,b int) int{
	return  a*b
}

func Div(a int, b int) float64  {
	return  float64(a/b)
}

func Chuvi(a,b,h int) float64  {
	tong := Tong(a,b)
	tich := Mul(tong,h)
	result:=  float64(tich/2)
	return  result
}
func main()  {
	 fmt.Println(Chuvi(2,3,4))
}