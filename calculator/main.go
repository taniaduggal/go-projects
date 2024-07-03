package main

import "fmt"

func addition(a int, b int) int {
	return a + b
}
func subs(a, b int) int {
	return a - b
}
func product(a, b int) int {
	return a * b
}
func quo(a, b int) int {
	return b / a
}
func main() {
	var choice int
	var a, b int
	for {
		fmt.Println("enter 1 for addition, 2 for subtraction, 3 for product, 4 for quotient, 5 for exit")
		fmt.Scanln(&choice)

		if choice == 5 {
			break
		}
		fmt.Println("enter two numbers")
		fmt.Scanln(&a, &b)

		if choice == 1 {
			add := addition(a, b)
			fmt.Println("The sum is = ", add)
		} else if choice == 2 {
			sub := subs(a, b)
			fmt.Println("The difference is = ", sub)
		} else if choice == 3 {
			prod := product(a, b)
			fmt.Println("The product is = ", prod)
		} else if choice == 4 {
			quot := quo(a, b)
			fmt.Println("The quotient is = ", quot)
		} else {
			fmt.Println("Invalid choice")
		}
	}
}
