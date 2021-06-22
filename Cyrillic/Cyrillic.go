package main

import (
	"reflect"
	"regexp"
	"time"
)

func main() {

	p := Person{
		Name:      "NurbekГОДЕВ",
		Surname:   "Nessipbay",
		Age:       21,
		Birthdate: time.Time{},
	}
	deleteCyrillic(&p)
}

type Person struct {
	Name      string
	Surname   string
	Age       int
	Birthdate time.Time
}

func deleteCyrillic(s *Person) []string {
	v := reflect.ValueOf(*s)

	var array []string

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Type().String() == "string" {
			re := regexp.MustCompile("[[:^ascii:]]")
			var a = re.ReplaceAllLiteralString(v.Field(i).String(), "")
			array =  append(array,a)
		}
	}
	return array
}
