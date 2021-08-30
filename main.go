package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Hello")
	http.HandleFunc("/", MyHandler)
	http.ListenAndServe(":8080", nil)
}
func MyHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	byt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err %v", err)))
	}
	// w.Write([]byte(fmt.Sprintf("hello %s", r.Method)))
	w.Write(byt)
}
