package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var roman = map[string]string{
	"C":    "100",
	"XC":   "90",
	"L":    "50",
	"XL":   "40",
	"X":    "10",
	"IX":   "9",
	"VIII": "8",
	"VII":  "7",
	"VI":   "6",
	"V":    "5",
	"IV":   "4",
	"III":  "3",
	"II":   "2",
	"I":    "1",
}

var arabian = map[string]string{
	"100": "C",
	"90":  "XC",
	"50":  "L",
	"40":  "XL",
	"10":  "X",
	"9":   "IX",
	"8":   "VIII",
	"7":   "VII",
	"6":   "VI",
	"5":   "V",
	"4":   "IV",
	"3":   "III",
	"2":   "II",
	"1":   "I",
}

func IsCorrect(s string, arr [4]string) (string, string, string) {
	var verif, mPerem, mOper []string
	var oper string
	verif = strings.Split(s, " ")

	if len(verif) == 3 {
		for i := 0; i < len(arr); i++ {
			if strings.Contains(s, arr[i]) == true {
				mOper = append(mOper, arr[i])
				oper = arr[i]
			}
		}

		text := strings.ReplaceAll(s, " ", "")

		if len(mOper) >= 1 {
			mPerem = strings.Split(text, oper)
			return oper, mPerem[0], mPerem[1]
		}

		panic("Нет ни одного оператора")

	} else {
		panic("Неверный формат")
	}

}

func isRomainFunc(a string, b string, roman map[string]string) (bool, int, int) {
	var romain, arabic []string
	var two string
	for k, v := range roman {
		if len(romain) == 2 || len(arabic) == 2 {
			break
		} else if k == a && k == b {
			romain = append(romain, a, a)
		} else if v == a && v == b {
			arabic = append(arabic, a, a)
		} else if k == a {
			romain = append(romain, a)
			if two != "" {
				romain = append(romain, two)
			}
		} else if k == b {
			two = k
			if len(romain) == 1 {
				romain = append(romain, b)
			}
		} else if v == a {
			arabic = append(arabic, a)
			if two != "" {
				arabic = append(arabic, two)
			}
		} else if v == b {
			two = v
			if len(arabic) == 1 {
				arabic = append(arabic, b)
			}
		}
	}

	if len(arabic) == 2 {
		one, _ := strconv.Atoi(arabic[0])
		two, _ := strconv.Atoi(arabic[1])
		return false, one, two
	} else if len(romain) == 2 {
		one, _ := strconv.Atoi(roman[romain[0]])
		two, _ := strconv.Atoi(roman[romain[1]])
		return true, one, two
	} else {
		panic("Введите числами от 1 до 10 вкл")
	}
}

func cal(a int, b int, oper string, isRomain bool, arabian map[string]string) {
	var itog int
	if a >= 1 && b >= 1 && a <= 10 && b <= 10 {
		switch oper {
		case "+":
			itog = a + b
		case "-":
			itog = a - b
		case "*":
			itog = a * b
		case "/":
			itog = a / b
		}
	} else {
		panic("Введите одинаковые вид значение")
	}

	if isRomain == false {
		fmt.Println(itog)
	} else {
		if itog == 0 {
			panic("В римской системе нет 0")
		} else if itog > 0 {
			if itog >= 1 && itog <= 10 || itog == 20 || itog == 40 || itog == 50 || itog == 90 || itog == 100 {
				fmt.Print(arabian[strconv.Itoa(itog)])
			} else if itog > 11 && itog < 40 {
				fmt.Println(strings.Repeat("X", itog/10) + arabian[strconv.Itoa(itog%10)])
			} else if itog > 40 && itog < 50 {
				fmt.Println("XL" + arabian[strconv.Itoa(itog%10)])
			} else if itog > 50 && itog < 90 {
				fmt.Println("L" + strings.Repeat("X", itog/10-5) + arabian[strconv.Itoa(itog%10)])
			} else if itog > 90 && itog < 99 {
				fmt.Println("XC" + arabian[strconv.Itoa(itog%10)])
			}
		} else {
			panic("В римской числе нет отрицательных чисел")
		}
	}

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите значения:")
	operator := [4]string{"+", "-", "/", "*"}
	text, _ := reader.ReadString('\n')
	text = strings.ToUpper(strings.TrimSpace(text))
	oper, a, b := IsCorrect(text, operator)
	isRomain, aInt, bInt := isRomainFunc(a, b, roman)
	cal(aInt, bInt, oper, isRomain, arabian)
}
