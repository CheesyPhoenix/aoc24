package day4

import (
	"fmt"
	"os"
	"strings"
)

func Run() {
	fmt.Println(getResult("./src/day4/input.txt"))
	fmt.Println(getResultP2("./src/day4/input.txt"))
}

// This function is kinda scuffed as it does not preserve the proper alignment of characters
// between lines, and therefore cannot be chained multiple times.
func rotate45(text string) string {
	out := ""
	lines := strings.Split(text, "\n")

	start_y := 0
	start_x := 0
	for start_y < len(lines) && start_x < len(lines[start_y]) {
		x := start_x
		y := start_y

		for y >= 0 && x < len([]rune(lines[y])) {
			out += string([]rune(lines[y])[x])
			y--
			x++
		}

		out += "\n"
		if start_y < len(lines)-1 {
			start_y++
		} else {
			start_x++
		}
	}

	return out
}

func rotate90(text string) string {
	out := ""
	lines := strings.Split(text, "\n")

	width := len([]rune(lines[0]))
	height := len(lines)

	for x := 0; x < width; x++ {
		for y := height - 1; y >= 0; y-- {
			out += string(([]rune(lines[y]))[x])
		}
		out += "\n"
	}

	return out[:len(out)-3] // Remove trailing newline
}

func getResult(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to read file %s. Error: %s", filename, err))
	}

	text := string(data)

	sum := 0
	sum += strings.Count(text, "XMAS")
	sum += strings.Count(text, "SAMX")

	text45 := rotate45(text)
	sum += strings.Count(text45, "XMAS")
	sum += strings.Count(text45, "SAMX")

	text90 := rotate90(text)
	sum += strings.Count(text90, "XMAS")
	sum += strings.Count(text90, "SAMX")

	text135 := rotate45(text90)
	sum += strings.Count(text135, "XMAS")
	sum += strings.Count(text135, "SAMX")

	return sum
}

func getResultP2(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to read file %s. Error: %s", filename, err))
	}

	lines := strings.Split(string(data), "\n")

	width := len([]rune(lines[0]))
	height := len(lines)

	sum := 0

	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			if string([]rune(lines[x])[y]) != "A" {
				continue
			}

			tl := string([]rune(lines[x-1])[y-1])
			bl := string([]rune(lines[x-1])[y+1])
			tr := string([]rune(lines[x+1])[y-1])
			br := string([]rune(lines[x+1])[y+1])

			if (tl == "M" && br == "S") || (tl == "S" && br == "M") {
				if (bl == "M" && tr == "S") || (bl == "S" && tr == "M") {
					sum += 1
				}
			}
		}
	}

	return sum
}
