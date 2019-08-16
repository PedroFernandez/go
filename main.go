package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var beers = map[string]string{
	"AMAKLPF52B4U44H0UZ4L": "Voll-Damm",
	"K43EXHJNGYDPN4LUD18N": "Estrella Galicia",
	"7FYRUXSZIDCD3NY38FDR": "Heineken",
}

func main() {
	beersCmd := flag.NewFlagSet("beers", flag.ExitOnError)
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("Command beers must be included in the statement")
		os.Exit(2)
	}

	switch flag.Arg(0) {
	case "beers":
		ID := beersCmd.String("id", "", "find by Id")

		beersCmd.Parse(os.Args[2:])

		if *ID != "" {
			fmt.Println(beers[*ID])
		} else {
			fmt.Println(beers)
		}
	}
}
