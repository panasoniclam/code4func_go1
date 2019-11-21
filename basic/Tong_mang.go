package main


func  sum(arr [] int) int  {
	result:= 0
	for i:=1;i<len(arr);i++{
		result+=arr[i]
	}
	return  result
}

