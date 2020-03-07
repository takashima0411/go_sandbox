package main

import (
	"./money"
	"fmt"
)

func main() {
	mon := []money.Money{
		money.NewSen(),
		money.NewGosen(),
		money.NewSen(),
		money.NewSen(),
		money.NewSen(),
		money.NewSen(),
	}

	reversed := money.Reverse(mon)

	for _, val := range reversed {
		fmt.Println(val)
	}
}
