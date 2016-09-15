package world

type Port struct {
	Name string
}

type ShipDockedEvent struct {
	Ship *Ship
	Port *Port
}

func DockShip(ship *Ship, port *Port) {
	ship.Location = port

	fireShipDocked(ship, port)
}

func (*ShipDockedEvent) GetEventType() string {
	return "ShipDocked"
}

func fireShipDocked(ship *Ship, port *Port) {
	logEvent(&ShipDockedEvent{
		Ship: ship,
		Port: port})
}