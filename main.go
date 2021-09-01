package main

import (
	"fmt"
	"net/http"

	"github.com/el-zacharoo/go-101/handler"
	"github.com/el-zacharoo/go-101/store"
	"github.com/el-zacharoo/go-101/test"
)

func main() {

	s := store.Store{}
	s.Connect()

	p := handler.Person{
		Store: s,
	}
	o := handler.Org{
		Store: s,
	}

	http.HandleFunc("/person", p.Create)
	http.HandleFunc("/org", o.Create)
	http.HandleFunc("/test", test.Test)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Print(err)
	}
}
