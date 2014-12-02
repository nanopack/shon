package main

import "testing"

func TestParseJson(t *testing.T) {
  b := []byte(`{"b":true,"i":1234,"st":"hi friend","map":{"hi":"ono"},"arr":[1.4,2,3,"what"]}`)
  r := ParseJson(b)

  if r.(map[string]interface{})["b"].(bool) != true {
    t.Errorf("I could not parse json")
  }
}