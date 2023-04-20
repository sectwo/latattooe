package main

import (
	"flag"
	"fmt"
	"os"
)

func CliStart() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 8000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		//start rest api
		RestStart(*port)
	default:
		usage()
	}

	fmt.Println(*port, *mode)
}
func usage() {
	fmt.Printf("Welcome to LATATTOE\n")
	fmt.Printf("Please use the following flags : \n")
	fmt.Printf("-port:		Set the PORT of the service\n")
	fmt.Printf("-mode:		Choose 'rest'\n\n")
	os.Exit(0)
}
