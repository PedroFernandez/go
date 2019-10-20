package main

import (
	."fmt"
)

func main() {
	Println("Go runs in: ")
	switch language := "goLang"; language {
	case "pepe":
		Println("This is not go")
	case "goLang":
		Println("This is GO")
	} 
}
