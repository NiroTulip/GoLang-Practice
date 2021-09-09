/* 
The input is a float64 number. You need to output the converted number according to the rule: the number is squared, then the fractional part is cut off so that 4 decimal places remain. But if the number is equal to or less than zero - output:
"number R does not match", where R is the original number with 2 digits after the decimal point and 2 in width. And if the number is more than 10,000 - display the original number in exponential notation (lowercase exponent sign).

Sample input:
-000012.2123
Sample output:
the number -12.21 is not suitable

Sample input:
1000001
Sample output:
1.000001e + 06

Sample Input:
12.12345678
Sample Output:
146.9782 
*/


/*
На вход подается число типа float64. Вам нужно вывести преобразованное число по правилу: число возводится в квадрат, затем обрезается дробная часть так что остается 4 знака после запятой. Но если число равно или меньше нуля - выводить:
"число R не подходит", где R - исходное число с 2 цифрами после запятой и с 2 по ширине. А если число больше чем 10 000 - выводить исходное число в экспоненциальном представлении (знак экспоненты в нижнем регистре).

Sample input:
-000012.2123
Sample output:
число -12.21 не подходит

Sample input:
1000001
Sample output:
1.000001e+06

Sample Input:
12.12345678
Sample Output:
146.9782
*/

package main
import "fmt"
func main() {

    var n, n2 float64
    fmt.Scan(&n)
    if n<=0{
    fmt.Printf("число %2.2f не подходит", n)
    } else if n>10000{
    fmt.Printf("%e", n)
    } else {
    n2 = n*n
    fmt.Printf("%.4f", n2)
    }
}
