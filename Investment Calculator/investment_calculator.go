package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	var investAmount float64 = 1000
	expectedReturnRate := 5.5
	var years float64

	outputText("Enter investment amount : ")
	// fmt.Print("Enter investment amount : ")
	fmt.Scan(&investAmount)

	outputText("Enter years : ")
	// fmt.Print("Enter years : ")
	fmt.Scan(&years)

	outputText("Enter expectedReturnRate : ")
	// fmt.Print("Enter expectedReturnRate : ")
	fmt.Scan(&expectedReturnRate)

	// futureValue := investAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	futureValue, futureRealValue := CalculateFutureValues(investAmount, expectedReturnRate, years)

	FormatedFV := fmt.Sprintf("Future Value : %0.1f\n", futureValue)
	FormatedRFV := fmt.Sprintf("Future Real Value : %0.1f\n", futureRealValue)

	// fmt.Println(futureValue)
	// fmt.Println(futureRealValue)
	// fmt.Printf("Future Value : %0.1f\nFuture Real Value : %0.1f\n", futureValue, futureRealValue)
	fmt.Print(FormatedFV, FormatedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func CalculateFutureValues(investAmount, expectedReturnRate, years float64) (float64, float64) {
	futureValue := investAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, futureRealValue
}
