package main

import (
	"fmt"
	"reflect"
	"regexp"
)

func main() {

	g := Foo{Bar: "NurbekГОДЕВ"}
	deleteCyrillic(&g)
	fmt.Println(g)
}

type Foo struct {
	Id, Bar string
}

func deleteCyrillic(data interface{}) {
	val := reflect.ValueOf(data).Elem()

	v := reflect.ValueOf(data)
	i := reflect.Indirect(v)
	t := i.Type()

	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			if (t.Field(i).Type).String() == "string" {
				re := regexp.MustCompile("[[:^ascii:]]")
				var a = re.ReplaceAllLiteralString(val.Field(i).String(), "")
				val.FieldByName(t.Field(i).Name).SetString(a)
			}
		}
	} else {
		fmt.Println("not a stuct")
	}
}
