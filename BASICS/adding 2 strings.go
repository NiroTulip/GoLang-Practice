/*
You need to write a function that reads two variables of type string , and returns a number of type int64 that you need to get by adding these strings. First you need to remove excess garbage from the strings and convert them to numbers. Garbage means extra characters and special characters.
*/

/*
Необходимо написать функцию, которая считывает две переменных типа string, а возвращает число типа int64 которое нужно получить сложением этих строк. Предварительно вам нужно убрать из строк лишний мусор и конвертировать их в числа. Под мусором имеются ввиду лишние символы и спец знаки.
*/

package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func extractNumber (s string) int64 {
	var res string
	var res_int64 int64
	runes := []rune(s)
	for _,r := range runes{
		if unicode.IsDigit(r) {res = res + string(r)}
	}
	res_int64, err := strconv.ParseInt(res, 10, 64)
	if err != nil {
		panic(err)
	}
	return res_int64
}

func adding(s1, s2 string) int64 {
	var num1, num2 int64
	num1 = extractNumber(s1)
	num2 = extractNumber(s2)
	return num1 + num2
}


func main() {
fmt.Println(adding("s12f", "#23d"))
}

