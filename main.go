package main

import (
	"flag"
	"fmt"
)

var players = map[string]string{
	"AMAKLPF52B4U44H0UZ4L": "Wu Lei",
	"K43EXHJNGYDPN4LUD18N": "Víctor Sánchez",
	"7FYRUXSZIDCD3NY38FDR": "Diego López",
}

func main() {
	b := flag.Bool("players", false, "RCD Espanyol Players")
	flag.Parse()
	if *b {
		fmt.Println(players)
	}
}
