package handlers

import (
	"encoding/json"
	"net/http"
)

// Ping es el manejador para el EP v1/ping
func Ping(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}{
		Code: 200,
		Msg:  "Bienvenido a la plantilla minimal Chi de Gophers Latam",
	}

	js, err := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		_, _ = w.Write([]byte(`{"code": 500, msg:"couldnt serialize response"}`))
		return
	}

	_, _ = w.Write(js)
}
