package main

import (
	"fmt"
	"reflect"
)

type model struct {
	firstField  int
	secondField string
}

func main() {
	var b = model{12, "af"}
	cyrillic(b)
	//BindStruct(&b)
}

func cyrillic(a interface{}) {
	//e := reflect.ValueOf(&a).Elem()
	val := reflect.ValueOf(a)
	//fmt.Println(e)

	for i := 0; i < val.NumField(); i++ {
		//valueField := val.Field(i)
		//typeField := val.Type().Field(i).Name
		fmt.Println(val.Field(i))
		//fmt.Println(typeField)
		//val.Field(i) =
	}

}
