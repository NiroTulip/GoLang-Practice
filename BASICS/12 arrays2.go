/* 
An array of integers is given. The numbering of elements starts from 0. Write a program that displays the elements of an array whose indices are even (0, 2, 4 ...).
Input data
First, the number NNN is given - the number of elements in the array (1≤N≤1001 \ leq N \ leq 1001≤N≤100). Then NNN numbers are written separated by a space - the elements of the array. The array consists of integers.
Output
It is necessary to display all the even-numbered elements of the array.
Sample Input:
6
4 5 3 4 2 3
Sample Output:
4 3 2 
*/


/*
Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0. Напишите программу, которая выведет элементы массива, индексы которых четны (0, 2, 4...).
Входные данные
Сначала задано число NNN — количество элементов в массиве (1≤N≤1001 \leq N \leq 1001≤N≤100). Далее через пробел записаны NNN чисел — элементы массива. Массив состоит из целых чисел.
Выходные данные
Необходимо вывести все элементы массива с чётными номерами.
Sample Input:
6
4 5 3 4 2 3
Sample Output:
4 3 2
*/

package main
import "fmt"
func main() {
    var n,x int
    var a [] int
    fmt.Scan(&n)
       
    for i:=1; i<=n; i++{
        fmt.Scan(&x)
        a = append(a,x)
        }
    
    for idx, elem := range a {
        if idx % 2 == 0 {fmt.Print(elem, " ")}
    }     
    
 }
