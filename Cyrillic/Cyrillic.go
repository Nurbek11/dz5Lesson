package main

import (
	"fmt"
	"reflect"
	"regexp"
	"time"
)

type Person interface {
	change()
}

type John struct {
	Name      string
	Surname   string
	Age       int
	Birthdate time.Time
}
type Ben struct {
	Name string
}

func deleteCyrillic(person Person) {
	person.change()
}

func (c *John) change() {
	v := reflect.ValueOf(*c)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Kind() == reflect.String {
			re := regexp.MustCompile("[[:^ascii:]]")
			var a = re.ReplaceAllLiteralString(v.Field(i).String(), "")
			switch i {
			case 0:
				c.Name = a
				break
			case 1:
				c.Surname = a
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
}
func (c *Ben) change() {
	v := reflect.ValueOf(*c)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Kind() == reflect.String {
			re := regexp.MustCompile("[[:^ascii:]]")
			var a = re.ReplaceAllLiteralString(v.Field(i).String(), "")
			switch i {
			case 0:
				c.Name = a
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
}

func main() {

	john := John{
		Name:      "NurbekГОДЕВ",
		Surname:   "Nessipbay",
		Age:       21,
		Birthdate: time.Time{},
	}

	ben := Ben{
		Name: "Ben ПЫХЕР",
	}

	deleteCyrillic(&john)
	deleteCyrillic(&ben)

	fmt.Println(john, ben)

}
