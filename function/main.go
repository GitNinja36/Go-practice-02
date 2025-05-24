package main

import "fmt"

func main() {
	fmt.Println("hii there")
	greeter()
	greeterTwo()

	result := adder(3, 5)
	fmt.Println("Result is :", result)

	proResult, myMsg := proAdder(3, 6, 9)
	fmt.Println("pro Result is :", proResult)
	fmt.Println("pro msg is :", myMsg)
}

func greeter() {
	fmt.Println("Im rohit")
}

func greeterTwo() {
	fmt.Println("Im niitian")
}

func adder(a int, b int) int {
	sum := a + b
	return sum
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total = total + val
	}
	return total, "so the value is calculated"
}
