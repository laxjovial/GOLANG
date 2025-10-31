package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printMsg() {
	msg := []rune{120, 32, 61, 32, 52, 50, 44, 32, 121, 32, 61, 32, 50, 49, 10}
	for i := 0; i < 15; i++ {
		z01.PrintRune(msg[i])
	}
}

func main() {
	points := &point{}
	setPoint(points)
	printMsg()
}
