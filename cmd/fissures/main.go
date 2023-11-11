package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func getInput() string {
	return `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`
}

type Point struct {
	x int
	y int
}

type Line struct {
	p1 *Point
	p2 *Point
}

func isStraight(line Line) bool {
	return line.p1.x == line.p2.x || line.p1.y == line.p2.y
}

func parsePoint(p string) (*Point, error) {
	parts := strings.Split(p, ",")

	xPos, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	yPos, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	return &Point{
		x: xPos,
		y: yPos,
	}, err
}

func parseLine(line string) (*Line, error) {
	points := strings.Split(line, " -> ")

	point1, err := parsePoint(points[0])
	if err != nil {
		return nil, err
	}

	point2, err := parsePoint(points[1])
	if err != nil {
		return nil, err
	}

	return &Line{
		p1: point1,
		p2: point2,
	}, err
}

func main() {
	for _, l := range strings.Split(getInput(), "\n") {
		line, err := parseLine(l)
		if err != nil {
			log.Fatal("This should not ever happen")
		}

		fmt.Printf("%+v : %+v \n", *line, isStraight(*line))
	}
}
