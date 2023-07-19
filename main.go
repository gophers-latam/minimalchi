package main

import (
	"fmt"
	"os"
	"strings"
	_ "time/tzdata"

	"github.com/gophers-latam/minimal-chi/cmd"
	"github.com/gophers-latam/minimal-chi/internal/provider"
)

var (
	// BinVersion guarda la versión actual del binario de esta aplicación.
	// Normalmente la versión es el último git tag
	BinVersion string

	// GitStatus guarda info acer del estado del repo git
	GitStatus string

	// BuildTime almacena la fecha de compilación del binario
	BuildTime string

	// BuildGoVersion indica la versión de go usada para compilar este binario
	BuildGoVersion string
)

func init() {
	if GitStatus == "" {
		GitStatus = "up to date"
	} else {
		GitStatus = strings.Replace(strings.Replace(GitStatus, "\r\n", " | ", -1), "\n", " | ", -1)
	}
}

// main
// Obtiene el comando a ejecutar desde la línea de comandos y lo invoca
func main() {
	fmt.Printf("%s Version %s Build with Go %v at %v\n", provider.GetContainer().Config().APPName, BinVersion, BuildGoVersion, BuildTime)
	command := selectCmds()
	command()
}

// selectCmds Ejecuta los commands alojados en cmd/ según segundo parámetro de cli
func selectCmds() func() {

	if len(os.Args) == 1 {
		fmt.Println(`Ud. debería ejecutar esta aplicación con un comando`)

		return cmd.RunHelp
	}
	// os.Args = []string{"", "serve", "-config=./config_example.yml", "-port=4999", "-log=\"./b.b\"", "-development"}
	switch c := os.Args[1]; c {
	case "serve":
		return cmd.RunServe
	case "help":
		return cmd.RunHelp
	}

	return cmd.RunHelp
}
