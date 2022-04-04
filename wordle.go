package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/chloesgit/wordle/lib"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{}))
	motcible := "Bonjour"
	motcible = strings.ToLower(motcible)
	r.Get("/{mot}", func(w http.ResponseWriter, r *http.Request) {
		motsaisi := chi.URLParam(r, "mot")
		motsaisi = strings.ToLower(motsaisi)
		if motcible == motsaisi {
			w.Write([]byte("Vous avez gagné"))
			return
		}
		lettrescorrectes, err := lib.Isword(motcible, motsaisi)
		if err != nil {
			http.Error(w, "7", 500)
			return
		}
		reponse := make(map[string]lib.InfoLettre)
		for i, _ := range lettrescorrectes {
			reponse[string(motsaisi[i])] = lettrescorrectes[i]
		}
		json.NewEncoder(w).Encode(reponse)

	})
	fmt.Println("http://localhost:8000")
	http.ListenAndServe(":8000", r)

}
