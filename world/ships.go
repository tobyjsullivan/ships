package world

type Ship struct {
	Name string
	Location Location
}

type ShipBuiltEvent struct {
	Ship *Ship
}

func BuildShip(name string, port *Port) *Ship {
	ship := &Ship{
		Name: name}
	fireShipBuilt(ship)
	DockShip(ship, port)

	return ship
}

func (*ShipBuiltEvent) GetEventType() string {
	return "ShipBuilt"
}

func fireShipBuilt(ship *Ship) {
	logEvent(&ShipBuiltEvent{
		Ship: ship})
}
