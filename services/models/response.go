package models

import (
	"net/http"
	"encoding/json"
)

type Response struct {
	Ok     bool
	Data interface{}
	Message string
	TokenChanged bool
	Token string
}

func (r *Response) SendJson(w http.ResponseWriter)  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	val, _ := json.Marshal(r)
	w.Write(val)

}