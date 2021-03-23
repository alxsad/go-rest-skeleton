package handler

import (
	"github.com/alxsad/go-rest-skeleton/response"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response.RenderJSON(w, map[string]string{
		"status": "ok",
	})
}

func PanicHandler(w http.ResponseWriter, r *http.Request) {
	panic("panic")
}
