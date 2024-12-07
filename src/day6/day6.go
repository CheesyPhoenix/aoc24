package day6

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

func Run() {
	fmt.Println(getResult("./src/day6/input.txt"))
	fmt.Println(getResultP2("./src/day6/input.txt"))
}

type Vector2 struct {
	x int
	y int
}
type Vector2WithDirection struct {
	x     int
	y     int
	dir_x int
	dir_y int
}

func getResult(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to read file %s. Error: %s", filename, err))
	}

	text := string(data)
	lines := strings.Split(text, "\n")

	width := len([]rune(lines[0]))
	height := len(lines)

	x := -1
	y := -1

	for i, line := range lines {
		pos := strings.Index(line, "^")

		if pos != -1 {
			x = pos
			y = i
			break
		}
	}

	if x == -1 || y == -1 {
		panic("Unable to find guard")
	}

	dir_x := 0
	flip_x := false
	dir_y := -1
	flip_y := false

	// x -> y -> bool
	positions := map[Vector2]bool{}
	unique_positions := 0

	for {
		if !positions[Vector2{x: x, y: y}] {
			unique_positions++
		}
		positions[Vector2{x: x, y: y}] = true

		if y+dir_y < 0 || y+dir_y >= width || x+dir_x < 0 || x+dir_x >= height {
			break
		}

		next_pos := string([]rune(lines[y+dir_y])[x+dir_x])

		if next_pos != "#" {
			x += dir_x
			y += dir_y
		} else {
			if flip_x {
				dir_x--
			} else {
				dir_x++
			}

			if flip_y {
				dir_y--
			} else {
				dir_y++
			}

			if dir_x > 1 {
				dir_x = 0
				flip_x = true
			} else if dir_x < -1 {
				dir_x = 0
				flip_x = false
			}

			if dir_y > 1 {
				dir_y = 0
				flip_y = true
			} else if dir_y < -1 {
				dir_y = 0
				flip_y = false
			}
		}
	}

	return unique_positions
}

func isLoop(data string) bool {
	lines := strings.Split(data, "\n")

	width := len([]rune(lines[0]))
	height := len(lines)

	x := -1
	y := -1

	for i, line := range lines {
		pos := strings.Index(line, "^")

		if pos != -1 {
			x = pos
			y = i
			break
		}
	}

	if x == -1 || y == -1 {
		panic("Unable to find guard")
	}

	dir_x := 0
	flip_x := false
	dir_y := -1
	flip_y := false

	// x -> y -> bool
	positions := map[Vector2WithDirection]bool{}

	for {
		if positions[Vector2WithDirection{x: x, y: y, dir_x: dir_x, dir_y: dir_y}] {
			return true
		}
		positions[Vector2WithDirection{x: x, y: y, dir_x: dir_x, dir_y: dir_y}] = true

		if y+dir_y < 0 || y+dir_y >= height || x+dir_x < 0 || x+dir_x >= width {
			break
		}

		next_pos := string([]rune(lines[y+dir_y])[x+dir_x])

		if next_pos != "#" {
			x += dir_x
			y += dir_y
		} else {
			if flip_x {
				dir_x--
			} else {
				dir_x++
			}

			if flip_y {
				dir_y--
			} else {
				dir_y++
			}

			if dir_x > 1 {
				dir_x = 0
				flip_x = true
			} else if dir_x < -1 {
				dir_x = 0
				flip_x = false
			}

			if dir_y > 1 {
				dir_y = 0
				flip_y = true
			} else if dir_y < -1 {
				dir_y = 0
				flip_y = false
			}
		}
	}

	return false
}

func getResultP2(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to read file %s. Error: %s", filename, err))
	}

	characters := []rune(string(data))

	count := 0
	count_lock := sync.Mutex{}

	wg := sync.WaitGroup{}

	for i := 0; i < len(characters); i++ {
		if string(characters[i]) == "^" || string(characters[i]) == "\n" {
			continue
		}

		wg.Add(1)

		go func() {
			if isLoop(string(characters[0:i]) + "#" + string(characters[i+1:])) {
				count_lock.Lock()
				count++
				count_lock.Unlock()
			}

			wg.Done()
		}()
	}

	wg.Wait()

	return count
}
