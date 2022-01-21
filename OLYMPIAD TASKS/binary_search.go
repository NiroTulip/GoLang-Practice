/* 
Write a program that uses the binary search algorithm to perform the following task:
In an ordered array of unique elements (string type words), finds the index of a specific element from this array.
Input: (["азбука", "дети", "ягода"], "ягода")
Output: 2
*/


/*
Напишите программу, которая с помощью алгоритма бинарного поиска выполняет задачу:
В упорядоченном массиве из уникальных элементов (слова строкового типа) находит индекс конкретного элемента из этого массива.

Ввод: (["азбука", "дети", "ягода"], "ягода")
Вывод: 2
*/

package main
import "fmt"
var iteration = 0

func getElemIdxBinary(source [] string, target string, start, end int) int{
	iteration += 1
	if (target < source[0] || target > source[len(source)-1]){return -1}
	hit := start + (end-start)/2

    fmt.Printf("Текущая итерация: %d, текущий индекс: %d, текущее слово: %q \n", iteration, hit, source[hit])

	if (source[hit] == target) {return hit} else{
		if (end <= start) {return -1}
		if (source[hit] < target){
			start = hit + 1
			result := getElemIdxBinary(source, target, start, end)
			return result
		}else{
			end = hit - 1
			result := getElemIdxBinary(source, target, start, end)
			return result
		}
	}
}

func main() {

example := [] string {"авангард", "аншлаг", "билет", "бокал", "вера", "виноград", "глобус", "город", "деревня", "дружба",
"еж", "енот", "жаба", "жизнь", "заря", "звон", "ива", "иголка", "йемен", "йод", "корова", "крыша", "лимон", "лист", "мама",
"мишка", "надежда", "ностальгия", "одежда", "орех", "планета", "пьедестал", "равнина", "ресурс", "сажа", "степень", "тарелка",
"тетрадь", "улитка", "уравнение", "фара", "финал", "хакатон", "харизма", "цель", "цунгцванг", "чаша", "черный", "шалаш",
"шепот", "щавель","щи", "энергия", "эра", "юг", "юность", "яблоко", "ядро"}

target := "юность"

result := getElemIdxBinary(example, target, 0, len(example)-1)
fmt.Println()
if (result == -1) {fmt.Println("Слово не найдено")} else {
	fmt.Printf("Cлово %q найдено. Его индекс: %d \n", target, result)
	fmt.Println("Всего итераций:", iteration)
}
iteration=0
}

