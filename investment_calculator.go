package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investAmount float64 = 1000
	expectedReturnRate := 5.5
	var years float64

	fmt.Print("Enter investment amount : ")
	fmt.Scan(&investAmount)

	fmt.Print("Enter years : ")
	fmt.Scan(&years)

	fmt.Print("Enter expectedReturnRate : ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
