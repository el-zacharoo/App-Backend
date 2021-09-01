package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/el-zacharoo/go-101/data"
	"github.com/el-zacharoo/go-101/store"
)

type Person struct {
	Store store.Store
}

func (p *Person) Create(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	reqByt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err %v", err)))
	}
	var psn data.Person
	json.Unmarshal(reqByt, &p)

	p.Store.AddPerson(psn)
	w.Write([]byte("done"))
}
