package main

import (
	"strconv"
)

type Wire struct {
	Path []Coordinate
}

func (w Wire) lastX() int {
	return w.Path[len(w.Path)-1][0]

}
func (w Wire) lastY() int {
	return w.Path[len(w.Path)-1][1]
}

type Coordinate [2]int

func NewWire(c []string) Wire {
	//wire := Wire{Path: []Coordinate{Coordinate{0, 0}}}
	var wire Wire

	curX := 0
	curY := 0
	for _, coord := range c {
		direction := string(coord[0])
		length, _ := strconv.Atoi(coord[1:])
		switch direction {
		case "U":
			for i := curX + 1; i <= curX+length; i++ {
				wire.Path = append(wire.Path, Coordinate{i, curY})
			}
			curX += length
		case "D":
			for i := curX - 1; i >= curX-length; i-- {
				wire.Path = append(wire.Path, Coordinate{i, curY})
			}
			curX -= length
		case "R":
			for i := curY + 1; i <= curY+length; i++ {
				wire.Path = append(wire.Path, Coordinate{curX, i})
			}
			curY += length
		case "L":
			for i := curY - 1; i >= curY-length; i-- {
				wire.Path = append(wire.Path, Coordinate{curX, i})
			}
			curY -= length
		}
	}
	return wire
}

func FindCrosses(x, y Wire) []Coordinate {
	crosses := []Coordinate{}
	for _, xCoords := range x.Path {
		for _, yCoords := range y.Path {
			if xCoords == yCoords && xCoords != [2]int{0, 0} {
				crosses = append(crosses, yCoords)
			}
		}
	}
	return crosses
}

func FindClosestCross(x []Coordinate) int {
	shortest := abs(x[0][0]) + abs(x[0][1])
	for _, i := range x {
		if abs(i[0])+abs(i[1]) < shortest {
			shortest = abs(i[0]) + abs(i[1])
		}
	}
	return shortest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
