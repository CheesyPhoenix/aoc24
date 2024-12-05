package day5

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println(getResult("./src/day5/input.txt"))
	fmt.Println(getResultP2("./src/day5/input.txt"))
}

func getResult(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to read file %s. Error: %s", filename, err))
	}

	text := string(data)

	before_list := map[int][]int{}

	for _, line := range strings.Split(text, "\n") {
		if line == "" {
			break
		}

		items := strings.Split(line, "|")

		first, err := strconv.Atoi(items[0])
		if err != nil {
			panic(fmt.Sprintf("Got error trying to parse %s as number. Line %s. Error: %s", items[0], line, err))
		}
		last, err := strconv.Atoi(items[1])
		if err != nil {
			panic(fmt.Sprintf("Got error trying to parse %s as number. Line %s. Error: %s", items[0], line, err))
		}

		before_list[first] = append(before_list[first], last)
	}

	sum := 0

outer:
	for _, line := range strings.Split(strings.Split(text, "\n\n")[1], "\n") {
		pages_str := strings.Split(line, ",")
		pages_out := []int{}

		for _, page_str := range pages_str {
			page, err := strconv.Atoi(page_str)
			if err != nil {
				panic(fmt.Sprintf("Got error trying to parse %s as number. Error: %s", page_str, err))
			}

			before_pages := before_list[page]
			// after_pages := after_list[page]

			loc := len(pages_out)

			for _, before_page := range before_pages {
				index := slices.Index(pages_out, before_page)
				if index != -1 && index < loc {
					// loc = index

					continue outer
				}
			}

			pages_out = slices.Insert(pages_out, loc, page)
		}

		sum += pages_out[len(pages_out)/2]
	}

	return sum
}

func getResultP2(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to read file %s. Error: %s", filename, err))
	}

	text := string(data)

	before_list := map[int][]int{}

	for _, line := range strings.Split(text, "\n") {
		if line == "" {
			break
		}

		items := strings.Split(line, "|")

		first, err := strconv.Atoi(items[0])
		if err != nil {
			panic(fmt.Sprintf("Got error trying to parse %s as number. Line %s. Error: %s", items[0], line, err))
		}
		last, err := strconv.Atoi(items[1])
		if err != nil {
			panic(fmt.Sprintf("Got error trying to parse %s as number. Line %s. Error: %s", items[0], line, err))
		}

		before_list[first] = append(before_list[first], last)
	}

	sum := 0

	for _, line := range strings.Split(strings.Split(text, "\n\n")[1], "\n") {
		pages_str := strings.Split(line, ",")
		pages_out := []int{}

		valid := true

		for _, page_str := range pages_str {
			page, err := strconv.Atoi(page_str)
			if err != nil {
				panic(fmt.Sprintf("Got error trying to parse %s as number. Error: %s", page_str, err))
			}

			before_pages := before_list[page]
			// after_pages := after_list[page]

			loc := len(pages_out)

			for _, before_page := range before_pages {
				index := slices.Index(pages_out, before_page)
				if index != -1 && index < loc {
					loc = index
					valid = false
				}
			}

			pages_out = slices.Insert(pages_out, loc, page)
		}

		if !valid {
			sum += pages_out[len(pages_out)/2]
		}
	}

	return sum
}
