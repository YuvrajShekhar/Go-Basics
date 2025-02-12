package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	balancebyte, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to find file")
	}

	balanceString := string(balancebyte)
	balanceFloat, err := strconv.ParseFloat(balanceString, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}

	return balanceFloat, nil
}

func WriteFloatToFile(value float64, fileName string) {
	balance_text := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(balance_text), 0644)
}
