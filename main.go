package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type elements struct {
	num_1, num_2 int
	op           string
	isRoman      bool
}
type key_value struct {
	Key   int
	Value string
}

func Decode(roman string) int {
	var sum int
	var Roman = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	for k, v := range roman {
		if k < len(roman)-1 && Roman[byte(roman[k+1])] > Roman[byte(roman[k])] {
			sum -= Roman[byte(v)]
		} else {
			sum += Roman[byte(v)]
		}
	}
	return sum
}

func NumCheck(num string) (int, bool) {
	var int_num int
	var err error
	var is_roman_num bool = false
	if num != "" {
		if num[0] == 73 || num[0] == 86 || num[0] == 88 {
			is_roman_num = true
			int_num = Decode(num)
		} else {
			int_num, err = strconv.Atoi(num)
			if err != nil {
				fmt.Println("Паника, с числами что-то не так")
				panic(err)
			}
		}
		if int_num < 0 || int_num > 10 {
			fmt.Println("Panicking!")
			panic(fmt.Sprintf("Паника, Число или формат не удовлетворяют заданию"))
		}
	} else {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("Паника, с числами что-то не так"))
	}
	return int_num, is_roman_num
}

func strToNum(s string) elements {
	var num_1, num_2, op string = "", "", ""
	var isRoman bool
	i := 0
	for c := len(s); i < c; i++ {
		if s[i] > 47 && s[i] < 58 {
			num_1 += string(s[i])
		} else if s[i] == 73 || s[i] == 86 || s[i] == 88 {
			num_1 += string(s[i])
		} else if s[i] == 42 || s[i] == 43 || s[i] == 45 || s[i] == 47 {
			op = string(s[i])
			break
		}
	}

	i++
	if op != "" {
		for c := len(s); i < c; i++ {
			if s[i] > 47 && s[i] < 58 {
				num_2 += string(s[i])
			} else if s[i] == 73 || s[i] == 86 || s[i] == 88 {
				num_2 += string(s[i])

			}
		}
	}
	int_num_1, is_roman_num_1 := NumCheck(num_1)
	int_num_2, is_roman_num_2 := NumCheck(num_2)
	if is_roman_num_1 != is_roman_num_2 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("Паника, используются одновременно разные системы счисления."))
	}
	if is_roman_num_1 == true {
		isRoman = true
	} else {
		isRoman = false
	}
	v := elements{int_num_1, int_num_2, op, isRoman}
	return v
}
func to_Roman(num int) string {
	var roman string = ""
	var roman_num [9]string
	roman_num[0] = "C"
	roman_num[1] = "XC"
	roman_num[2] = "L"
	roman_num[3] = "XL"
	roman_num[4] = "X"
	roman_num[5] = "IX"
	roman_num[6] = "V"
	roman_num[7] = "IV"
	roman_num[8] = "I"

	var arabic_num [9]int
	arabic_num[0] = 100
	arabic_num[1] = 90
	arabic_num[2] = 50
	arabic_num[3] = 40
	arabic_num[4] = 10
	arabic_num[5] = 9
	arabic_num[6] = 5
	arabic_num[7] = 4
	arabic_num[8] = 1

	for num > 0 {
		for key := range arabic_num {
			for num >= arabic_num[key] {
				roman += roman_num[key]
				num -= arabic_num[key]
			}
		}

	}
	return roman
}
func calc(e elements) {
	var calcs int8
	if e.isRoman == false {
		switch e.op {
		case "+":
			calcs = int8(e.num_1) + int8(e.num_2)
		case "-":
			calcs = int8(e.num_1) - int8(e.num_2)
		case "*":
			calcs = int8(e.num_1) * int8(e.num_2)
		case "/":
			calcs = int8(e.num_1) / int8(e.num_2)
		default:
			fmt.Println("Panicking!")
			panic(fmt.Sprintf("Паника, с оператором что-то не так"))
		}
		fmt.Println(calcs)
	} else {
		switch e.op {
		case "+":
			calcs = int8(e.num_1) + int8(e.num_2)
		case "-":
			calcs = int8(e.num_1) - int8(e.num_2)
		case "*":
			calcs = int8(e.num_1) * int8(e.num_2)
		case "/":
			calcs = int8(e.num_1) / int8(e.num_2)
		default:
			fmt.Println("Panicking!")
			panic(fmt.Sprintf("Паника, с оператором что-то не так"))
		}
		if calcs >= 0 {
			fmt.Println(to_Roman(int(calcs)))
		} else {
			fmt.Println("Panicking!")
			panic(fmt.Sprintf("Паника, в римской системе нет отрицательных чисел."))
		}

	}

}

func main() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	calc(strToNum(s))
}
