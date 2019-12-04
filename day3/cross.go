package main

import (
	"math"
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

func FindCrosses(x, y Wire) map[Coordinate]int {
	crosses := make(map[Coordinate]int)
	for xIndex, xCoords := range x.Path {
		for yIndex, yCoords := range y.Path {
			if xCoords == yCoords && xCoords != [2]int{0, 0} {
				// Have to add 2 because index starts at 0
				crosses[yCoords] = xIndex + yIndex + 2
			}
		}
	}
	return crosses
}

func FindClosestCross(x map[Coordinate]int) int {
	//shortest := abs(x[0][0]) + abs(x[0][1])
	shortest := math.MaxInt64
	for k, _ := range x {
		if abs(k[0])+abs(k[1]) < shortest {
			shortest = abs(k[0]) + abs(k[1])
		}
	}
	return shortest
}

func FindLeastStepsCross(x map[Coordinate]int) int {
	shortest := math.MaxInt64
	for _, v := range x {
		if v < shortest {
			shortest = v
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
