package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/el-zacharoo/go-101/data"
	"github.com/el-zacharoo/go-101/store"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
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

	if r.Method == http.MethodOptions {
		return
	}

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

	if r.Method == http.MethodOptions {
		return
	}
	// parts := strings.Split(r.RequestURI, "/")
	// id := parts[len(parts)-1]
	id := chi.URLParam(r, "id")

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

// func (p *Person) Query(w http.ResponseWriter, r *http.Request) {

// 	if r.Method == http.MethodOptions {
// 		return
// 	}

// 	// defer res.Body.Close()
// 	// byt, err := ioutil.ReadAll(res.Body)
// 	// if err != nil {
// 	// 	return &pb.QueryResponse{}, err
// 	// }
// 	// if res.StatusCode != http.StatusOK {
// 	// 	return &pb.QueryResponse{}, fmt.Errorf("an error occurred fetching users: %s", string(byt))
// 	// }

// 	//  psn := data.Person

// 	var psn data.People

// 	psn, err := p.Store.GetPeople()
// 	reqByt, err := ioutil.ReadAll(r.Body)

// 	if err = json.Unmarshal(reqByt); err != nil {
// 		return &pb.QueryResponse{}, err
// 	}

// 	for _, c := range psn.Data {
// 		if len(c.Credential) == 0 {
// 			continue
// 		}
// 		var cred pb.Credential
// 		if err = protojson.Unmarshal(c.Credential, &cred); err != nil {
// 			return &pb.QueryResponse{}, err
// 		}
// 		cred.Id = c.ID

// 		qrsp.Cursor = append(qrsp.Cursor, &cred)
// 	}

// 	rspByt, err := json.Marshal(psn)
// 	if err != nil {
// 		w.Write([]byte(fmt.Sprintf("error %v", err)))
// 	}

// 	w.Write(rspByt)
// }

func (p *Person) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	reqByt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err %v", err)))
	}
	var psn data.Person
	json.Unmarshal(reqByt, &psn)

	id := chi.URLParam(r, "id")

	p.Store.UpdatePerson(id, psn)
	w.Write([]byte("done"))
}

func (p *Person) Delete(w http.ResponseWriter, r *http.Request) {

	// parts := strings.Split(r.RequestURI, "/")
	// id := parts[len(parts)-1]
	id := chi.URLParam(r, "id")

	if err := p.Store.DeleteUser(id); err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}
	w.Write([]byte("done"))
}
