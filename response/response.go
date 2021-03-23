package response

import (
	"encoding/json"
	"net/http"
)

func RenderJSON(w http.ResponseWriter, v interface{}) {
	var (
		js []byte
		err error
	)
	if js, err = json.Marshal(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(js); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
