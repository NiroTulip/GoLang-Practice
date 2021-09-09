/* 
For this three-digit number, determine if all of its numbers are different.
Input data format
The input is one natural three-digit number.
Output data format
Print "YES" if all digits of the number are different, otherwise - "NO".
Sample Input 1:
237
Sample Output 1:
YES
Sample Input 2:
117
Sample Output 2:
NO 
*/

/*
По данному трехзначному числу определите, все ли его цифры различны.
Формат входных данных
На вход подается одно натуральное трехзначное число.
Формат выходных данных
Выведите "YES", если все цифры числа различны, в противном случае - "NO".
Sample Input 1:
237
Sample Output 1:
YES
Sample Input 2:
117
Sample Output 2:
NO
*/

package main
import "fmt"
func main() {
var n int 
fmt.Scan(&n)
var a int = n % 10   
var b int = (n % 100) / 10
var c int = n/100 
    if (a!=b && a!=c && b!=c){
    fmt.Print("YES")   
    } else {fmt.Print("NO")}  
}
