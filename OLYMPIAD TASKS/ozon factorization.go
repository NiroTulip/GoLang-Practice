/* 
OLYMPIAD TASK FROM OZON.RU FROM ROUTE 256 CONTEST
You are given an integer x.
Can you get x by summing some number of 11, 111, 1111, 11111, ...? (You can use any number among them any number of times).
For instance,
    33=11+11+11
    144=111+11+11+11
Input data format
The first input line contains a single integer
t is the number of test cases.
The first and only line of each test case contains one integer
x is the number you should get.
Output format
For each test case, you must output one line. If you can get x, print "YES" (without quotes). Otherwise print "NO".
You can output each letter from "YES" and "NO" in any case (upper or lower).

Examples
Input data:
4
33
144
69
9

Output:
YES
YES
NO
NO
*/

/*
ОЛИМПИАДНАЯ ЗАДАЧА ОТ OZON.RU С КОНТЕСТА ROUTE 256
Вам дано целое число x.
Можете ли вы получить x, просуммировав некоторое количество 11, 111, 1111, 11111, …? (Вы можете использовать любое число среди них любое количество раз).
Например,
    33=11+11+11
    144=111+11+11+11
Формат входных данных
Первая строка ввода содержит одно целое число 
t — количество наборов входных данных.
Первая и единственная строка каждого набора входных данных содержит одно целое число 
x — число, которое вы должны получить.
Формат выходных данных
Для каждого набора входных данных вы должны вывести одну строку. Если вы можете получить x, выведите «YES» (без кавычек). В противном случае выведите «NO».
Вы можете вывести каждую букву из «YES» и «NO» в любом регистре (верхнем или нижнем).

Примеры
Входные данные:
4
33
144
69
9
Выходные данные:
YES
YES
NO
NO
*/

package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

var FactorizationSuccess = false

//Вспомогательная функция возвращает одно число вида 11, 111, 1111, ... указанной в аргументе длины.
func getOneSequenceNumber(bitness int) int{
	str := strings.Repeat("1", bitness)
	result, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

//Функция возвращает срез из всех возможных кандидатов в слагаемые вида 11, 111, 1111, ... для данного числа, т.е. все
//числа вида 11, 111, 1111, ..., которые меньше заданного числа
func getListOfdivisors(div int) [] int{
	var result [] int
	if (div < 11) {return result}
	bitness := int(math.Log10(float64(div)))+1
	if (div >= getOneSequenceNumber(bitness)) {result = append(result, getOneSequenceNumber(bitness))}
	for i:=bitness-1; i>=2; i-- {
    	result = append(result, getOneSequenceNumber(i))
	}
return result
}

//Рекурсивная функция путем деления на множители вида 11, 111, 1111, ... проверяет все возможные комбинации разложения
//заданного числа на множители, кратные числам вида 11, 111, 1111, ..., пока не найдет нужное разложение (или не найдет)
//Используется разложение на множители, т.к. если число делится нацело на одно из чисел вида 11.., значит оно может
//быть представлено в виде некоторого количества таких слагаемых. Рекурсивно проверяя остатки от деления на делимость
//для разных чисел вида 11.. мы решаем поставленную задачу без сложения и вычитания.
func checkFactorization(div int) {
	if (div < 11) {return}
	listOfDivisors :=  getListOfdivisors(div)
	for _, divisor := range listOfDivisors {
		remainder := div % divisor
		if (remainder % 11 == 0) {FactorizationSuccess = true} else{
			if (remainder > 11) {checkFactorization(remainder)}
		}
	}
}


func main() {
	var t,d int
	var digits [] int
	fmt.Scan(&t)
	for i:=0; i<t; i++{
		fmt.Scan(&d)
		digits = append (digits, d)
	}
	for _,digit := range digits{
		checkFactorization(digit)
		if FactorizationSuccess {fmt.Println("YES")} else {
			fmt.Println("NO")
		}
	FactorizationSuccess = false
	}
}

