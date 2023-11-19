package main

import "fmt"

func main() {
	theSmallestNumber([]int{348,2198,-3478,384739,-3820})
}

func theSmallestNumber(values []int) {
	min := values[0]

	for _, v := range values {
		if v < min {
			min = v
		}
	}
	fmt.Println(min)
}
