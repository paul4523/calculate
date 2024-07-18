package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	nope     = "В римской системе числа только от I до X включительно."
	differen = "используются одновременно разные системы счисления"
	negative = "в римской системе нет отрицательных чисел и нуля."
	zero     = "на ноль делить нельзя"
	NotMath  = "так как строка не является математической операцией."
	ranges   = "калькулятор принимает только значения от 1 до 10."
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
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VIII",
		9:   "IX",
		10:  "X",
		11:  "XI",
		12:  "XII",
		13:  "XIII",
		14:  "XIV",
		15:  "XV",
		16:  "XVI",
		17:  "XVII",
		18:  "XVIII",
		19:  "XIX",
		20:  "XX",
		21:  "XXI",
		22:  "XXII",
		23:  "XXIII",
		24:  "XXIV",
		25:  "XXV",
		26:  "XXVI",
		27:  "XXVII",
		28:  "XXVIII",
		29:  "XXIX",
		30:  "XXX",
		31:  "XXXI",
		32:  "XXXII",
		33:  "XXXIII",
		34:  "XXXIV",
		35:  "XXXV",
		36:  "XXXVI",
		37:  "XXXVII",
		38:  "XXXVIII",
		39:  "XXXIX",
		40:  "XL",
		41:  "XLI",
		42:  "XLII",
		43:  "XLIII",
		44:  "XLIV",
		45:  "XLV",
		46:  "XLVI",
		47:  "XLVII",
		48:  "XLVIII",
		49:  "XLIX",
		50:  "L",
		51:  "LI",
		52:  "LII",
		53:  "LIII",
		54:  "LIV",
		55:  "LV",
		56:  "LVI",
		57:  "LVII",
		58:  "LVIII",
		59:  "LIX",
		60:  "LX",
		61:  "LXI",
		62:  "LXII",
		63:  "LXIII",
		64:  "LXIV",
		65:  "LXV",
		66:  "LXVI",
		67:  "LXVII",
		68:  "LXVIII",
		69:  "LXIX",
		70:  "LXX",
		71:  "LXXI",
		72:  "LXXII",
		73:  "LXXIII",
		74:  "LXXIV",
		75:  "LXXV",
		76:  "LXXVI",
		77:  "LXXVII",
		78:  "LXXVIII",
		79:  "LXXIX",
		80:  "LXXX",
		81:  "LXXXI",
		82:  "LXXXII",
		83:  "LXXXIII",
		84:  "LXXXIV",
		85:  "LXXXV",
		86:  "LXXXVI",
		87:  "LXXXVII",
		88:  "LXXXVIII",
		89:  "LXXXIX",
		90:  "XC",
		91:  "XCI",
		92:  "XCII",
		93:  "XCIII",
		94:  "XCIV",
		95:  "XCV",
		96:  "XCVI",
		97:  "XCVII",
		98:  "XCVIII",
		99:  "XCIX",
		100: "C"}

	r, ok := intMap[s]
	if !ok {

	}

	return r
}
func main() {
	for {
		var x, y, z string
		var q int
		var input string
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()
		fields := strings.Fields(input)
		if len(fields) != 3 {
			panic(NotMath)
		}

		x = fields[0]
		z = fields[1]
		y = fields[2]

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
		if a < 1 || a > 10 {
			panic(ranges)
		}
		{
		}
		if b < 1 || b > 10 {
			panic(ranges)
		}
		{
		}
		if romanToInt(x) != 0 && romanToInt(y) == 0 || romanToInt(x) == 0 && romanToInt(y) != 0 {
			panic(differen)
		} else {
			if romanToInt(x) != 0 && romanToInt(y) != 0 {
				if romanToInt(x) < 1 || romanToInt(x) > 10 {
					panic(ranges)
				}
				{
				}
				if romanToInt(y) < 1 || romanToInt(y) > 10 {
					panic(ranges)
				}
				{
				}
				switch z {

				case "+":
					{
						q = a + b

						println(intToRoman(q))

					}

				case "-":
					{
						q = a - b
						if q < 1 {
							panic(negative)
						}
						println(intToRoman(q))
					}

				case "*":
					{
						q = a * b

						println(intToRoman(q))

					}

				case "/":
					{
						if b == 0 {
							panic(zero)
						} else {
							q = a / b
							if q == 0 {
								panic(nope)
							}
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
						q = a - b
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
