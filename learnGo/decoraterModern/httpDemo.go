package main

import (
	"fmt"
	"log"
	"net/http"
)

func isAuthorized(endPoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("checking to see if Authorized header set ...")
		//t := reflect.TypeOf(r)
		//v := reflect.ValueOf(r)
		//for k := 0; k < t.NumField(); k++ {
		//	fmt.Printf("%s -- %v \n", t.Field(k).Name, v.Field(k).Interface())
		//	//当结构体中含有非导出字段(即私有属性)时，v.Field(k).Interface()会panic
		//}
		if val, ok := r.Header["Authorized"]; ok {
			fmt.Println(val)
			if val[0] == "true" {
				fmt.Println("Header is set! We can serve content!")
				endPoint(w, r)
			}
		} else {
			fmt.Println("Not Authorized!!")
			_, _ = fmt.Fprintf(w, "Not Authorized!!")
		}
	})
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EndPoint Hit: homePage")
	_, _ = fmt.Fprintf(w, "Welcome to the homePage")
}
func newEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("My New Endpoint")
	_, _ = fmt.Fprintf(w, "My second endpoint")
}

func handleRequests() {
	http.Handle("/", isAuthorized(homePage))
	http.Handle("/new", isAuthorized(newEndpoint))
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
