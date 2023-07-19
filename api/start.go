package api

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/arl/statsviz"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	v1 "github.com/gophers-latam/minimal-chi/api/v1"
	"github.com/gophers-latam/minimal-chi/internal/provider"
)

// Start inicializa los routers y subrouters, cargando los endpoints registrados en ellos
// e inicializa el servicio de la api.
func Start() {
	r := chi.NewRouter()

	// middlewares sett up.  Revise la documentación de chi sobre middlewares en: https://go-chi.io/#/pages/middleware
	r.Use(middleware.Logger) // <--<< Logger should come before Recoverer
	r.Use(middleware.Recoverer)

	r.Get("/debug/statsviz/ws", statsviz.Ws)
	r.Get("/debug/statsviz", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/debug/statsviz/", http.StatusMovedPermanently)
	})
	r.Handle("/debug/statsviz/*", statsviz.Index)

	r.HandleFunc("/pprof", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, r.RequestURI+"/", http.StatusMovedPermanently)
	})

	r.HandleFunc("/pprof/*", pprof.Index)
	r.HandleFunc("/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/pprof/profile", pprof.Profile)
	r.HandleFunc("/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/pprof/trace", pprof.Trace)

	r.Handle("/pprof/goroutine", pprof.Handler("goroutine"))
	r.Handle("/pprof/threadcreate", pprof.Handler("threadcreate"))
	r.Handle("/pprof/mutex", pprof.Handler("mutex"))
	r.Handle("/pprof/heap", pprof.Handler("heap"))
	r.Handle("/pprof/block", pprof.Handler("block"))
	r.Handle("/pprof/allocs", pprof.Handler("allocs"))

	// Vaya a v1/routes para registrar sus rutas. Esto permitirá que cuando lo necesite pueda escalar a una v2 de su api con bajo costo
	v1.Routes(r)

	fmt.Printf("Escuchando en puerto %d\n", provider.GetContainer().Config().Port)
	// Empieza la magia!

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", provider.GetContainer().Config().Port), r))
}
