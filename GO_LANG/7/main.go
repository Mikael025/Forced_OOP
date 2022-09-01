package main

import "fmt"

func luasSegitiga (alas int, tinggi int) int {

	luas := alas * tinggi / 2

	return luas
}

func main() {
	var alas,tinggi int
	fmt.Print ("alas segitiga adalah : ")
	fmt.Scan(&alas)
	fmt.Println("tinggi segitiga adalah : ")
	fmt.Scan(&tinggi)

	var luas int
	
	luas = luasSegitiga(alas,tinggi)

	fmt.Println("luas segitiga adalah ",luas)
}
