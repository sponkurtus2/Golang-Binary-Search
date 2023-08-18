package main

import (
	"fmt"
	"os"
)

func main() {
	numbers := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	lowNum := 0
	highNum := len(numbers) -1
	target := 3	
	
	for lowNum <= highNum {
		midNum := (lowNum + highNum) / 2
		fmt.Printf("\n lowNum: %d \n highNum: %d \n midNum %d \n", lowNum, highNum, midNum)

		if numbers[midNum] == target {
			fmt.Printf("\n == lowNum: %d \n highNum: %d \n midNum %d \n", lowNum, highNum, midNum)
			fmt.Println("Index found at ", midNum)
			os.Exit(0)
		} else if numbers[midNum] > target {
			highNum = midNum - 1
			fmt.Printf("\n >> lowNum: %d \n highNum: %d \n midNum %d \n", lowNum, highNum, midNum)
		} else {
			lowNum = midNum + 1
			fmt.Printf("\n << lowNum: %d \n highNum: %d \n midNum %d \n", lowNum, highNum, midNum)
		}
	}

}