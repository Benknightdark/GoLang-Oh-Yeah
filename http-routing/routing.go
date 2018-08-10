package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
	m string
}

func (p *MyMux) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}
func sayhelloName(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello aaa")

}
func main() {
	mux := &MyMux{}
	http.HandleFunc("/", mux.ServerHTTP)
	//	http.HandleFunc(mux.ServerHTTP())
	fmt.Println(mux)
	http.ListenAndServe(":8222", nil)
}
