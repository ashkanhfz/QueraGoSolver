package main

import "fmt"

func GuessMyNumber(game Game) string {
	lo, hi := 1, 360
	for lo <= hi {
		mid := (lo + hi) / 2
		switch game.CheckNumber(mid) {
		case "CORRECT":
			return fmt.Sprintf("Your Number was %d", mid)
		case "My Number is Greater":
			lo = mid + 1
		case "My Number is Lower":
			hi = mid - 1
		}
	}
	return "Your Number was ?"
}
