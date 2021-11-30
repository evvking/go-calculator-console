package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("\nGo-Kalkulator v1.0 napisany w Go")

	var x, y int
	var sign string

	fmt.Print("\nPierwsza liczba: ")
	fmt.Scanln(&x)

	fmt.Print("Znak działania (+, -, *, /, ^, sq): ")
	fmt.Scanln(&sign)

	fmt.Print("Druga liczba (nieuźywana przy sq): ")
	fmt.Scanln(&y)

	var resultInt int
	var resultInt64 int64
	var resultFloat32 float32
	var resultFloat64 float64

	switch sign {
	case "+":
		resultInt = addition(x, y)
	case "-":
		resultInt = subtraction(x, y)
	case "*":
		resultInt = multiplication(x, y)
	case "/":
		resultFloat32 = division(x, y)
	case "^":
		resultInt64 = power(x, y)
	case "sq":
		resultFloat64 = sqrt(x)
	}

	fmt.Print("\n\nWynik: ")
	if sign == "+" {
		fmt.Printf("Suma liczb %d i %d wynosi %d\n\n", x, y, resultInt)
	} else if sign == "-" {
		fmt.Printf("Różnica liczb %d i %d wynosi %d\n\n", x, y, resultInt)
	} else if sign == "*" {
		fmt.Printf("Iloczyn liczb %d i %d wynosi %d\n\n", x, y, resultInt)
	} else if sign == "/" {
		fmt.Printf("Iloraz liczb %d i %d wynosi %f\n\n", x, y, resultFloat32)
	} else if sign == "^" {
		fmt.Printf("Liczba %d podniesiona do potęgi %d jest równa %d\n\n", x, y, resultInt64)
	} else if sign == "sq" {
		fmt.Printf("Pierwiastek kwadratowy z liczby %d wynosi %f\n\n", x, resultFloat64)
	}
}

func addition(num1, num2 int) int {
	var result int = num1 + num2
	return result
}

func subtraction(num1, num2 int) int {
	var result int = num1 - num2
	return result
}

func multiplication(num1, num2 int) int {
	var result int = num1 * num2
	return result
}

func division(num1, num2 int) float32 {
	var num1_float32, num2_float32 float32

	num1_float32 = float32(num1)
	num2_float32 = float32(num2)

	var result float32 = num1_float32 / num2_float32
	return result
}

func power(num1, num2 int) int64 {
	var num1_64, num2_64 int64

	num1_64 = int64(num1)
	num2_64 = int64(num2)

	var result int64 = num1_64 ^ num2_64
	return result
}

func sqrt(num int) float64 {
	var numFloat float64
	numFloat = float64(num)

	var result = math.Sqrt(numFloat)
	return result
}
