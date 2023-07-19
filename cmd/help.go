package cmd

import (
	"fmt"

	"github.com/gophers-latam/minimal-chi/internal/provider"
)

func RunHelp() {
	appName := provider.GetContainer().Config().APPName

	fmt.Printf(`%s es la plantilla mínima para crear apis propuesta por la comunidad Gophers Latam usando al router chi.

	Comandos:

	Serve.   Sirve esta api.
	$ %s serve

	  Flags:

         (psst! Ud. debería definir aquí sus flags si su aplicación los necesita)

	Help. Muestra esta ayuda.
	$ cart Help.`, appName, appName)
}
