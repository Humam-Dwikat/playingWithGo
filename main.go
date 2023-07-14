package main

import (
	"firstProject/trying_Interface"
	"fmt"
)

func main() {
	// circle area
	circle := Interface.Circle{Radius: 1.2}
	circle.Draw()
	fmt.Print("\n")
	//square area
	square := Interface.Square{Length: 9}
	square.Draw()
	fmt.Print("\n")
	// triangle area
	triangle := Interface.Triangle{Base: 4, Height: 3}
	triangle.Draw()
}
