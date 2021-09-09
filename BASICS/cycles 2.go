/* 
The sequence consists of natural numbers and ends with the number 0. Determine the number of elements of this sequence that are equal to its largest element.
Input data format
A non-empty sequence of natural numbers is introduced, ending with the number 0 (the number 0 itself is not included in the sequence, but serves as a sign of its end).
Output data format
Print the answer to the problem.
Sample Input:
1
3
3
1
0
Sample Output:
2 
*/


/*
Последовательность состоит из натуральных чисел и завершается числом 0. Определите количество элементов этой последовательности, которые равны ее наибольшему элементу.
Формат входных данных
Вводится непустая последовательность натуральных чисел, оканчивающаяся числом 0 (само число 0 в последовательность не входит, а служит как признак ее окончания).
Формат выходных данных
Выведите ответ на задачу.
Sample Input:
1
3
3
1
0
Sample Output:
2
*/


package main
import "fmt"
func main() {
    var a, max, kol_max int
    
    for fmt.Scan(&a); a != 0; fmt.Scan(&a){
        
        if a==max {
        kol_max++
        } else if a>max {
        kol_max=1;
        max=a
        }
           
    }
    fmt.Println(kol_max)
}
