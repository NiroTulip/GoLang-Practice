/* 
A three-digit number is given. Flip it over and then take it out.
Input data format
The input is a three-digit number that does not end in zero.
Output data format
Print the inverted number.
Sample Input:
843
Sample Output:
348 
*/

/*
Дано трехзначное число. Переверните его, а затем выведите. 
Формат входных данных
На вход дается трехзначное число, не оканчивающееся на ноль.
Формат выходных данных
Выведите перевернутое число.
Sample Input:
843
Sample Output:
348
*/

package main
import "fmt"
func main() {
    var n, m, cifra int
    fmt.Scan(&n)
    
    for i:=1; i<=3; i++{
        switch i {
            case 1: {
                cifra = n % 10
                m+=cifra*100
            }
            case 2: {
                cifra = (n % 100)/10
                m+=cifra*10
            }
            case 3: {
                cifra = (n % 1000)/100
                m+=cifra}    
        }
        
    } 
    fmt.Println(m)
 }  
