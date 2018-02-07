package utils

import (
	"encoding/json"
	"net/http"
)

func Write(w http.ResponseWriter, source interface{}) {

	js, err := json.Marshal(source)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func Read(r *http.Request, destination interface{}) {
	err := json.NewDecoder(r.Body).Decode(&destination)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()
}
