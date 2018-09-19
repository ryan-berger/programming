package main

import (
	"math"
	"strconv"
		"fmt"
)

const FIRSTINPUT = `2
212
1253`

const BIGGERINPUT = `5
23
17
43
52
22`

const HARDESINPUT = `3
213
102
45`

type Exponent struct {
	base  float64
	power float64
}

func ParseNumber(number string) Exponent  {
	base, _ := strconv.ParseFloat(number[:len(number) - 1], 64)

	power, _ := strconv.ParseFloat(string(number[len(number) - 1]), 64)
	return Exponent{
		base:base,
		power: power,
	}
}

func (exponent Exponent) Eval() int {
	return int(math.Pow(exponent.base, exponent.power))
}



func main() {
	var numInput int
	fmt.Scanln(&numInput)

	result := 0

	for i := 0; i < numInput; i++ {
		var numberString string
		fmt.Scanln(&numberString)
		result += ParseNumber(numberString).Eval()
	}

	fmt.Println(result)
}
