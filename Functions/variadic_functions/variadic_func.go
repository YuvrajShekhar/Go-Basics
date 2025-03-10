package variadic

import "fmt"

func main() {
	totalsum := sumpup(1, 20, 30, 23, 232)
	fmt.Println(totalsum, "============ total sum")
}

func sumpup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum = sum + val
	}

	return sum
}
