/* 
You are given a sequence of integers. Write a program that counts the number of positive numbers among the elements of a sequence.
Input data
First, the number NNN is given - the number of elements in the sequence (1≤N≤1001 \ leq N \ leq1001≤N≤100). Then NNN numbers are written separated by a space - the elements of the sequence. The sequence consists of integers.
Output
You must print a single number - the number of positive elements in the sequence.
Sample Input:
5
1 2 3 -1 -4
Sample Output:
3 
*/

/*
Дана последовательность, состоящая из целых чисел. Напишите программу, которая подсчитывает количество положительных чисел среди элементов последовательности.
Входные данные
Сначала задано число NNN — количество элементов в последовательности (1≤N≤1001\leq N\leq1001≤N≤100). Далее через пробел записаны NNN чисел — элементы последовательности. Последовательность состоит из целых чисел.
Выходные данные
Необходимо вывести единственное число - количество положительных элементов в последовательности.
Sample Input:
5
1 2 3 -1 -4
Sample Output:
3
*/

package main
import "fmt"
func main() {
    var n,x,j int
    var a [] int
    fmt.Scan(&n)
       
    for i:=1; i<=n; i++{
        fmt.Scan(&x)
        a = append(a,x)
        }
    
    for _, elem := range a {
        if elem > 0 {j++}
    }     
    fmt.Print(j)
 }   
