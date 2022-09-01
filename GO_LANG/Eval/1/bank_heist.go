package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true

	eludedGuards := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've manage to make it past the guards. Good job but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("you've caught, run!!!")
	}

	fmt.Println("isHeistOn status currently: ", isHeistOn)

	openedVault := rand.Intn(100)

	if openedVault >= 70 && isHeistOn == true {
		fmt.Println("Vault is open, grab it and go now!!")
	} else if isHeistOn == true {
		isHeistOn = false
		fmt.Println("Vault can't be opened, better luck next time")
	}

	leftSafely := rand.Intn(5)

	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Looks like you tripped     the laser trap...RUN!!??")
		case 1:
			isHeistOn = false
			fmt.Println("welp,apparently vault door is locked from outside ")
		case 2:
			isHeistOn = false
			fmt.Println("you forgot to bring the bag")
		case 3:
			isHeistOn = false
			fmt.Println("looks like the guards are rushing to your position ")
		default:
			fmt.Println("Get in to the car and Let's GO!!!")
		}

		if isHeistOn {
			amtStolen := 10000 + rand.Intn(100000)
			fmt.Println("We got around ", amtStolen, " worth of cash. Huh i guess its not bad at all")
		}
	}
}
