package day10

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println(getResult("./src/day10/input.txt"))
	fmt.Println(getResultP2("./src/day10/input.txt"))
}

type Vector2 struct {
	x int
	y int
}

func parsePos(x int, y int, lines []string) int {
	pos, err := strconv.Atoi(string([]rune(lines[y])[x]))
	if err != nil {
		panic(fmt.Sprintf("Unable to parse %s as number. Error: %s", string([]rune(lines[y])[x]), err))
	}

	return pos
}

func addTops(tops_map map[Vector2]bool, new_tops map[Vector2]bool) {
	for k, v := range new_tops {
		if v {
			tops_map[k] = true
		}
	}
}

func getReachableTops(x int, y int, lines []string) map[Vector2]bool {
	tops := map[Vector2]bool{}

	currLevel := parsePos(x, y, lines)

	if currLevel == 9 {
		tops[Vector2{x: x, y: y}] = true
		return tops
	}

	// fmt.Printf("X: %d, Y: %d, CURR: %d.", x, y, currLevel)

	// if y-1 >= 0 {
	// 	fmt.Printf(" UP: %d", parsePos(x, y-1, lines))
	// }
	// if x-1 >= 0 {
	// 	fmt.Printf(" Left: %d", parsePos(x-1, y, lines))
	// }
	// if y+1 < len(lines) {
	// 	fmt.Printf(" Down: %d", parsePos(x, y+1, lines))
	// }
	// if x+1 < len([]rune(lines[y])) {
	// 	fmt.Printf(" Right: %d", parsePos(x+1, y, lines))
	// }
	// fmt.Println()

	if y-1 >= 0 && parsePos(x, y-1, lines) == currLevel+1 {
		addTops(tops, getReachableTops(x, y-1, lines))
	}
	if x-1 >= 0 && parsePos(x-1, y, lines) == currLevel+1 {
		addTops(tops, getReachableTops(x-1, y, lines))
	}

	if y+1 < len(lines) && parsePos(x, y+1, lines) == currLevel+1 {
		addTops(tops, getReachableTops(x, y+1, lines))
	}
	if x+1 < len([]rune(lines[y])) && parsePos(x+1, y, lines) == currLevel+1 {
		addTops(tops, getReachableTops(x+1, y, lines))
	}

	return tops
}

func getResult(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to read file %s. Error: %s", filename, err))
	}

	text := string(data)
	lines := strings.Split(text, "\n")

	starts := []Vector2{}

	for y, line := range lines {
		for x, char := range []rune(line) {
			if string(char) == "0" {
				starts = append(starts, Vector2{x: x, y: y})
			}
		}
	}

	sum := 0

	for _, start := range starts {
		tops := getReachableTops(start.x, start.y, lines)

		for _, v := range tops {
			if v {
				sum++
			}
		}
	}

	return sum
}

func getTrailCount(x int, y int, lines []string) int {
	currLevel := parsePos(x, y, lines)

	if currLevel == 9 {
		return 1
	}

	sum := 0

	if y-1 >= 0 && parsePos(x, y-1, lines) == currLevel+1 {
		sum += getTrailCount(x, y-1, lines)
	}
	if x-1 >= 0 && parsePos(x-1, y, lines) == currLevel+1 {
		sum += getTrailCount(x-1, y, lines)
	}

	if y+1 < len(lines) && parsePos(x, y+1, lines) == currLevel+1 {
		sum += getTrailCount(x, y+1, lines)
	}
	if x+1 < len([]rune(lines[y])) && parsePos(x+1, y, lines) == currLevel+1 {
		sum += getTrailCount(x+1, y, lines)
	}

	return sum
}

func getResultP2(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to read file %s. Error: %s", filename, err))
	}

	text := string(data)
	lines := strings.Split(text, "\n")

	starts := []Vector2{}

	for y, line := range lines {
		for x, char := range []rune(line) {
			if string(char) == "0" {
				starts = append(starts, Vector2{x: x, y: y})
			}
		}
	}

	sum := 0

	for _, start := range starts {
		sum += getTrailCount(start.x, start.y, lines)
	}

	return sum
}
