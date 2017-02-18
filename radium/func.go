package radium

import "math"

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Given a list of Coordinates C, and a Coordinate X, find all c that are n degrees away from X

// In general, x and y must satisfy (x - center_x)^2 + (y - center_y)^2 < radius^2

func (c *Coordinate) circleContains(radius float64, loc *Coordinate) bool {
	return math.Pow(loc.Latitude-c.Latitude, 2)+math.Pow(loc.Longitude-c.Longitude, 2) <= math.Pow(radius, 2)
}

func (c *Coordinate) GetCoordinatesWithinRadius(radius float64, arr []*Coordinate) (result []*Coordinate) {
	for _, point := range arr {
		if c.circleContains(radius, point) {
			result = append(result, point)
		}
	}
	return result
}
