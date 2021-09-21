/* 
Write a function sumInt that takes a variable number of return int arguments, the number of arguments returned by the function, and their sum.
An example of calling your function:
a, b: = sumInt (1, 0)
fmt.Println (a, b)
Result: 2, 1 
*/


/*
Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных функцией аргументов и их сумму. 
Пример вызова вашей функции:
a, b := sumInt(1, 0)
fmt.Println(a, b)
Результат: 2, 1
*/

package main

import "fmt"

func sumInt(x ... int) (int, int) {
 var sum, kol int
 for idx, elem := range x {
 	sum += elem
 	kol = idx + 1
 }
 return kol, sum
}

func main() {
a,b := sumInt(1,2,3,4)
fmt.Println(a,b)
}
