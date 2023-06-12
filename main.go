package main

import (
	"fmt"
	"gokota/pkg/cups"
	"log"
	"os"
	"syscall"
)

const BACKEND_NAME = "gokota"

func setupLogging() error {
	file, err := os.OpenFile(
		"/tmp/gokota.log",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return err
	}
	log.SetOutput(file)
	return nil
}

func main() {
	args := os.Args

	err := setupLogging()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}

	// Device discovery on stdout
	if len(args) == 1 {
		fmt.Printf("direct %s \"Unknown\" \"no device info\"\n", BACKEND_NAME)
		syscall.Exit(0)
	}

	if len(args) != 6 && len(args) != 7 {
		fmt.Fprintf(
			os.Stderr,
			"Error: %s job-id user title copies options [file]\n",
			args[0],
		)
		syscall.Exit(1)
	}

	backend := cups.Backend{}

	backend.InitParameters(args)
}
