package provider

import "github.com/gophers-latam/minimal-chi/internal/config"

type Containerable interface {
	Config() config.Conf
}

// container es una implementación naif de un contenedor de dependencias para ser usdo rápidamente
// Si lo necesita agregue las dependencias de su aplicación como campos del contenedor.
type container struct {
	// cfg provee las dependencias de la aplicación de manera inmutable
	cfg config.Conf
}

// Config devuelve la configuración de la app cargada en el contenedor de dependencias
func (c *container) Config() config.Conf {
	return c.cfg
}

var (
	// cont es la instancia del contenedor de dependencias
	cont Containerable

	// este idioma permite validar en tiempó de compilación que la struct implementa la interface
	_ Containerable = &container{}
)

// GetContainer retorna al contenedor de dependencias. Devuelve un puntero para garantizar
// que en todas partes se tenga acceso a las mismas dependencias contenidas
func GetContainer() Containerable {
	if cont == nil {
		cont = &container{
			cfg: config.NewConf(),
		}
	}

	return cont
}
