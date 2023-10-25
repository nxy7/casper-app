package main

import (
	"fmt"
	"net/http"
	"nxy7/casper-app/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("start")
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	makeRoutes(r)

	err := http.ListenAndServe(":3333", r)
	if err != nil {
		fmt.Println(err)
	}
}

func makeRoutes(r *chi.Mux) {
	r.Get("/", routes.Hello)
	r.Get("/user/nfts", routes.Hello)
	r.Get("/queryCasper", routes.Hello)
}
