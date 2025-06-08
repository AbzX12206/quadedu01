package main

import "github.com/01-edu/z01"

type BorderChars struct {
	first  rune
	middle rune
	last   rune
}

var (
	topBorder    = BorderChars{'A', 'B', 'A'}
	middleBorder = BorderChars{'B', ' ', 'B'}
	bottomBorder = BorderChars{'C', 'B', 'C'}
)

func printBorder(b BorderChars, width int) {
	if width <= 0 {
		return
	}
	z01.PrintRune(b.first)
	for i := 0; i < width-2; i++ {
		z01.PrintRune(b.middle)
	}
	if width > 1 {
		z01.PrintRune(b.last)
	}
	z01.PrintRune('\n')
}

func printVertical(width, height int) {
	for i := 0; i < height-2; i++ {
		printBorder(middleBorder, width)
	}
}

func QuadC(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	if y == 1 {
		printBorder(topBorder, x)
		return
	}

	if x == 1 {
		z01.PrintRune('A')
		z01.PrintRune('\n')
		for i := 0; i < y-2; i++ {
			z01.PrintRune('B')
			z01.PrintRune('\n')
		}
		z01.PrintRune('C')
		z01.PrintRune('\n')
		return
	}

	printBorder(topBorder, x)
	printVertical(x, y)
	printBorder(bottomBorder, x)
}

func main() {
	testCases := []struct {
		x, y  int
		title string
	}{
		{5, 3, "Program #1:"},
		{5, 1, "Program #2:"},
		{1, 1, "Program #3:"},
		{1, 5, "Program #4:"},
	}

	for _, tc := range testCases {
		println(tc.title)
		QuadC(tc.x, tc.y)
		println()
	}
}