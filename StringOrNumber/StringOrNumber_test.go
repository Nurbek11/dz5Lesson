package main

import (
	"encoding/json"
	"testing"
)

func TestStringOrNumber(t *testing.T)  {
	var users []User
	err := json.Unmarshal(rawJson, &users)
	if err != nil {
		t.Error("something is wrong with decoding")
	}

}