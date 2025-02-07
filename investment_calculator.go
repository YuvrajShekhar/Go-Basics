package main

import (
	"fmt"
	"math"
)

func main() {
	investAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
	futureValue := investAmount * math.Pow((1+expectedReturnRate/100), years)
	fmt.Println(futureValue)
}
