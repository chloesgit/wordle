package main

import (
	"fmt"
	"net/http"
	"recette/wordle/lib"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	motcible := "Bonjour"
	motcible = strings.ToLower(motcible)
	r.Get("/{mot}", func(w http.ResponseWriter, r *http.Request) {
		motsaisi := chi.URLParam(r, "mot")
		motsaisi = strings.ToLower(motsaisi)
		if motcible == motsaisi {
			w.Write([]byte("Vous avez gagné"))
			return
		}
		lettrescorrectes, _ := lib.Isword(motcible, motsaisi)
		message := ""
		for i, _ := range lettrescorrectes {
			message = message + string(motsaisi[i]) + " est " + string(lettrescorrectes[i]) + "\n"
		}
		w.Write([]byte(message))

	})
	fmt.Println("http://localhost:3000")
	http.ListenAndServe(":3000", r)

}
