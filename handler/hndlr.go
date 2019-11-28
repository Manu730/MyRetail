package hndlr

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"test/db"

	"github.com/MyRetail/common"

	"github.com/gorilla/mux"
)

const (
	SUCCESS        = "SUCCESS"
	FAILURE        = "FAILURE"
	INTERNAL_ERROR = "INTERNAL_ERROR"
	NONE           = "NONE"
)

func ServeProducts(w http.ResponseWriter, r *http.Request) {
	var output GetProductOutput
	vars := mux.Vars(r)
	proId := vars["id"]
	prod, e := GetProduct(proId)
	if e != nil {
		fmt.Println("Error getting products", e)
		output.Result = FAILURE
		output.RespCode = INTERNAL_ERROR
		output.RespDescription = e.Error()
	} else {
		output.Result = SUCCESS
		output.RespCode = NONE
		output.RespDescription = NONE
		output.Product = *prod
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if output.Result == FAILURE {
		w.WriteHeader(422)
	} else {
		w.WriteHeader(200)
	}
	if err := json.NewEncoder(w).Encode(output); err != nil {
		fmt.Println("Error sending response", err)
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var input UpdateProdInput
	var output UpdateProdOutput
	vars := mux.Vars(r)
	proId := vars["id"]

	output.Result = FAILURE
	defer func() {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if output.Result == FAILURE {
			w.WriteHeader(422)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
		if err := json.NewEncoder(w).Encode(output); err != nil {
		}
	}()

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, common.MAX_REST_PAYLOAD))
	if err != nil {
		output.RespCode = "INVALID_INPUT"
		output.RespDescription = err.Error()
		return
	}

	if err = r.Body.Close(); err != nil {
		output.RespCode = "INVALID_INPUT"
		output.RespDescription = err.Error()
		return
	}

	if err = json.Unmarshal(body, &input); err != nil {
		output.RespCode = "INVALID_INPUT"
		output.RespDescription = err.Error()
		return
	}

	prod, e := db.GetProduct(proId)
	if e != nil {
		output.RespCode = INTERNAL_ERROR
		output.RespDescription = e.Error()
		return
	}
	if input.Product.Price.Value != prod.Price.Value {
		prod.Price.Value = input.Product.Price.Value
		if e = db.UpdateProduct(*prod); e != nil {
			output.RespCode = INTERNAL_ERROR
			output.RespDescription = e.Error()
			return
		}
	}

	output.Result = SUCCESS
	output.RespCode = NONE
	output.RespDescription = NONE

}
