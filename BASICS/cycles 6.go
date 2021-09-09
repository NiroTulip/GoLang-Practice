/* 
Find the largest multiple of 7 between a and b.
Input data
Enter two integers a and b (a≤b).
Output
Find the largest number on the segment from a to b (the segment includes numbers a and b), a multiple of 7, or print "NO" - if there are none.
Sample Input:
100
500
Sample Output:
497 
*/

/*
Найдите самое большее число на отрезке от a до b, кратное 7 .
Входные данные 
Вводится два целых числа a и b (a≤b).
Выходные данные 
Найдите самое большее число на отрезке от a до b (отрезок включает в себя числа a и b), кратное 7 , или выведите "NO" - если таковых нет.
Sample Input:
100
500
Sample Output:
497
*/

package main
import "fmt"
func main() {
    var a,b,res int
    fmt.Scan(&a, &b)
    res = 1
    for i:=b; i>=a; i--{
        if i%7 == 0 {
            res = i
            break
        }
    }
       if res != 1 {
            fmt.Print(res)
        } else {
            fmt.Print("NO")
        }

} 
