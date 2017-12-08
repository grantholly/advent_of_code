package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"sort"
)

func main() {
	f, err := os.Open("./day2.txt")
	if err != nil {
		fmt.Println("can't open file!")
	}
	defer f.Close()

	var part_one_sum, part_two_sum int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "\t")
		sorted_row := make([]int, len(row))
		for i, v := range row {
			sorted_row[i], _ = strconv.Atoi(v)
			for j := 0; j < len(row); j++ {
				if v != row[j] {
					val, _ := strconv.Atoi(v)
					other, _ := strconv.Atoi(row[j])
					if val % other == 0 {
						fmt.Println(val, other)
						part_two_sum += val / other
					}
				}
			}
		}
		sort.Slice(sorted_row, func(i, j int) bool {
			return sorted_row[i] < sorted_row[j]
		})
		part_one_sum += sorted_row[len(sorted_row)-1] - sorted_row[0]
	}
	fmt.Println(part_one_sum)
	fmt.Println(part_two_sum)
}
