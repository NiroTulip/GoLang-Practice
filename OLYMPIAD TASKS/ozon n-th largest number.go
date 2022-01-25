/* 
OLYMPIAD TASK FROM OZON.RU FROM ROUTE 256 CONTEST

Finding the N-th largest number in an array
Conditions
There is an array of integers, you need to find the largest N-th number in it.

Input data format
The input is an integer 0 < M ≤ 105, an integer 0 < N ≤ 105, separated by a space on the first line.
The second line contains an array of M integers separated by a space.

Output format
It is required to print one number - the N-th largest number or -1 if there is no such number.

Examples
Input data:
6 3
1 2 3 -9 3 5

Output:
2

Explanation for example:
    1. the first largest number is 5, it occurs once;
    2. the second largest number is 3, it occurs twice;
    3. The third largest number is 2, it occurs once and is the desired one.
*/


/*
ОЛИМПИАДНАЯ ЗАДАЧА ОТ OZON.RU С КОНТЕСТА ROUTE 256

Поиск N-го самого большого числа в массиве
Условия
Есть массив целых чисел, требуется найти в нём самое большое N-ое число.

Формат входных данных
На вход подаётся целое число 0 < M ≤ 105, целое число 0 < N ≤ 105, разделённые пробелом на первой строке.
На второй строке массив из M целых чисел, разделённых знаком пробел.

Формат выходных данных
Требуется вывести одно число – N-ое самое большое число или -1, если такого числа нет.

Примеры
Входные данные:
6 3
1 2 3 -9 3 5

Выходные данные:
2 

Объяснение к примеру:
    1. первое самое большое число – это 5, оно встречается один раз;
    2. второе самое большое число – это 3, оно встречается два раза;
    3. третье самое большое число – это 2, оно встречается один раз и является искомым.
*/


package main

import (
	"fmt"
	"sort"
)

var M,N,m int
var numbers [] int

//Функция сортирует входной срез по убыванию, убирает из него повторы и возвращает срез с уникальными элементами.
func getUniqueNumbersSlice(input [] int) [] int {
	var result []int

	sort.Slice(input, func(i, j int) bool {
		return input[i] > input[j]
	})

	result = append(result, input[0])
	for i:=1; i<len(input); i++{
        if (input[i] != result[len(result)-1]){
			result = append(result, input[i])
		}
	}
	return result
}

//Функция возвращает N-ый элемент среза, который и будет равен искомому N-ому самому большому числу, т.к. массив
//имеет только уникальные элементы и отсортирован по убыванию.
func getNBigNumber(input [] int, N int) int {
	if (len(input)<N) {return -1} else{
		return input[N-1]
	}
}

func main() {
	fmt.Scan(&M,&N)

	for i:=0; i<M; i++{
		fmt.Scan(&m)
		numbers = append (numbers, m)
	}

	if (len(numbers)>1) {numbers=getUniqueNumbersSlice(numbers)}
	fmt.Println(getNBigNumber(numbers, N))
}
