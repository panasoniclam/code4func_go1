package main
import "fmt"

func main(){
	//var num int
	//var float float64
	//var str string

	//var a [5] interface{}
	// khai bao keiu arry, mac dinh chieu dai array
	//a.append(a,2,3)
	//a[1] =4
	//a[4] =5
	//fmt.Println(len(a),a)
	//var e[0] string
	//fmt.Println(cap(e),e)

	//hello( "hello")
	// a:=map[string] int{} // tap hop key, value
	//b:= make(map[string]int)
	//var c map[string]int
	//encode(a)
	//encode(b)
	//encode(c)
	// var c map[string]interface{}
	// c = make(map[string]interface{})
	//c["key1"] =2
	//
	//  fmt.Println(c)


	myMap := make(map[string]int)
	myMap["key1"] = 123
	myMap["key2"] = 456
	fmt.Println(myMap)

	value, ok := myMap["key1"]
	if ok {
		//found
		fmt.Println(value)
		fmt.Println(ok)
	}

	//func encode(in map[string]int)  {
	//	d,_ := json.Marshal(in) //
	//	fmt.Println(string((d)))
	//}