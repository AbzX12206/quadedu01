package main

import "github.com/01-edu/z01"

func Quad1d(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	drawBorder := func(first, middle, last rune, length int) {
		z01.PrintRune(first)
		for i := 0; i < length-2; i++ {
			z01.PrintRune(middle)
		}
		if length > 1 {
			z01.PrintRune(last)
		}
		z01.PrintRune('\n')
	}

	drawMiddle := func(length int) {
		z01.PrintRune('B')
		for i := 0; i < length-2; i++ {
			z01.PrintRune(' ')
		}
		if length > 1 {
			z01.PrintRune('B')
		}
		z01.PrintRune('\n')
	}

	drawBorder('A', 'B', 'C', x)
	for i := 0; i < y-2; i++ {
		drawMiddle(x)
	}
	if y > 1 {
		drawBorder('A', 'B', 'C', x)
	}
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
		Quad1d(tc.x, tc.y)
		println()
	}
}