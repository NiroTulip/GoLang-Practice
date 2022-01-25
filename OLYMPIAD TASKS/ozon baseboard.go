/* 
OLYMPIAD TASK FROM OZON.RU FROM ROUTE 256 CONTEST

The store sells an unlimited number of skirting boards with lengths 1,2,…,n. You want to buy some set of rails to get the total length L.
It is allowed that in your set there are slats having the same length. What is the minimum number of rails needed to get the length L?

Input data format
The only line of the input contains two integers n and L
(1 <= n <= 100000, 1 <= L < 1010M)

Output format
Print exactly one integer — the minimum number of rails needed to get the length L.

Examples
Input data:
5 11
Output:
3

The task is more for mathematical thinking, and not for programming, because, if understood correctly, it is implemented in 3 lines of code:
*/


/*
ОЛИМПИАДНАЯ ЗАДАЧА ОТ OZON.RU С КОНТЕСТА ROUTE 256

В магазине продается неограниченное количество плинтусных реек с длинами 1,2,…,n. Вы хотите купить некоторый набор реек, чтобы получить суммарную длину L.
Разрешается чтобы в вашем наборе были рейки имеющие одну и ту же длину. Какое минимальное количество реек нужно, чтобы получить длину L?

Формат входных данных
Единственная строка входных данных содержит два целых числа n и L 
(1 <= n <= 100000, 1 <= L < 1010M)

Формат выходных данных
Выведите ровно одно целое число — минимальное количество реек, которое нужно чтобы получить длину L.

Примеры
Входные данные:
5 11
Выходные данные:
3

Задача скорее на математическое мышление, нежели на программирование, т.к. при правильном понимании, она легко реализуется 3 строками кода:
*/

package main

import "fmt"

var n,L uint64

func getNumberOfPieces(n,L uint64) int{
	if (L<n) {return 1}	else {
		if L % n == 0 {return int(L/n)} else {
			return int(L/n)+1
		}
	}
}

func main() {
 fmt.Scan(&n,&L)
 fmt.Println(getNumberOfPieces(n,L))
}



