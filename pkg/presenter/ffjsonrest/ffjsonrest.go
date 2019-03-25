package ffjsonrest

import (
	"encoding/json"
	"github.com/pquerna/ffjson/ffjson"
	"net/http"
	"presenters-benchmark/pkg/presenter"
)

type Candidate struct{}

var _ presenter.CandidateHandlerFunc = (*Candidate)(nil)

func (Candidate) HandlerFunc(options []*presenter.Option) (http.HandlerFunc, error) {
	return func(w http.ResponseWriter, r *http.Request) {
		var params struct {
			Query         string                 `json:"query"`
			OperationName string                 `json:"operationName"`
			Variables     map[string]interface{} `json:"variables"`
		}
		// mandatory to check this in all example

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response := NewResponse(options)

		//fmt.Println(response)
		// deserialize and return option

		//w.Write(prefix)
	//	w.68745267
		err :=  ffjson.NewEncoder(w).Encode(response) //json.NewEncoder(w).Encode(options)
	///	w.Write(sufix)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}, nil
}

func (Candidate) UnmarshalOptions(b []byte) ([]*presenter.Option, error) {
	return presenter.JSONUnmarshalOptions(b)
}

var prefix = []byte(`{"data": {"hotelX": {"search": {"options": `)
var sufix = []byte(`,"errors": {"code": "","type": "","description": ""}}}}}`)
