package controller

import (
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		var response = map[string]string{
			"Messsage": "Method Not Allowed",
		}
		res, _ := json.Marshal(response)
		w.Write(res)
		return
	}
	w.WriteHeader(http.StatusOK)
	var response = map[string]string{
		"Messsage": "Hello World",
	}
	res, _ := json.Marshal(response)
	w.Write(res)
}
