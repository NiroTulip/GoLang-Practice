/* 
The deposit in the bank is x rubles. It increases annually by p percent, after which the fractional part of kopecks is discarded. Each year, the amount of the deposit becomes larger. Determine how many years the contribution will be at least y rubles.
Input data
The program receives three natural numbers as input: x, p, y.
Output
The program should output one integer.
Sample Input:
100
10
200
Sample Output:
8
*/


/*
Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов, после чего дробная часть копеек отбрасывается. Каждый год сумма вклада становится больше. Определите, через сколько лет вклад составит не менее y рублей.
Входные данные
Программа получает на вход три натуральных числа: x, p, y.
Выходные данные
Программа должна вывести одно целое число.
Sample Input:
100
10
200
Sample Output:
8
*/

package main
import "fmt"
func main() {
    var x,p,y int
    var year int = 1
    fmt.Scan(&x)
    fmt.Scan(&p)
    fmt.Scan(&y)
    
    for ; x<y; year++{
        x=x+x*p/100
        if x>=y{
          fmt.Println(year) 
          break
        }
           
    }
 }
