package main

import "fmt"

func main() {
	elements := []int{999, 99, 9}

	for i := 0; i < len(elements); i = i + 1 {
		fmt.Print(elements[i] + 1)
		fmt.Print(" ")
	}

	fmt.Println("Done!")

	name := "Mittens"
	weight := 12
	fmt.Printf("The cat is named %s.", name)
	fmt.Printf("Its weight is %d.", weight)

}
