// Sort My Numbers

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func AscendingAndDescendingOrder() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		number, _ := strconv.Atoi(line)
		numbers = append(numbers, number)
	}

	fmt.Println("---Ascending Order--")
	sort.Ints(numbers)
	fmt.Println(strings.Trim(fmt.Sprint(numbers), "[]"))

	fmt.Println("---Descending Order---")
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Println(strings.Trim(fmt.Sprint(numbers), "[]"))
}

func main() {
	AscendingAndDescendingOrder()

}
