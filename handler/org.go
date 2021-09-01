package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/el-zacharoo/go-101/data"
	"github.com/el-zacharoo/go-101/store"
)

type Org struct {
	Store store.Store
}

func (o *Org) Create(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	reqByt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err %v", err)))
	}
	var org data.Org
	json.Unmarshal(reqByt, &o)

	o.Store.AddOrg(org)
	w.Write([]byte("done"))
}
