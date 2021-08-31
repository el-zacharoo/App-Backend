package print

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DataStructure struct {
	Key    string `json:"Key"`
	Record struct {
		Name     string `json:"name"`
		Type     string `json:"type"`
		Validity bool   `json:"validity"`
	} `json:"Record"`
}

// var data string = `[{"Key":"area1", "Record": {"name":"belfast","type":"surburban","validity":true}},{"Key":"area1", "Record": {"name":"belfast","type":"surburban","validity":false}}]`

func Message(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	reqByt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err %v", err)))
	}
	var datastruct []DataStructure

	var result []DataStructure
	if err := json.Unmarshal([]byte(reqByt), &datastruct); err != nil {
		panic(err)
	}
	for _, item := range datastruct {
		if item.Record.Validity {
			result = append(result, item)
		}
	}
	fmt.Println(result)
}
