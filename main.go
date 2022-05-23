package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func main() {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}
	caller := conn.Object("com.deepin.utcloud.Daemon", "/com/deepin/utcloud/Daemon")
	caller.AddMatchSignal("org.freedesktop.DBus.Properties", "PropertiesChanged")

	go func() {
		c := make(chan *dbus.Signal, 10)
		conn.Signal(c)
		for v := range c {
			if len(v.Body) < 2 {
				continue
			}
			values, ok := v.Body[1].(map[string]dbus.Variant)
			if !ok {
				continue
			}

			v, ok := values["UserData"]
			if !ok {
				continue
			}
			fmt.Println(v)
		}
	}()
	select {}
}

type Foo struct {
	A int `tag1:"First Tag" tag2:"Second Tag"`
	B string
}

func SetMode() {
	sl := []int{1, 2, 3}
	greeting := "hello"
	greetingPtr := &greeting
	f := Foo{A: 10, B: "Salutations"}
	fp := &f

	slType := reflect.TypeOf(sl)
	gType := reflect.TypeOf(greeting)
	grpType := reflect.TypeOf(greetingPtr)
	fType := reflect.TypeOf(f)
	fpType := reflect.TypeOf(fp)

	examiner(slType, 0)
	examiner(gType, 0)
	examiner(grpType, 0)
	examiner(fType, 0)
	examiner(fpType, 0)
}

func examiner(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "Type is", t.Name(), "and kind is", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1), "Contained type:")
		examiner(t.Elem(), depth+1)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "name is", f.Name, "type is", f.Type.Name(), "and kind is", f.Type.Kind())
			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth+2), "Tag is", f.Tag)
				fmt.Println(strings.Repeat("\t", depth+2), "tag1 is", f.Tag.Get("tag1"), "tag2 is", f.Tag.Get("tag2"))
			}
		}
	}
}

func funcCaller() {
	timed := MakeTimedFunction(timeMe).(func())
	timed()
	timedToo := MakeTimedFunction(timeMeToo).(func(int) int)
	fmt.Println(timedToo(2))
}

func MakeTimedFunction(f interface{}) interface{} {
	rf := reflect.TypeOf(f)
	if rf.Kind() != reflect.Func {
		panic("expects a function")
	}
	vf := reflect.ValueOf(f)
	wrapperF := reflect.MakeFunc(rf, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := vf.Call(in)
		end := time.Now()
		fmt.Printf("calling %s took %v\n", runtime.FuncForPC(vf.Pointer()).Name(), end.Sub(start))
		return out
	})
	return wrapperF.Interface()
}

func timeMe() {
	fmt.Println("starting")
	time.Sleep(1 * time.Second)
	fmt.Println("ending")
}

func timeMeToo(a int) int {
	fmt.Println("starting")
	time.Sleep(time.Duration(a) * time.Second)
	result := a * 2
	fmt.Println("ending")
	return result
}
