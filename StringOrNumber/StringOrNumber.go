package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

var rawJson = []byte(`[
  {
    "id": 1,
    "address": {
      "city_id": 5,
      "street": "Satbayev"
    },
    "Age": 20
  },
  {
    "id": 1,
    "address": {
      "city_id": "6",
      "street": "Al-Farabi"
    },
    "Age": "32"
  }
]`)

type IntString int

func (st *IntString) UnmarshalJSON(b []byte) error {
	var item interface{}
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	switch v := item.(type) {
	case int:
		*st = IntString(v)
	case float64:
		*st = IntString(int(v))
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*st = IntString(i)
	}
	return nil
}

func main() {

	var users []User
	err := json.Unmarshal(rawJson, &users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", users)

}

type User struct {
	XMLName   xml.Name `xml:"user" json:"-"`
	ID      IntString `xml:"id" json:"id"`
	Address Address   `xml:"address" json:"address,omitempty"`
	Age     IntString `xml:"age" json:"age"`
}

type Address struct {
	CityID IntString `xml:"city_id" json:"city_id"`
	Street string    `xml:"street" json:"street"`
}
