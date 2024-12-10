package day9

import (
	"fmt"
	"os"
	"strconv"
)

func Run() {
	fmt.Println(getResult("./src/day9/input.txt"))
	fmt.Println(getResultP2("./src/day9/input.txt"))
}

func getResult(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to read file %s. Error: %s", filename, err))
	}

	text := string(data)

	disk := []*int{}

	for i, rune := range []rune(text) {
		number, err := strconv.Atoi(string(rune))
		if err != nil {
			panic(fmt.Sprintf("Failed to parse number %s. Error: %s", string(rune), err))
		}
		ids := make([]*int, number)

		if i%2 == 0 {
			// file
			id := i / 2
			for n := 0; n < number; n++ {
				ids[n] = &id
			}
		}
		disk = append(disk, ids...)
	}

	last_index := len(disk) - 1

outer:
	for i, place := range disk {
		if place != nil {
			continue
		}

		for {
			if last_index <= i {
				break outer
			}

			if disk[last_index] == nil {
				last_index--
				continue
			}

			disk[i] = disk[last_index]
			disk[last_index] = nil
			last_index--
			if last_index <= i {
				break outer
			}
			break
		}
	}

	sum := 0

	for i, el := range disk {
		if el == nil {
			break
		}

		sum += i * *el
	}

	return sum
}

func getResultP2(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to read file %s. Error: %s", filename, err))
	}

	text := string(data)

	disk := []*int{}

	for i, char := range []rune(text) {
		number, err := strconv.Atoi(string(char))
		if err != nil {
			panic(fmt.Sprintf("Failed to parse number %s. Error: %s", string(char), err))
		}
		ids := make([]*int, number)

		if i%2 == 0 {
			// file
			id := i / 2
			for n := 0; n < number; n++ {
				ids[n] = &id
			}
		}
		disk = append(disk, ids...)
	}

	last_index := len(disk) - 1

	attempted := map[int]bool{}

	for last_index >= 0 {
		id := disk[last_index]
		if id == nil || attempted[*id] {
			last_index--
			continue
		}
		attempted[*id] = true

		len := 1
		for {
			if last_index-len < 0 {
				break
			}
			if disk[last_index-len] == nil {
				break
			}
			if *disk[last_index-len] != *id {
				break
			}
			len++
		}

		for i, el := range disk[:last_index-len+1] {
			if el != nil {
				continue
			}

			block_len := 1
			for block_len < len && disk[i+block_len] == nil {
				block_len++
			}

			if block_len >= len {
				for x := 0; x < len; x++ {
					disk[i+x] = id
					disk[last_index-x] = nil
				}
				break
			}
		}

		last_index -= len
	}

	sum := 0

	for i, el := range disk {
		if el == nil {
			continue
		}

		sum += i * *el
	}

	return sum
}
