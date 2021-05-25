package main

import (
	"encoding/json"
	"fmt"
	"time"
)

var testJSON = `{"num":5,"duration":"5s"}`

type Nested struct {
	Dur time.Duration `json:"duration"`
}

func (obj *Object) UnmarshalJSON(data []byte) error {
	tmp := struct {
		Dur string `json:"duration"`
		Num int    `json:"num"`
	}{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	dur, err := time.ParseDuration(tmp.Dur)
	if err != nil {
		return err
	}
	obj.Dur = dur
	obj.Num = tmp.Num
	return nil
}

type Object struct {
	Nested
	Num int `json:"num"`
}

var _ json.Unmarshaler = (*Object)(nil)

func main() {
	obj := Object{}
	_ = json.Unmarshal([]byte(testJSON), &obj)
	fmt.Printf("result: %+v \n", obj)
}
