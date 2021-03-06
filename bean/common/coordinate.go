package common

import "fmt"

//Coordinate ***
type Coordinate struct {
	X float64
	Y float64
	Z float64
}

//NewCoordinate ***
func NewCoordinate(x, y, z float64) *Coordinate {
	o := &Coordinate{
		X: x,
		Y: y,
		Z: z,
	}
	return o
}

// Equals ***
func (o *Coordinate) Equals(d *Coordinate) bool {
	if d == nil {
		return false
	}

	return (o.X == d.X) && (o.Y == d.Y) && (o.Z == d.Z)
}

// ToString get the string
func (o *Coordinate) ToString() string {
	return fmt.Sprintf("Coordinate [X=%v, Y=%v, Z=%v]", o.X, o.Y, o.Z)
}
