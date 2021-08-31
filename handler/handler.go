package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type person struct {
	ID         string `json:"id"`
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
	Email      string `json:"email"`
}

var people struct {
	NextCursor string   `json:"nextCursor"`
	Data       []person `json:"data"`
}

func PersonHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	reqByt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err %v", err)))
	}
	var p person
	json.Unmarshal(reqByt, &p)
	fmt.Println(p)

	rsp := person{
		GivenName:  p.GivenName,
		FamilyName: p.FamilyName,
		Email:      p.Email,
	}
	rspByt, err := json.Marshal(rsp)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err %v", err)))
	}
	w.Write(rspByt)

}

func PeopleHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	reqByt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err %v", err)))
	}
	var p person
	json.Unmarshal(reqByt, &p)
	fmt.Println(p)

	rsp := person{
		GivenName:  p.GivenName,
		FamilyName: p.FamilyName,
		Email:      p.Email,
	}
	rspByt, err := json.Marshal(rsp)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err %v", err)))
	}
	w.Write(rspByt)
}
