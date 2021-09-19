/* 
Remove the given digit from a natural number.
Input data
A natural number and a digit to be deleted are entered.
Output
Print a number without the specified digits.
Input example:
38012732
3
Output example:
801272 
*/


/*
Из натурального числа удалить заданную цифру.
Входные данные
Вводятся натуральное число и цифра, которую нужно удалить.
Выходные данные
Вывести число без заданных цифр.
Sample Input:
38012732
3
Sample Output:
801272
 */


package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var n, dig, resint int
	var res string
	fmt.Scan(&n)
	fmt.Scan(&dig)
	ns:=strconv.Itoa(n)

	for i:=0; i<len(ns); i++ {
		symb:=string(ns[i]);
		if (symb != strconv.Itoa(dig)) {
			res+=symb;
		}
	}
	resint, err := strconv.Atoi(res)
	if err != nil {
		log.Fatal(err)
	}
    fmt.Println(resint)
}
