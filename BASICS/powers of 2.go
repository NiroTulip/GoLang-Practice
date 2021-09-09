/* 
For a given number N, print all integer values of a power of two not exceeding N in ascending order.
Input data
A natural number is entered.
Output
Print the answer to the problem.
Sample Input:
50
Sample Output:
1 2 4 8 16 32 
*/

/*
По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.
Входные данные
Вводится натуральное число.
Выходные данные
Выведите ответ на задачу.
Sample Input:
50
Sample Output:
1 2 4 8 16 32
*/


package main
import "fmt"
func main() {
    var n,s2 int
    fmt.Scan(&n)
    s2=1
for {
    if s2 <= n {
      fmt.Print(s2, " ")
      s2=s2*2
      } else {break}
    }
      
} 
