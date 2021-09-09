/* 
It is required to determine whether a given year is a leap year, recall:
A year is a leap year if it meets at least one of the following conditions:
- multiple of 400;
- a multiple of 4, but not a multiple of 100.
Input data
A single number is entered - the number of the year (integer, positive, does not exceed 10000).
Output
You want to output the word YES if the year is a leap year and NO otherwise.
Sample Input:
2000
Sample Output:
YES 
*/


/*
Требуется определить, является ли данный год високосным, напомним:
Год является високосным если он соответствует хотя бы одному из нижеперечисленных условий:
- кратен 400;
- кратен 4, но не кратен 100.
Входные данные
Вводится единственное число - номер года (целое, положительное, не превышает 10000).
Выходные данные
Требуется вывести слово YES, если год является високосным и NO - в противном случае.
Sample Input:
2000
Sample Output:
YES
*/


package main
import "fmt"
func main() {
var n int 
fmt.Scan(&n)
    if n % 400 == 0 || (n % 4 == 0 && n % 100 != 0) {
        fmt.Println("YES")
    } else {fmt.Println("NO")}
}
