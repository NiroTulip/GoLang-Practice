/* 
Number of minimums
Find the number of minimum elements in the sequence.
Input data
A natural number N is entered, followed by N integers of the sequence.
Output
Output the number of minimum elements in a sequence.
Sample Input:
3
21
eleven
4
Sample Output:
1 
*/

/*
Количество минимумов
Найдите количество минимальных элементов в последовательности.
Входные данные
Вводится натуральное число N, а затем N целых чисел последовательности.
Выходные данные
Выведите КОЛИЧЕСТВО минимальных элементов последовательности.
Sample Input:
3
21
11
4
Sample Output:
1
*/

package main
import "fmt"
func main() {
    var n,x,min,kol_min int
    var a [] int
    fmt.Scan(&n)
       
    for i:=1; i<=n; i++{
        fmt.Scan(&x)
        a = append(a,x)
        }
    min = a[0]
    for _, elem := range a {

        if elem<min {min=elem
                     kol_min=1} else if elem==min {kol_min++}
    }     
    fmt.Print(kol_min)
 }   
