package main

import (
	"encoding/json"
	"fmt"
	"time"
)

var testJSON = `{"num":5,"duration":"5s"}`

type customTimeDuration time.Duration

type Nested struct {
	Dur customTimeDuration `json:"duration"`
}

func (ctd *customTimeDuration) UnmarshalJSON(b []byte) error {
	var durStr string
	if err := json.Unmarshal(b, &durStr); err != nil {
		return err
	}
	dur, err := time.ParseDuration(durStr)
	if err == nil {
		*ctd = customTimeDuration(dur)
	}
	return err
}

type Object struct {
	Nested
	Num int `json:"num"`
}

func main() {
	obj := Object{}
	_ = json.Unmarshal([]byte(testJSON), &obj)
	fmt.Printf("result: %+v \n", obj)
}
