package main

import "fmt"

func main() {
	var n int64

	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	result := Sum(n)
	fmt.Println(result)
}

func Sum(n int64) string {
	var sum int64
	var result string
	for i := int64(1); i <= n; i++ {
		if i%7 != 0 {
			sum += i
			result += fmt.Sprintf("%d+", i)
		}
	}
	result = result[:len(result)-1]
	result += fmt.Sprintf("=%d", sum)
	return result
}
