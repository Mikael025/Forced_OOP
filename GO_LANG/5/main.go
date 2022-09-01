package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	gacha := rand.Intn(100)
	var hasil string

	if (gacha <= 50) {
		hasil = " gachamu ampas"
	} else if (gacha>50 && gacha<80) {
		hasil = " lumayan gachamu"
	} else { 
		hasil = " Wuih bejo"
	}

	fmt.Println("Hasil gachamu ",gacha, hasil)
}