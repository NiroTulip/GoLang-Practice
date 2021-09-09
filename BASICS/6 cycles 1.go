/* 
Write a program that in a sequence of numbers finds the sum of two-digit numbers divisible by 8. The program in the first line receives as input the number n - the number of numbers in the sequence, in the second line - n numbers included in this sequence.
Sample Input:
5
38 24 800 8 16
Sample Output:
40 
*/


/*
Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8. Программа в первой строке получает на вход число n - количество чисел в последовательности, во второй строке -- n чисел, входящих в данную последовательность.
Sample Input:
5
38 24 800 8 16
Sample Output:
40
*/

package main
import "fmt"
func main() {
    var a,b,sum int
    fmt.Scan(&a)
  
    for i:=1; i <= a; i++{
       fmt.Scan(&b)
       if b>=10 && b<=99 && b%8==0{
         sum+=b;  
       }
    }
    fmt.Println(sum) 
}
