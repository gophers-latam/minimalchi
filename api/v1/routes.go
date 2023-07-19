package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/gophers-latam/minimal-chi/api/v1/handlers"
)

// Routes registra las rutas de los endpoints bajo v1 y las monta en el router principal
// recibido como argumento.
func Routes(r *chi.Mux) {
	// para agregar el tramo v1 en las ruta de dicha versión
	v1 := chi.NewRouter()

	v1.Get("/ping", handlers.Ping)

	r.Mount("/v1", v1)
}
