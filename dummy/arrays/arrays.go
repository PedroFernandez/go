package main

import "fmt"

func main() {
	var surnames[4] string
	fmt.Println("empty", surnames)

	names := [5]string {"helena", "paula", "marta", "abu", "pedro"}
	fmt.Println("names", names)

	names[0] = "dolly helena"
	fmt.Println("names", names)
}
