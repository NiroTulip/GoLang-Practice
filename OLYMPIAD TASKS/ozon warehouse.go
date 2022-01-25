/* 
OLYMPIAD TASK FROM OZON.RU FROM ROUTE 256 CONTEST

Once, several consignments of goods were brought to the warehouse at once and they were unloaded into N different piles, the size of each heap is ai. At the same time, it is known that in the first heap the goods have numbers from 1 to a1, in the second from a1 + 1 to a1 + a2, and so on.

But from these goods it is necessary to collect orders. Therefore, it is necessary as quickly as possible to determine in which heap it lies by the number of the goods.

Input data format
The first line contains the number 0 < N < 10(5) - this is the number of heaps.
On the next N integers a1, a2, …, an. Moreover, 1 < ai < 10(5), and in total there are no more than 10(6) goods.
Further on a new line 1 < M < 10(5) is the number of goods to be found.
The last line contains M numbers of goods to be found.

Output format
It is required to output M lines, each of which contains the number of the heap.
Examples

Input data:

5
8 52 50 35 11
5
78 45 79 99 66

Output:

3
2
3
3
3
*/


/*
ОЛИМПИАДНАЯ ЗАДАЧА ОТ OZON.RU С КОНТЕСТА ROUTE 256

Однажды на склад привезли сразу несколько партий товаров и выгрузили их в N разных куч, размер каждый кучи ai. При этом известно, что в первой куче товары имеют номера от 1 до a1, во второй от a1 + 1 до a1 + a2 и так далее.

Но из этих товаров надо собрать заказы. Поэтому надо как можно быстрее по номеру товара определять в какой куче он лежит.

Формат входных данных
В первой строке записано число 0 < N < 10(5) - это количество кучек.
На следующей N целых чисел a1, a2, …, an. Причём 1 < ai < 10(5), а всего не более чем 10(6) товаров.
Далее на новой строке 1 < M < 10(5) – это количество товаров, которые надо найти.
На последней строке написано M номеров товаров, которые надо найти.

Формат выходных данных
Требуется вывести M строк, каждая из которых содержит номер кучи.
Примеры

Входные данные:

5
8 52 50 35 11
5
78 45 79 99 66

Выходные данные:

3
2
3
3
3
*/


package main

import "fmt"

var N,a,M, resultHeapNumber int

var heaps = [] int {0}
var goods = [] int {}

func findHeapByNumber(target int)int{
	summ := 0
	for i,e := range heaps{
		summ += e
		if (target <= summ){return i}
	}
	return summ
}

func main() {
fmt.Scan(&N)
for i:=0;i<N;i++{
	fmt.Scan(&a)
	heaps = append(heaps, a)
}
fmt.Scan(&M)
for i:=0;i<M;i++{
	fmt.Scan(&a)
	goods = append(goods,a)
}

for _,good :=range goods{
	resultHeapNumber=0;
	fmt.Println(findHeapByNumber(good))
}

}
