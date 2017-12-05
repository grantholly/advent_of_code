package main

import(
	"fmt"
	"os"
	"strings"
	"strconv"
	"bufio"
)

func sum_pairs(in string, step int) int {
	sub_strings := strings.Split(in, "")
	var sum int
	
	for i := 0; i < len(sub_strings); i++ {
		ss := sub_strings[i]
		idx := i + step
		// don't go out of bounds here
		if idx >= len(in) {
			idx -= len(in)
		}
		// use the safe index value
		if ss == sub_strings[idx] {
			i, _ := strconv.Atoi(ss)
			sum += i
		}
	}
	return sum
}

func main () {
	f, err := os.Open("/home/gholly/github/advent_of_code/day1.txt")
	if err != nil {
		fmt.Println("couldn't open file")
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	f_input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("couldn't read string")
	}
	fmt.Println(sum_pairs(strings.TrimSpace(f_input), 1))

//	test := "311563"
//	fmt.Println(sum_pairs(test, 1))

	from_the_middle := len(strings.TrimSpace(f_input)) / 2
	// walk up by the length of half of the slice
	fmt.Println(sum_pairs(strings.TrimSpace(f_input), from_the_middle))
}
