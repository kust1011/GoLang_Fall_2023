package main

import (
	"fmt"
	"math/big"
	"syscall/js"
)

func CheckPrime(this js.Value, args []js.Value) interface{} {
	// TODO: Check if the number is prime

	// Get the input value from the HTML input element
	inputVal := js.Global().Get("document").Call("getElementById", "value").Get("value").String()
	fmt.Println("input value: " + inputVal)

	// Convert the input value to big.Int
	number := new(big.Int)
	number, ok := number.SetString(inputVal, 10)
	if !ok {
		js.Global().Get("document").Call("getElementById", "answer").Set("innerText", "Invalid input")
		return nil
	}

	// Check if the number is prime
	isPrime := number.ProbablyPrime(0)

	// Display the result in the HTML paragraph element
	resultText := "It's not prime"
	if isPrime {
		resultText = "It's prime"
	}
	js.Global().Get("document").Call("getElementById", "answer").Set("innerText", resultText)

	return nil

}

func registerCallbacks() {
	// TODO: Register the function CheckPrime
	js.Global().Set("CheckPrime", js.FuncOf(CheckPrime))
}

func main() {
	fmt.Println("Golang main function executed")
	registerCallbacks()

	//need block the main thread forever
	select {}
}
