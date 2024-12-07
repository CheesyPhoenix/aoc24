package day7

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println(getResult("./src/day7/input.txt"))
	fmt.Println(getResultP2("./src/day7/input.txt"))
}

func getValues(str_numbers []string) []int {
	first_number, err := strconv.Atoi(str_numbers[0])
	if err != nil {
		panic(fmt.Sprintf("Unable to parse number %s. Error: %s", str_numbers[0], err))
	}

	values := []int{first_number}

	for _, str_number := range str_numbers[1:] {
		new_values := []int{}
		number, err := strconv.Atoi(str_number)
		if err != nil {
			panic(fmt.Sprintf("Unable to parse number %s. Error: %s", str_number, err))
		}

		for _, option := range values {
			new_values = append(new_values, option+number, option*number)
		}
		values = new_values
	}

	return values
}

func getResult(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to read file %s. Error: %s", filename, err))
	}

	text := string(data)
	lines := strings.Split(text, "\n")

	sum := 0

	for _, line := range lines {
		split := strings.Split(line, ": ")
		target, err := strconv.Atoi(split[0])
		if err != nil {
			panic(fmt.Sprintf("Unable to parse number %s. Error: %s", split[0], err))
		}

		numbers := strings.Split(split[1], " ")
		if slices.Index(getValues(numbers), target) != -1 {
			sum += target
		}
	}

	return sum
}

func getValuesP2(str_numbers []string) []int {
	first_number, err := strconv.Atoi(str_numbers[0])
	if err != nil {
		panic(fmt.Sprintf("Unable to parse number %s. Error: %s", str_numbers[0], err))
	}

	values := []int{first_number}

	for _, str_number := range str_numbers[1:] {
		new_values := []int{}
		number, err := strconv.Atoi(str_number)
		if err != nil {
			panic(fmt.Sprintf("Unable to parse number %s. Error: %s", str_number, err))
		}

		for _, option := range values {
			concat, err := strconv.Atoi(strconv.Itoa(option) + str_number)
			if err != nil {
				panic(fmt.Sprintf("Unable to parse number %s. Error: %s", strconv.Itoa(option)+str_number, err))
			}
			new_values = append(new_values, option+number, option*number, concat)
		}
		values = new_values
	}

	return values
}

func getResultP2(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to read file %s. Error: %s", filename, err))
	}

	text := string(data)
	lines := strings.Split(text, "\n")

	sum := 0

	for _, line := range lines {
		split := strings.Split(line, ": ")
		target, err := strconv.Atoi(split[0])
		if err != nil {
			panic(fmt.Sprintf("Unable to parse number %s. Error: %s", split[0], err))
		}

		numbers := strings.Split(split[1], " ")
		if slices.Index(getValuesP2(numbers), target) != -1 {
			sum += target
		}
	}

	return sum
}
