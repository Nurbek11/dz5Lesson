package main

import (
	"reflect"
	"testing"
	"time"
)

func Cyrillic_test(t *testing.T) {
	p := Person{
		Name:      "NurbekГОДЕВ",
		Surname:   "Nessipbay",
		Age:       21,
		Birthdate: time.Time{},
	}
	b := []string{"Nurbek", "Nessipbay"}
	ans := deleteCyrillic(&p)
	if reflect.DeepEqual(ans, b) {
		t.Error("There are cyrillic symbols")
	}

}
