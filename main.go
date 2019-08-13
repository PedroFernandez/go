package main

import (
	"fmt"
	"os"
)

var beers = map[string]string{
	"AMAKLPF52B4U44H0UZ4L": "Wu Lei",
	"K43EXHJNGYDPN4LUD18N": "Víctor Sánchez",
	"7FYRUXSZIDCD3NY38FDR": "Diego López",
}

func main() {
	param := os.Args[1]
	if param == "beers" {
		fmt.Println(beers)
	}
	fmt.Println(param)
}
