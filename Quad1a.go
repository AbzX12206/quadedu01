package main

import "github.com/01-edu/z01"

func Quad1a(x, y int) {
    if x <= 0 || y <= 0 {
        return
    }

    printLine := func(left, middle, right rune) {
        z01.PrintRune(left)
        for i := 1; i < x-1; i++ {
            z01.PrintRune(middle)
        }
        if x > 1 {
            z01.PrintRune(right)
        }
        z01.PrintRune('\n')
    }

    printLine('o', '-', 'o')

    for i := 1; i < y-1; i++ {
        printLine('|', ' ', '|')
    }

    if y > 1 {
        printLine('o', '-', 'o')
    }
}

func main() {
     println("Program #1:")
    Quad1a(5, 3)
    println()

    println("Program #2:")
    Quad1a(5, 1)
    println()

    println("Program #3:")
    Quad1a(1, 1)
    println()

    println("Program #4:")
    Quad1a(1, 5)
    println()
}
