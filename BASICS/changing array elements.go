/* 
In the first step, 10 positive integers are fed to standard input, which must be written in the order they were entered into an array of 10 elements.
The type of numbers included in the array must correspond to the smallest possible unsigned integer.
At the second stage, 3 more pairs of numbers are fed to the standard input - the indices of the elements of this array that need to be swapped
if such a pair of numbers is 3 and 7, then the element with index 3 in the array must be swapped with the element whose index is 7).
The elements of the resulting array should be output separated by a space on the standard output. 
*/


/*
На первом этапе на стандартный ввод подается 10 целых положительных чисел, которые должны быть записаны в порядке ввода в массив из 10 элементов.
Тип чисел, входящих в массив, должен соответствовать минимально возможному целому беззнаковому числу. 
На втором этапе на стандартный ввод подаются еще 3 пары чисел - индексы элементов этого массива, которые требуется поменять местами 
если такая пара чисел 3 и 7, значит в массиве элемент с 3 индексом нужно поменять местами с элементом, индекс которого 7).
Элементы полученного массива должны быть выведены через пробел на стандартный вывод.
 */


package main
import "fmt";
func main() {
	var workArray [10] uint8
	var changeIndx [6] uint8
	var a uint8
	for i:=0; i<10; {
		fmt.Scan(&a)
		workArray [i]=a
		i++
	}
	for i:=0; i<6; {
		fmt.Scan(&a)
		changeIndx [i]=a
		i++
	}

	for i:=0; i<3; {
		var j = i*2
		var ind1 = changeIndx[j]
		var ind2 = changeIndx[j+1]
		workArray[ind1], workArray[ind2] = workArray[ind2], workArray[ind1];
		i++
	}

	for _, elem := range workArray {
		fmt.Print(elem)
		fmt.Print(" ")
	}
}
