
/* 
The Fibonacci sequence is defined as follows: φ1 = 1, φ2 = 1, φn = φn-1 + φn-2 for n> 1. The beginning of the Fibonacci series looks like this: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ...
Write a function that, given a natural n, returns φn.
Input data
One number n is entered.
Output
It is necessary to output the value of φn.
Sample Input:
4
Sample Output:
3 
*/


/*
Последовательность Фибоначчи определена следующим образом: φ1=1, φ2=1, φn=φn-1+φn-2 при n>1. Начало ряда Фибоначчи выглядит следующим образом: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ...
Напишите функцию, которая по указанному натуральному n возвращает φn.
Входные данные
Вводится одно число n.
Выходные данные
Необходимо вывести  значение φn.
Sample Input:
4
Sample Output:
3
*/


package main

import "fmt"

func fibonacci(n int) int {
	f1 :=1
	f2 :=1
	var fi int
	if (n == 1 || n == 2) {return 1
	} else {
	for i:=3; i<=n; i++{
		fi = f1 + f2
		if (i == n) {
			return fi
		} else {
			f1 = f2
			f2 = fi
		}
	}
	}
	return fi
}

func main() {
fmt.Println(fibonacci(6))
}
