package main

import (
	"fmt"
	"strconv"
)

const (
	nope     = "В римской системе числа только до X включительно."
	differen = "используются одновременно разные системы счисления"
	negative = "в римской системе нет отрицательных чисел."
	zero     = "на ноль делить нельзя"
	NotMath  = "так как строка не является математической операцией."
	ranges   = "калькулятор принимает только значения от 0 до 10."
)

func romanToInt(s string) int {
	romanMap := map[string]int{
		"X":    10,
		"IX":   9,
		"VIII": 8,
		"VII":  7,
		"VI":   6,
		"V":    5,
		"IV":   4,
		"III":  3,
		"II":   2,
		"I":    1}

	r, ok := romanMap[s]
	if !ok {
		r = 0
	}

	return r
}
func intToRoman(s int) string {
	intMap := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X"}

	r, ok := intMap[s]
	if !ok {

	}

	return r
}
func main() {
	for {
		var x, y, z string
		var q int
		fmt.Scanf("%s %s %s", &x, &z, &y)

		a, err := strconv.Atoi(x)
		if err != nil {
			a = romanToInt(x)
		}
		{
		}

		b, err := strconv.Atoi(y)
		if err != nil {
			b = romanToInt(y)
		}
		{
		}
		if romanToInt(x) != 0 && romanToInt(y) == 0 || romanToInt(x) == 0 && romanToInt(y) != 0 {
			panic(differen)
		} else {
			if romanToInt(x) != 0 && romanToInt(y) != 0 {
				if romanToInt(x) < 0 || romanToInt(x) > 10 {
					panic(ranges)
				}
				{
				}
				if romanToInt(y) < 0 || romanToInt(y) > 10 {
					panic(ranges)
				}
				{
				}
				switch z {

				case "+":
					{
						q = a + b
						if q > 10 {
							panic(nope)
						}
						{
							println(intToRoman(q))
						}
					}

				case "-":
					{
						q = a - b
						if q < 0 {
							panic(negative)
						}
						println(intToRoman(q))
					}

				case "*":
					{
						q = a * b
						if q > 10 {
							panic(nope)
						}
						{
							println(intToRoman(q))
						}
					}

				case "/":
					{
						if b == 0 {
							panic(zero)
						} else {
							q = a / b
							println(intToRoman(q))
						}
					}
				default:
					panic(NotMath)

				}
			} else {
				if a < 0 || a > 10 {
					panic(ranges)
				}
				{
				}
				if b < 0 || b > 10 {
					panic(ranges)
				}
				{
				}
				switch z {

				case "+":
					{
						q = a + b
						println(q)
					}

				case "-":
					{
						q = b - a
						println(q)
					}

				case "*":
					{
						q = a * b
						println(q)
					}

				case "/":
					{
						if b == 0 {
							panic(zero)
						} else {

							q = a / b
							println(q)
						}

					}
				default:
					panic(NotMath)
				}

			}

		}
	}
}
