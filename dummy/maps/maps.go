package main

import "fmt"

func main() {
	family := make(map[string]string)
	family["father"] = "pedro"
	family["mother"] = "helena"
	family["children"] = "marta y paula"
	fmt.Println(family)
}
