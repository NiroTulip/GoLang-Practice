/* 
Digital root
The digital root of a natural number is a digit obtained as a result of an iterative process of summing digits, at each iteration of which, to calculate the sum of digits, the result obtained at the previous iteration is taken. This process is repeated until one digit is received.
For example, the digital root 65536 is 7, because 6 + 5 + 5 + 3 + 6 = 25 and 2 + 5 = 7.
For this number, determine its digital root.
Input data
One natural number n is entered, not exceeding 10710 ^ 7107.
Output
Print the digital root of the number n.
Sample Input:
3456
Sample Output:
9
*/

/*
Цифровой корень
Цифровой корень натурального числа — это цифра, полученная в результате итеративного процесса суммирования цифр, на каждой итерации которого для подсчета суммы цифр берут результат, полученный на предыдущей итерации. Этот процесс повторяется до тех пор, пока не будет получена одна цифра.
Например цифровой корень 65536 это 7 , потому что 6+5+5+3+6=25 и 2+5=7 . 
По данному числу определите его цифровой корень.
Входные данные
Вводится одно натуральное число n, не превышающее 10710^7107.
Выходные данные
Вывести цифровой корень числа n.
Sample Input:
3456
Sample Output:
9
*/

package main
import "fmt"
func main() {
    var n,d,sum int
    fmt.Scan(&n)
  
for {
    for n>0 {
    d=n%10
    sum+=d
    n=n/10
    }
    
    if sum>10 {
    n=sum
    sum=0
    } else {break}
}
    fmt.Print(sum)
    
}   
