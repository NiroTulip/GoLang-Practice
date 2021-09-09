/* 
Find the first number between 1 and n, inclusive, that is a multiple of c but NOT a multiple of d.
Input data
Enter 3 natural numbers n, c, d, each of which does not exceed 10000.
Output
Print the first number from 1 to n inclusive, a multiple of c, but NOT a multiple of d. If there is no such number, you do not need to display anything.
Sample Input:
twenty
3
5
Sample Output:
3 
*/


/*
Найдите первое число от 1 до n включительно, кратное c, но НЕ кратное d.
Входные данные
Вводится 3 натуральных числа n, c, d, каждое из которых не превышает 10000.
Выходные данные
Вывести первое число от 1 до n включительно, кратное c, но НЕ кратное d. Если такого числа нет - выводить ничего не нужно.
Sample Input:
20
3
5
Sample Output:
3
*/

package main
import "fmt"
func main() {
    var n,c,d int
    fmt.Scan(&n)
    fmt.Scan(&c)
    fmt.Scan(&d)
    
    for i:=1; i<=n; i++{
        if i % c != 0{
            continue
        }
        if i % d != 0 {
        fmt.Println(i)
        break
        }
        
    }
  
}
