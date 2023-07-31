package cmd

import (
	"fmt"
)

func RunHelp() {

	fmt.Printf(`Es la plantilla mínima para crear apis propuesta por la comunidad Gophers Latam usando al router chi.

Comandos:

serve - Corre esta API.

Ejemplo
$go run ./main.go serve

help - Muestra esta ayuda.

Ejemplo
$go run ./main.go help

Flags:

(Ud. debería definir aquí sus flags si su aplicación los necesita)
`)
}
