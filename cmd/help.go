package cmd

import (
	"fmt"

	"github.com/gophers-latam/minimal-chi/internal/provider"
)

func RunHelp() {
	appName := provider.GetContainer().Config().APPName

	fmt.Printf(`%s es la plantilla mínima para crear apis propuesta por la comunidad Gophers Latam usando al router chi.

Comandos:

serve - Corre esta API.

Ejemplo
$go run ./main.go serve

help - Muestra esta ayuda.

Ejemplo
$go run ./main.go help

Flags:

(Ud. debería definir aquí sus flags si su aplicación los necesita)
`, appName)
}
