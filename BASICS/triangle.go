/* 
Input data
Three natural numbers a, b, c are given. Determine if there is a triangle with such sides.
Output
If the triangle exists, print the line "Exists", otherwise print the line "Doesn't exist". Print the string without quotes.
Sample Input:
4 5 6
Sample Output:
Exists 
*/

/*
Входные данные
Даны три натуральных числа a, b, c. Определите, существует ли треугольник с такими сторонами.
Выходные данные
Если треугольник существует, выведите строку "Существует", иначе выведите строку "Не существует". Строку выводите без кавычек.
Sample Input:
4 5 6
Sample Output:
Существует
*/

package main
import "fmt"
func main() {
    var a,b,c int
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&c)
    if a+b>c && a+c>b && b+c>a {
        fmt.Print("Существует")
    } else {fmt.Print("Не существует")}
 }   
