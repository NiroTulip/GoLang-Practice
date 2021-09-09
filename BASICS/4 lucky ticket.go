/* 
Determine if the ticket is lucky. A ticket is considered lucky if in its six-digit number the sum of the first three digits coincides with the sum of the last three.
Input data format
The ticket number is entered at the entrance - one six-digit number.
Output data format
Print "YES" if the ticket is lucky, otherwise - "NO".
Sample Input:
613244
Sample Output:
YES 
*/


/*
Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.
Формат входных данных
На вход подается номер билета - одно шестизначное  число.
Формат выходных данных
Выведите "YES", если билет счастливый, в противном случае - "NO".
Sample Input:
613244
Sample Output:
YES
*/

package main
import "fmt"
func main() {
var n int 
fmt.Scan(&n)
var s1 = (n%1000000)/100000 + (n%100000)/10000 + (n%10000)/1000
var s2 = n%10 + (n%100)/10 + (n%1000)/100 
    if s1==s2 {
        fmt.Println("YES")
    } else {fmt.Println("NO")}
}
