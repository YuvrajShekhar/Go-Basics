package recursion

import "fmt"

func main() {
	output := factorial(5)
	fmt.Println(output, "========output")
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}

	return number * factorial(number-1)
}
