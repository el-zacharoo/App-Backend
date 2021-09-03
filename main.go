package main

import (
	"fmt"
	"net/http"

	"github.com/el-zacharoo/go-101/handler"
	"github.com/el-zacharoo/go-101/store"
	"github.com/gorilla/mux"
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

	prd := handler.Product{
		Store: s,
	}

	r := mux.NewRouter()

	people := r.PathPrefix("/people").Subrouter()
	// people.Path("").Methods(http.MethodGet).HandlerFunc(p.Query)
	people.Path("").Methods(http.MethodPost).HandlerFunc(p.Create)
	people.Path("/{id}").Methods(http.MethodGet).HandlerFunc(p.Get)
	people.Path("/{id}").Methods(http.MethodPut).HandlerFunc(p.Update)
	people.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(p.Delete)

	org := r.PathPrefix("/org").Subrouter()
	org.Path("").Methods(http.MethodPost).HandlerFunc(o.Create)
	org.Path("/{id}").Methods(http.MethodGet).HandlerFunc(o.Get)
	org.Path("/{id}").Methods(http.MethodPut).HandlerFunc(o.Update)
	org.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(o.Delete)

	product := r.PathPrefix("/product").Subrouter()
	product.Path("").Methods(http.MethodPost).HandlerFunc(prd.Create)
	product.Path("/{id}").Methods(http.MethodGet).HandlerFunc(prd.Get)
	product.Path("/{id}").Methods(http.MethodPut).HandlerFunc(prd.Update)
	product.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(prd.Delete)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Print(err)
	}
}
