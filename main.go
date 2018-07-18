package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type TheTime struct {
	DateTime *time.Time `json:"dateTime"`
}

func main() {
	now := time.Now()
	a := TheTime{DateTime: &now}
	b, _ := json.Marshal(a)
	fmt.Println(string(b))
	c := TheTime{DateTime: nil}
	d, _ := json.Marshal(c)
	fmt.Println(string(d))
}
