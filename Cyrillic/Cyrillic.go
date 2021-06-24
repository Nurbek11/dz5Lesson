package main

import (
	"fmt"
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
	fmt.Println(deleteCyrillic(&p))
	fmt.Println(p)
}

type Person struct {
	Name      string
	Surname   string
	Age       int
	Birthdate time.Time
}

func deleteCyrillic(s *Person) error {
	v := reflect.ValueOf(*s)

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Type().String() == "string" {
			re := regexp.MustCompile("[[:^ascii:]]")
			var a = re.ReplaceAllLiteralString(v.Field(i).String(), "")
			switch i {
			case 0:
				s.Name = a
				break
			case 1:
				s.Surname = a
				break
			default:
				continue
			}

			//Tried but FAILED
			//fieldValue := reflect.ValueOf(v.Field(i))
			//if fieldValue.CanSet() {
			//	fieldValue.SetString(a)
			//}

		}

	}
	return nil
}
