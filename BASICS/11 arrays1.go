/* 
Five numbers are fed into the input, which are written to an array. You need to find and display the maximum number in this array. 
*/


/*
На ввод подаются пять чисел, которые записываются в массив. Нужно найти и вывести максимальное число в этом массиве.
*/

package main
import "fmt"

func main()  {
	array := [5]int{}
	var a, max int
	for i:=0; i < 5; i++{
		fmt.Scan(&a)
		array[i] = a
	}
    for _, elem := range array{
        if elem > max {max = elem}
    }
    fmt.Println(max)
}
