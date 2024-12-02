package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println(getResult("./src/day2/input.txt"))
	fmt.Println(getResultP2("./src/day2/input.txt"))
}

func getResult(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to read file %s. Error: %s", filename, err))
	}

	lines := strings.Split(string(data), "\n")

	num_safe := 0

outer:
	for _, line := range lines {
		values := strings.Split(line, " ")

		prev_value, err := strconv.Atoi(values[0])
		if err != nil {
			panic(fmt.Sprintf("Got error parsing number %s", values[0]))
		}

		second_value, err := strconv.Atoi(values[1])
		if err != nil {
			panic(fmt.Sprintf("Got error parsing number %s", values[0]))
		}

		is_increasing := false
		if second_value > prev_value {
			is_increasing = true
		} else if second_value == prev_value {
			continue // unsafe
		}

		for i := 1; i < len(values); i++ {
			value, err := strconv.Atoi(values[i])
			if err != nil {
				panic(fmt.Sprintf("Got error parsing number %s", values[i]))
			}

			if is_increasing {
				if value <= prev_value || value-3 > prev_value {
					continue outer
				}
			} else {
				if value >= prev_value || value+3 < prev_value {
					continue outer
				}
			}

			prev_value = value
		}

		num_safe++
	}

	return num_safe
}

func isSafe(values []string, problem_dampened bool) bool {
	prev_value, err := strconv.Atoi(values[0])
	if err != nil {
		panic(fmt.Sprintf("Got error parsing number %s", values[0]))
	}

	second_value, err := strconv.Atoi(values[1])
	if err != nil {
		panic(fmt.Sprintf("Got error parsing number %s", values[0]))
	}

	is_increasing := false

	if second_value > prev_value {
		is_increasing = true
	} else if second_value == prev_value {
		return false
	}

	for i := 1; i < len(values); i++ {
		value, err := strconv.Atoi(values[i])
		if err != nil {
			panic(fmt.Sprintf("Got error parsing number %s", values[i]))
		}

		if is_increasing {
			if value <= prev_value || value-3 > prev_value {
				if problem_dampened {
					return false
				} else {
					problem_dampened = true
					continue
				}
			}
		} else {
			if value >= prev_value || value+3 < prev_value {
				if problem_dampened {
					return false
				} else {
					problem_dampened = true
					continue
				}
			}
		}

		prev_value = value
	}

	return true
}

func getResultP2(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to read file %s. Error: %s", filename, err))
	}

	lines := strings.Split(string(data), "\n")

	num_safe := 0

	for _, line := range lines {
		values := strings.Split(line, " ")

		if !isSafe(values, false) && !isSafe(values[1:], true) && !isSafe(append(values[:1], values[2:]...), true) {
			continue
		}

		num_safe++
	}

	return num_safe
}
