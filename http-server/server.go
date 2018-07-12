package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	for k, v := range r.Form {
		fmt.Println(k, strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello aaa")

}
func main() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":8388", nil)
	if err != nil {
		log.Fatal(err)
	}
}
