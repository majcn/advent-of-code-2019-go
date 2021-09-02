package util

import "strings"

type Location struct {
	X int
	Y int
}

func (p Location) Add(q Location) Location {
	return Location{p.X + q.X, p.Y + q.Y}
}

func (p Location) Mul(k int) Location {
	return Location{p.X * k, p.Y * k}
}

type Grid map[Location]void

func (g *Grid) Add(location Location) {
	(*g)[location] = voidVar
}

func (g *Grid) Remove(location Location) {
	delete(*g, location)
}

func (g *Grid) Contains(location Location) bool {
	if _, ok := (*g)[location]; ok {
		return true
	}

	return false
}

func NewGridFromString(data string, el rune) Grid {
	grid := make(Grid)
	for y, line := range strings.Split(data, "\n") {
		for x, c := range line {
			if c == el {
				grid.Add(Location{X: x, Y: y})
			}
		}
	}
	return grid
}

var neighbours8 = []Location{
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
}

var neighbours4 = []Location{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func GetNeighbours4() []Location {
	return neighbours4
}

func GetNeighbours8() []Location {
	return neighbours8
}
