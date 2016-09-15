package main

import (
	"fmt"
	"github.com/tobyjsullivan/ships/world"
	"time"
)

func main() {
	la := &world.Port{
		Name: "Los Angeles"}
	van := &world.Port{
		Name: "Vancouver"}

	shipA := world.BuildShip("Ship A", la)
	shipB := world.BuildShip("Ship B", van)
	world.DockShip(shipA, van)
	world.DockShip(shipB, la)
	world.DockShip(shipA, la)

	time.Sleep(2 * time.Second)

	history := world.GetEventHistory()
	for i := 0; i < len(history); i++ {
		e := history[i]
		switch t := e.(type) {
		default:
			fmt.Println(fmt.Sprintf("Unexpected event %s", t))
		case *world.ShipDockedEvent:
			sde := e.(*world.ShipDockedEvent)
			fmt.Println(fmt.Sprintf("The ship %s docked at %s", sde.Ship.Name, sde.Port.Name))
		case *world.ShipBuiltEvent:
			sbe := e.(*world.ShipBuiltEvent)
			fmt.Println(fmt.Sprintf("The ship %s was built", sbe.Ship.Name))
		}
	}
}

