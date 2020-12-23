package helpers

import (
	"fmt"
	"strconv"
)

func StringToFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return f
}

func Float64ToString(f float64) string {
	return fmt.Sprintf("%f", f)
}
