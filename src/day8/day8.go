package day8

import (
	"fmt"
	"os"
	"strings"
)

func Run() {
	fmt.Println(getResult("./src/day8/input.txt"))
	fmt.Println(getResultP2("./src/day8/input.txt"))
}

type Vector2 struct {
	x int
	y int
}

func getResult(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to read file %s. Error: %s", filename, err))
	}

	text := string(data)
	lines := strings.Split(text, "\n")

	nodes := map[string][]Vector2{}

	for y, line := range lines {
		for x, char_rune := range []rune(line) {
			char := string(char_rune)

			if char == "." {
				continue
			}

			if nodes[char] != nil {
				nodes[char] = append(nodes[char], Vector2{x: x, y: y})
			} else {
				nodes[char] = []Vector2{{x: x, y: y}}
			}
		}
	}

	unique_locations := map[Vector2]bool{}
	count := 0

	width := len([]rune(lines[0]))
	height := len(lines)

	for _, locations := range nodes {
		for i1, loc := range locations {
			for i2, other_loc := range locations {
				if i1 == i2 {
					continue
				}

				diff_x := other_loc.x - loc.x
				diff_y := other_loc.y - loc.y

				x := other_loc.x + diff_x
				y := other_loc.y + diff_y

				if x < width && x >= 0 && y < height && y >= 0 {
					if !unique_locations[Vector2{x: x, y: y}] {
						count++
						unique_locations[Vector2{x: x, y: y}] = true
					}
				}
			}
		}
	}

	return count
}

func getResultP2(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to read file %s. Error: %s", filename, err))
	}

	text := string(data)
	lines := strings.Split(text, "\n")

	nodes := map[string][]Vector2{}

	for y, line := range lines {
		for x, char_rune := range []rune(line) {
			char := string(char_rune)

			if char == "." {
				continue
			}

			if nodes[char] != nil {
				nodes[char] = append(nodes[char], Vector2{x: x, y: y})
			} else {
				nodes[char] = []Vector2{{x: x, y: y}}
			}
		}
	}

	unique_locations := map[Vector2]bool{}
	count := 0

	width := len([]rune(lines[0]))
	height := len(lines)

	for _, locations := range nodes {
		for i1, loc := range locations {
			for i2, other_loc := range locations {
				if i1 == i2 {
					continue
				}

				diff_x := other_loc.x - loc.x
				diff_y := other_loc.y - loc.y

				x := other_loc.x
				y := other_loc.y
				for x < width && x >= 0 && y < height && y >= 0 {
					if !unique_locations[Vector2{x: x, y: y}] {
						count++
						unique_locations[Vector2{x: x, y: y}] = true
					}

					x += diff_x
					y += diff_y
				}
			}
		}
	}

	return count
}
