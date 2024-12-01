package day1

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println(getResult("./src/day1/input.txt"))
	fmt.Println(getResultP2("./src/day1/input.txt"))
}

func getResult(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to read file %s. Error: %s", filename, err))
	}

	lines := strings.Split(string(data), "\n")

	list1 := make([]int, len(lines))
	list2 := make([]int, len(lines))

	for i, line := range lines {
		values := strings.Split(line, " ")
		list1[i], err = strconv.Atoi(values[0])
		if err != nil {
			panic(fmt.Sprintf("Got error trying to parse first int on line %d. Value %s. Error: %s", i, values[0], err))
		}
		list2[i], err = strconv.Atoi(values[len(values)-1])
		if err != nil {
			panic(fmt.Sprintf("Got error trying to parse second int on line %d. Value %s. Error: %s", i, values[len(values)-1], err))
		}
	}

	sort.Ints(list1)
	sort.Ints(list2)

	sum := 0

	for i, first := range list1 {
		second := list2[i]

		if first > second {
			sum += first - second
		} else {
			sum += second - first
		}
	}

	return sum
}

func getResultP2(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to read file %s. Error: %s", filename, err))
	}

	lines := strings.Split(string(data), "\n")

	list1 := make([]int, len(lines))
	list2 := make([]int, len(lines))

	for i, line := range lines {
		values := strings.Split(line, " ")
		list1[i], err = strconv.Atoi(values[0])
		if err != nil {
			panic(fmt.Sprintf("Got error trying to parse first int on line %d. Value %s. Error: %s", i, values[0], err))
		}
		list2[i], err = strconv.Atoi(values[len(values)-1])
		if err != nil {
			panic(fmt.Sprintf("Got error trying to parse second int on line %d. Value %s. Error: %s", i, values[len(values)-1], err))
		}
	}

	sort.Ints(list1)

	appearance_count := map[int]int{}

	for _, num := range list2 {
		appearance_count[num] += 1
	}

	sum := 0

	for _, first := range list1 {
		sum += first * appearance_count[first]
	}

	return sum
}
