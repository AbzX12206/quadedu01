package main

import "github.com/01-edu/z01"

type corners struct {
	topLeft     rune
	topRight    rune
	bottomLeft  rune
	bottomRight rune
}

var quadECorners = corners{
	topLeft:     'A',
	topRight:    'C',
	bottomLeft:  'C',
	bottomRight: 'A',
}

func getChar(col, row, x, y int, c corners) rune {
	switch {
	case row == 1 && col == 1:
		return c.topLeft
	case row == 1 && col == x:
		return c.topRight
	case row == y && col == 1:
		return c.bottomLeft
	case row == y && col == x:
		return c.bottomRight
	case row == 1 || row == y || col == 1 || col == x:
		return 'B'
	default:
		return ' '
	}
}

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			z01.PrintRune(getChar(col, row, x, y, quadECorners))
		}
		z01.PrintRune('\n')
	}
}

type testCase struct {
	x, y  int
	title string
}
func runTests() {
	tests := []testCase{
		{5, 3, "Program #1:"},
		{5, 1, "Program #2:"},
		{1, 1, "Program #3:"},
		{1, 5, "Program #4:"},
	}

	for _, test := range tests {
		println(test.title)
		QuadE(test.x, test.y)
		println()
	}
}

func main() {
	runTests()
}