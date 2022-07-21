package main

import (
	"fmt"

	dotagsi "github.com/dlujetic/kladara/data"
)

func main()  {
	game := dotagsi.New(0)
	
	go func() {
		for state := range game.Channel {
			for kills := range state.Player.Name {
				kills := state.Player.Name
				if(kills.State == "active") {
					fmt.Printf("%s \n", hero.Name)
				}
			}
		}
	}()

	game.Listen(":3000")
}