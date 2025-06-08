package main

import "github.com/01-edu/z01"

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			switch {
			case i == 0 && j == 0:
				z01.PrintRune('/')
			case i == 0 && j == x-1:
				z01.PrintRune('\\')
			case i == y-1 && j == 0:
				z01.PrintRune('\\')
			case i == y-1 && j == x-1:
				z01.PrintRune('/')
			case i == 0 || i == y-1 || j == 0 || j == x-1:
				z01.PrintRune('*')
			default:
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func main() {
	tests := []struct {
		x, y int
		desc string
	}{
		{5, 3, "Program #1:"},
		{5, 1, "Program #2:"},
		{1, 1, "Program #3:"},
		{1, 5, "Program #4:"},
	}

	for _, test := range tests {
		println(test.desc)
		QuadB(test.x, test.y)
		println()
	}
}