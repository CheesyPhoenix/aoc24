package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println(getResult("./src/day3/input.txt"))
	fmt.Println(getResultP2("./src/day3/input.txt"))
}

func parseInt(str string) int {
	out, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("Unable to parse %s as int. Error: %s", str, err))
	}
	return out
}

func getResult(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to read file %s. Error: %s", filename, err))
	}

	r, err := regexp.Compile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to compile regex. Error: %s", err))
	}

	res := r.FindAllStringSubmatch(string(data), -1)

	sum := 0

	for _, match := range res {
		sum += parseInt(match[1]) * parseInt(match[2])
	}

	return sum
}

func getResultP2(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to read file %s. Error: %s", filename, err))
	}

	mul_r, err := regexp.Compile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	if err != nil {
		panic(fmt.Sprintf("Got error trying to compile regex. Error: %s", err))
	}

	mul_res := mul_r.FindAllStringSubmatchIndex(string(data), -1)

	sum := 0

	input := string(data)
	for _, match := range mul_res {
		dont_index := strings.LastIndex(input[0:match[0]], "don't()")
		do_index := strings.LastIndex(input[0:match[0]], "do()")
		if do_index < dont_index {
			continue
		}

		sum += parseInt(input[match[2]:match[3]]) * parseInt(input[match[4]:match[5]])
	}

	return sum
}
