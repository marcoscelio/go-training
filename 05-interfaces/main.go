package main

import "fmt"

type Instrument interface {
	Sound()
}

type Guitar struct {
	Player string
}

type Drum struct {
	Player string
}

func (b Guitar) Sound() {
	fmt.Printf("%s does a guitar sound\n", b.Player)
}

func (b Drum) Sound() {
	fmt.Printf("%s does a drum sound\n", b.Player)
}

func main() {
	var instrument1 Guitar
	instrument1.Player = "Paul"
	var instrument2 Drum
	instrument2.Player = "Ringo"

	players := map[string]Instrument{"Paul": instrument1, "Ringo": instrument2}

	for k, v := range players {
		fmt.Println(k)
		v.Sound()
	}

}
