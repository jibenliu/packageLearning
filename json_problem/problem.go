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

func (n *Nested) UnmarshalJSON(data []byte) error {
	*n = Nested{}
	tmp := struct {
		Dur string `json:"duration"`
	}{}
	fmt.Printf("parsing nested json %s \n", string(data))
	if err := json.Unmarshal(data, &tmp); err != nil {
		fmt.Printf("failed to parse nested: %v", err)
		return err
	}
	tmpDur, err := time.ParseDuration(tmp.Dur)
	if err != nil {
		fmt.Printf("failed to parse duration: %v", err)
		return err
	}
	(*n).Dur = tmpDur
	return nil
}

type Object struct {
	Nested
	Num int `json:"num"`
}

//uncommenting this method still doesnt help.
//tmp is parsed with the completed json at Nested
//which doesnt take care of Num field, so Num is zero value.
func (o *Object) UnmarshalJSON(data []byte) error {
	*o = Object{}
	tmp := struct {
		Nested
		Num int `json:"num"`
	}{}
	fmt.Printf("parsing object json %s \n", string(data))
	if err := json.Unmarshal(data, &tmp); err != nil {
		fmt.Printf("failed to parse object: %v", err)
		return err
	}
	fmt.Printf("tmp object: %+v \n", tmp)
	(*o).Num = tmp.Num
	(*o).Nested = tmp.Nested
	return nil
}

func main() {
	obj := Object{}
	if err := json.Unmarshal([]byte(testJSON), &obj); err != nil {
		fmt.Printf("failed to parse result: %v", err)
		return
	}
	fmt.Printf("result: %+v \n", obj)
}