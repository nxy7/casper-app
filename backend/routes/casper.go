package routes

import "net/http"

func CasperQuery(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
