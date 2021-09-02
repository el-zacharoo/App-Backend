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

type Person struct {
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

func (p *Person) Create(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	reqByt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err %v", err)))
	}
	var psn data.Person
	json.Unmarshal(reqByt, &psn)

	psn.ID = uuid.New().String()
	p.Store.AddPerson(psn)
	w.Write([]byte("done"))
}

func (p *Person) Get(w http.ResponseWriter, r *http.Request) {

	// parts := strings.Split(r.RequestURI, "/")
	// id := parts[len(parts)-1]
	id := mux.Vars(r)["id"]

	psn, err := p.Store.GetPerson(id)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}

	rspByt, err := json.Marshal(psn)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}

	w.Write(rspByt)
}

func (p *Person) Update(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	reqByt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err %v", err)))
	}
	var psn data.Person
	json.Unmarshal(reqByt, &psn)

	id := mux.Vars(r)["id"]

	p.Store.UpdatePerson(id, psn)
	w.Write([]byte("done"))
}

func (p *Person) Delete(w http.ResponseWriter, r *http.Request) {

	// parts := strings.Split(r.RequestURI, "/")
	// id := parts[len(parts)-1]
	id := mux.Vars(r)["id"]

	if err := p.Store.DeleteUser(id); err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}
	w.Write([]byte("done"))
}
