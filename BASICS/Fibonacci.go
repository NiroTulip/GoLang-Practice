/* 
Fibonacci number
Given a natural number A> 1. Determine the number at Fibonacci series of numbers it is, that is, print a number n such that φn = A. If A is not a Fibonacci number, print the number -1.
Input data
A natural number is entered.
Output
Print the answer to the problem.
Sample Input:
eight
Sample Output:
6 
*/

/*
Номер числа Фибоначчи
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, что φn=A. Если А не является числом Фибоначчи, выведите число -1.
Входные данные
Вводится натуральное число.
Выходные данные
Выведите ответ на задачу.
Sample Input:
8
Sample Output:
6
*/

package main
import "fmt"
func main() {
    var n, f1, f2, num, i int
    fmt.Scan(&n)
    f1=1
    f2=1
    i=1;
    for {
          
        if i==2 {num=3} else {num++}
        if i==n {
           fmt.Print(num)
           break
        }
        if i>n {
           num=-1
           fmt.Print(num)
           break
        }

        i=f1+f2
        f1=f2
        f2=i
    }
   
} 
