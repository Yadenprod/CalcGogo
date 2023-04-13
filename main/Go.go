package main

import (
	"fmt"
	"strconv"
)

type Calculator struct {
	num1, num2 int
	oper       string
}

func (c *Calculator) SetNumbers(num1, num2 int) {
	c.num1 = num1
	c.num2 = num2
}

func (c *Calculator) SetOperator(oper string) {
	c.oper = oper
}

func (c *Calculator) Calculate() string {
	operatory := [4]string{"+", "-", "*", "/"}
	if c.oper != operatory[0] && c.oper != operatory[1] && c.oper != operatory[2] && c.oper != operatory[3] {
		return "Введите оператор для операции +,-,*,/ любой из них"
	} else {
		switch c.oper {
		case "+":
			return strconv.Itoa(c.num1 + c.num2)
		case "-":
			return strconv.Itoa(c.num1 - c.num2)
		case "*":
			return strconv.Itoa(c.num1 * c.num2)
		case "/":
			return strconv.Itoa(c.num1 / c.num2)
		}
	}
	return "Error"
}

func main() {
	c := Calculator{}
	var num1, num2 int
	var oper string
	fmt.Println("\"Целочисленный Калькулятор\" Введите выражение в одну строку формата 5 * 24")
	fmt.Print("Ввод: ")
	fmt.Scan(&num1, &oper, &num2)
	c.SetNumbers(num1, num2)
	c.SetOperator(oper)
	fmt.Println("Результат: " + c.Calculate())
}
