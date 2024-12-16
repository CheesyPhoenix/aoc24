package day11

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println(getResult("./src/day11/input.txt"))
	fmt.Println(getResultP2("./src/day11/input.txt"))
}

func getResult(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to read file %s. Error: %s", filename, err))
	}

	text := string(data)

	stones_str := strings.Split(text, " ")
	stones := make([]int, len(stones_str))

	for i, stone := range stones_str {
		num, err := strconv.Atoi(stone)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse num %s. Error: %s", stone, err))
		}
		stones[i] = num
	}

	for i := 0; i < 25; i++ {
		for x := 0; x < len(stones); x++ {
			stone := stones[x]

			if stone == 0 {
				stones[x] = 1
				continue
			}

			str := []rune(strconv.Itoa(stone))

			if len(str)%2 == 0 {
				left, err := strconv.Atoi(string(str[:len(str)/2]))
				if err != nil {
					panic(fmt.Sprintf("Failed to parse num %s. Error: %s", string(str[:len(str)/2]), err))
				}

				right, err := strconv.Atoi(string(str[len(str)/2:]))
				if err != nil {
					panic(fmt.Sprintf("Failed to parse num %s. Error: %s", string(str[len(str)/2:]), err))
				}

				stones[x] = left
				stones = slices.Insert(stones, x+1, right)
				x++
			} else {
				stones[x] = stone * 2024
			}
		}
	}

	return len(stones)
}

func getResultP2(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to read file %s. Error: %s", filename, err))
	}

	text := string(data)

	stones_str := strings.Split(text, " ")
	stones := map[int]int{}

	for _, stone := range stones_str {
		num, err := strconv.Atoi(stone)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse num %s. Error: %s", stone, err))
		}
		stones[num]++
	}

	for i := 0; i < 75; i++ {
		new_stones := map[int]int{}

		for stone, count := range stones {
			if count == 0 {
				continue
			}

			if stone == 0 {
				new_stones[1] += count

				continue
			}

			str := []rune(strconv.Itoa(stone))

			if len(str)%2 == 0 {
				left, err := strconv.Atoi(string(str[:len(str)/2]))
				if err != nil {
					panic(fmt.Sprintf("Failed to parse num %s. Error: %s", string(str[:len(str)/2]), err))
				}

				right, err := strconv.Atoi(string(str[len(str)/2:]))
				if err != nil {
					panic(fmt.Sprintf("Failed to parse num %s. Error: %s", string(str[len(str)/2:]), err))
				}

				new_stones[left] += count
				new_stones[right] += count
			} else {
				new_stones[stone*2024] += count
			}
		}

		stones = new_stones
	}

	sum := 0

	for _, count := range stones {
		sum += count
	}

	return sum
}
