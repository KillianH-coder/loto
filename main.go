package main

import (
	"Loto/Drawing" // directory path for generate_numbers.go
	"Loto/Format"  // directory path for format.go
	"Loto/Result"  // directory path for print_result.go
	"fmt"
)

func main() {
	numbers := Drawing.GenerateNumbers() // directory name.function
	result := Format.Formated(numbers)
	finalResult := Result.PrintResult(result, numbers) // directory name.function
	fmt.Println(finalResult)
}
