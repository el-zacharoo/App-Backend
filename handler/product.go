package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/el-zacharoo/go-101/data"
	"github.com/el-zacharoo/go-101/store"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Product struct {
	Store store.Store
}

// func (p *Person) Handle(w http.ResponseWriter, r *http.Request) {

// 	switch r.Method {
// 	case http.MethodPost:
// 		p.create(w, r)
// 	case http.MethodDelete:
// 		p.delete(w, r)
// 	}

// }

func (prd *Product) Create(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	reqByt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err %v", err)))
	}
	var prod data.Product
	json.Unmarshal(reqByt, &prod)

	prod.ID = uuid.New().String()
	prd.Store.AddProduct(prod)
	w.Write([]byte("done"))
}

func (prd *Product) Get(w http.ResponseWriter, r *http.Request) {

	// parts := strings.Split(r.RequestURI, "/")
	// id := parts[len(parts)-1]
	id := mux.Vars(r)["id"]

	prod, err := prd.Store.GetProduct(id)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}

	rspByt, err := json.Marshal(prod)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}

	w.Write(rspByt)
}

func (prd *Product) Update(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	reqByt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err %v", err)))
	}
	var prod data.Product
	json.Unmarshal(reqByt, &prod)

	id := mux.Vars(r)["id"]

	prd.Store.UpdateProduct(id, prod)
	w.Write([]byte("done"))
}

func (prd *Product) Delete(w http.ResponseWriter, r *http.Request) {

	// parts := strings.Split(r.RequestURI, "/")
	// id := parts[len(parts)-1]
	id := mux.Vars(r)["id"]

	if err := prd.Store.DeleteProduct(id); err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}
	w.Write([]byte("done"))
}
