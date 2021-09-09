/* 
Two numbers are given. Determine the numbers included in the recording of both the first and second numbers.
Input data
The program receives two numbers as input. It is guaranteed that numbers in numbers do not repeat. Numbers ranging from 0 to 10000.
Output
The program should print the numbers that are in both numbers, separated by a space. The numbers are displayed in the order they appear in the first number.
Sample Input:
564 8954
Sample Output:
5 4 
*/


/*
Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.
Входные данные
Программа получает на вход два числа. Гарантируется, что цифры в числах не повторяются. Числа в пределах от 0 до 10000.
Выходные данные
Программа должна вывести цифры, которые имеются в обоих числах, через пробел. Цифры выводятся в порядке их нахождения в первом числе.
Sample Input:
564 8954
Sample Output:
5 4
*/

package main
import "fmt"
func main() {
    var a,b,c1,c2,i,va,vb int
    fmt.Scan(&a, &b)
     
    va=a
    
for va>0{
    
    c1=va
    i=1
    for c1>10 {
        c1=c1/10
        i=i*10       
    }
    
    vb=b
    for vb>0{
    c2=vb%10
        if c2==c1 {
            fmt.Print(c2, " ")
            break
        }
    vb=vb/10
    }
    
    va=va-c1*i
} 
}
