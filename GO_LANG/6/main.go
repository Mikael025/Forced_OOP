package main

import "fmt"
import "os"

func main() {
	var input string
	var heistReady bool
	var exit string
	tipe:= "type anything to quit"

	fmt.Println("is the heist ready?? (input 1 for ready, 2 if not ready)")

	fmt.Scan(&input)

	if (input == "1") {
		heistReady = true
	} else if (input == "2") {
		heistReady = false
	} else {
		fmt.Println("Invalid input")
		fmt.Println(tipe)
		fmt.Scan(&exit)
		os.Exit(0)
	}


	if heistReady {
		fmt.Println("we Go NOW!!!")
	} else if !heistReady {
		fmt.Println("we need to prepare")
	}

	fmt.Println(tipe)
	fmt.Scan(&exit)
}