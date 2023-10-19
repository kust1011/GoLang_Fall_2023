package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Calculator(w http.ResponseWriter, r *http.Request) {
	// TODO: implement a calculator
	
	// Get the path from the URL
	path := r.URL.Path

	parts := strings.Split(path, "/")
	if  len(parts) != 4 {
		fmt.Fprint(w, "Error!")
		return
	}

	operation := parts[1]

	num1, err := strconv.Atoi(parts[2])
	if err != nil {
		fmt.Fprint(w, "Error!")
		return
	}
	num2, err := strconv.Atoi(parts[3])
	if err != nil {
		fmt.Fprint(w, "Error!")
		return
	}

	// Calculate the result
	var result string
	switch operation {
		case "add":
			result = fmt.Sprintf("%d + %d = %d", num1, num2, num1+num2)
		case "sub":
			result = fmt.Sprintf("%d - %d = %d", num1, num2, num1-num2)
		case "mul":
			result = fmt.Sprintf("%d * %d = %d", num1, num2, num1*num2)
		case "div":
			if num2 == 0 {
				result = "Error!"
			} else {
				result = fmt.Sprintf("%d / %d = %d, reminder = %d", num1, num2, num1/num2, num1%num2)
			}
		default:
			result = "Error!"
		}

	fmt.Fprint(w, result)

}

func main() {
	http.HandleFunc("/", Calculator)
	log.Fatal(http.ListenAndServe(":8083", nil))
}
