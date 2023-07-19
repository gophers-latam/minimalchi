package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gophers-latam/minimal-chi/api"
	"github.com/gophers-latam/minimal-chi/internal/provider"
)

// RunServe pone a escuchar a la api en el puerto indicado
func RunServe() {

	// Generamos un canal que reciba señales desde el OS para detectar una orden de detención como [CTRL-C]
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go api.Start()

	<-c

	// Put here anything you want to do at the shutdown of the process
	// ensure closing of log file
	fmt.Printf("Shutting down %s. Bye!\n", provider.GetContainer().Config().APPName)

	os.Exit(0)
}
