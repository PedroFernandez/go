package main

import "fmt"

func main() {
	var nbaPlayers = make([]string, 3)
	nbaPlayers[0] = "M Jordan"
	nbaPlayers[1] = "Kobe"
	nbaPlayers[2] = "Lebron James"

	fmt.Println("Original NBA Players: ", nbaPlayers)

	nbaPlayers = append(nbaPlayers, "Tim Duncan", "Bill Russell")

	fmt.Println("New NBA Players have been added: ", nbaPlayers)

	anotherPlayers := nbaPlayers[:5]

	fmt.Println("Another Players have been added: ", anotherPlayers)

	anotherPlayers[4] = "Kevin Love"

	fmt.Println("Tim Duncan has been replaced by injury: ", anotherPlayers)
}
