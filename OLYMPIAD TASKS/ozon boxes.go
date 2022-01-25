/* 
OLYMPIAD TASK FROM OZON.RU FROM ROUTE 256 CONTEST

There are n boxes of apples in the warehouse (an even number).
The storekeeper wants to form n/2 pallets. Each pallet must consist of exactly two boxes. Two boxes can only be stable on a pallet if their weight is the same.
You can put apples in boxes, each of which increases the weight of the box by one.
The storekeeper wants to know what is the minimum number of apples that must be put into boxes to form exactly n/2 pallets (that is, each pair of boxes must lie on a pallet). Your task is to find this number.

Input data format
The first line of the input contains one integer n (2 <= n <= 100) — the number of boxes. It is guaranteed that n is always an even number.
The second line of the input contains n integers a1, a2, …, an (1 <= a[i] <= 100), where a[i] is equal to the weight of the i-th box.

Output format
Print a single integer — the minimum number of apples you need to put into the boxes to form exactly n/2 pallets.

Examples
Input data:
6
5 10 2 3 14 5
Output:
5 
*/

/*
ОЛИМПИАДНАЯ ЗАДАЧА ОТ OZON.RU С КОНТЕСТА ROUTE 256

На складе лежат n коробок с яблоками (четное число).
Кладовщик хочет сформировать n/2 паллет. Каждая паллета должна состоять ровно из двух коробок. Две коробки могут устойчиво лежать на паллете только тогда, когда их вес одинаков.
В коробки можно докладывать яблоки, каждое из которых увеличивает вес коробки на единицу.
Кладовщик хочет знать, какое минимальное количество яблок необходимо доложить в коробки, чтобы сформировать ровно n/2 паллет (то есть каждая пара коробок должна лежать на паллете). Ваша задача — найти это количество.

Формат входных данных
Первая строка входных данных содержит одно целое число n (2 <= n <= 100) — количество коробок. Гарантируется, что n всегда является четным числом.
Вторая строка входных данных содержит n целых чисел a1, a2, …, an (1 <= a[i] <= 100), где a[i] равно весу i-й коробки.

Формат выходных данных
Выведите одно целое число — минимальное количество яблок, которое необходимо доложить в коробки, чтобы сформировать ровно n/2 паллет.

Примеры
Входные данные:
6 
5 10 2 3 14 5
Выходные данные:
5
*/

package main

import (
	"fmt"
	"sort"
)

var n,w byte
var weight [] byte

//Функция находит и возвращает позицию элемента, который равен следующему элементу. Равные элементы будут идти один за другим, т.к. срез уже был отсортирован.
func findTwoEqualElemsInSlice(input [] byte) int{
	for i:=0; i<len(input)-1;i++{
		if (input[i] == input[i+1]) {return i}
	}
return -1;
}

//Функция удаляет два равных элемента из среза.
func deleteTwoElemFromSlice(pos1 int, input [] byte) [] byte{
	var result [] byte
	if (pos1 > 0) {result = append(result, input[:pos1]...)}
	if (pos1 < len(input)-2) {result = append(result, input[pos1+2:]...)}
	return result
}

//Функция подсчитывает, сколько яблок необходимо добавить в коробки, попарно проверяя соседние элементы в отсортированном срезе
func findCountAdditionalApples(input [] byte) int {
var summ int

	for i:=0; i<len(input)-1;i++{
		if (i % 2 != 0) {continue} else{
			summ += int(input[i+1] - input[i])
		}
	}
return summ
}

func main() {

//Цикл считывает количество коробок n
 for {
 	fmt.Scan(&n)
 	if (n % 2 == 0) {break}
 }

//Цикл считывает количество яблок в каждой коробке
 var i byte
 i=0
 for ; i<n; i++{
	 fmt.Scan(&w)
	 weight = append(weight,w)
 }

 //Сортировка среза, содержащего количество яблок в коробках
sort.Slice(weight, func(i, j int) bool {
	return weight[i] < weight[j]
})

 //Цикл находит и удаляет из среза пары равных элементов, т.к. для них не надо докладывать яблоки.
for {
	firstEqualPos := findTwoEqualElemsInSlice(weight)
	if (firstEqualPos != -1) {weight = deleteTwoElemFromSlice(firstEqualPos, weight)} else {break}
	if (len(weight) == 0) {break}
}

 //После удаления равных элементов остается отсортированный срез, в котором надо просто подсчитать количество недостающих яблок
fmt.Println(findCountAdditionalApples(weight))
}

