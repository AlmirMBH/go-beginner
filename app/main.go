package main

import (
	stuff "example/project/mypackage"
	"fmt"
	"reflect"
)


func main(){
	fmt.Println("Hello", stuff.Name)
	intArr := []int{1, 2, 3, 4, 5}
	fmt.Println("Array", stuff.IntArrToStrArr(intArr))
	fmt.Println("Type of:", reflect.TypeOf(stuff.IntArrToStrArr(intArr)))
}