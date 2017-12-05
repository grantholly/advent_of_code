package main

import(
	"fmt"
	"os"
)


func unpack(n int) []int {
	l := []int{}

	for {
		if n <= 0 {
			break
		}
		l = append(l, n % 10)
		n /= 10
	}
	return l
}

func reverse(l []int) []int {
	r := make([]int, len(l))

	for i, j := 0, len(l) - 1; i < j; 
	i, j = i + 1, j - 1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}

func get_pairs(l []int) [][]int {
	pairs := [][]int{}

	for i, v := range l {
		fmt.Println(i)
		fmt.Println(v)
	}
	return pairs
}

func main () {
	f, err := os.Open("/home/gholly/advent_of_code/day1.txt")
	if err != nil {
		fmt.Println("couldn't open file")
	}
	defer f.Close()
	
	nums := unpack(123334)
//	nums = reverse(nums)
	fmt.Println(nums)
//	get_pairs(nums)
}
