package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello World")
	//http://serveripaddress:8080/home
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")

	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello User")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "HOme Page")
	})

	http.ListenAndServe(":8080", nil)
}
