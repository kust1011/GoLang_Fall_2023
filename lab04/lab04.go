package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// TODO: Create a struct to hold the data sent to the template

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func Calculator(w http.ResponseWriter, r *http.Request) {
	// TODO: Finish this function

	// Parse the query string
	query := r.URL.Query()

	// Get the operation
	operation := query.Get("op")

	// Get the numbers
	num1, err := strconv.Atoi(query.Get("num1"))
	if err != nil {
		_ = template.Must(template.ParseFiles("error.html")).Execute(w, nil)
		return
	}
	num2, err := strconv.Atoi(query.Get("num2"))
	if err != nil {
		_ = template.Must(template.ParseFiles("error.html")).Execute(w, nil)
		return
	}

	var data struct {
		Expression string
		Result string
	}
	
	switch operation {
		case "add":
			data.Expression = fmt.Sprintf("%d + %d", num1, num2)
			data.Result = fmt.Sprintf("%d", num1+num2)
		case "sub":
			data.Expression = fmt.Sprintf("%d - %d", num1, num2)
			data.Result = fmt.Sprintf("%d", num1-num2)
		case "mul":
			data.Expression = fmt.Sprintf("%d * %d", num1, num2)
			data.Result = fmt.Sprintf("%d", num1*num2)
		case "div":
			if num2 == 0 {
				_ = template.Must(template.ParseFiles("error.html")).Execute(w, nil)
				return

			} else {
				data.Expression = fmt.Sprintf("%d / %d", num1, num2)
				data.Result = fmt.Sprintf("%d", num1/num2)
			}
		case "gcd":
			data.Expression = fmt.Sprintf("GCD(%d, %d)", num1, num2)
			data.Result = fmt.Sprintf("%d", gcd(num1, num2))
		case "lcm":
			data.Expression = fmt.Sprintf("LCM(%d, %d)", num1, num2)
			data.Result = fmt.Sprintf("%d", num1*num2/gcd(num1, num2))
		default:
			_ = template.Must(template.ParseFiles("error.html")).Execute(w, nil)
			return
		}

	err = template.Must(template.ParseFiles("index.html")).Execute(w, data)

	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	http.HandleFunc("/", Calculator)
	log.Fatal(http.ListenAndServe(":8084", nil))
}

