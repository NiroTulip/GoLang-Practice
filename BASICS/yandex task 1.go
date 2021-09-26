/* 
Input: an array of strings;
Output: an array of strings from input that can be typed by keys from the same keyboard row 

Input: ["фара", "лампочка", "свет"]
Output: ["фара"]

Input: ["вода", "лед", "пламя"]
Output: ["вода"]
*/


/*
Ввод: массив из строк;
Вывод: массив строк из ввода, которые могут быть напечатаны клавишами из одного ряда клавиатуры

Ввод: ["фара", "лампочка", "свет"]
Вывод: ["фара"]

Ввод: ["вода", "лед", "пламя"]
Вывод: ["вода"]
*/


package main

import (
	"fmt"
	"strings"
)

//Массив хранит шаблоны трех рядов клавиатуры
var layouts [3] string = [3] string {"йцукенгшщзхъ","фывапролджэ", "ячсмитьбю"}

//Функция возвращает номер индекса шаблона, в котором найден символ symb (аргумент функции)
func indexOfLayot (symb string) int{
	for idx, elem := range layouts{
		if (strings.Contains(elem, symb)) {return idx}
	}
	return -1;
}

//Функция проверяет, может ли слово word (аргумент функции) быть напечатано символами из одного шаблона
func isWordPrintableByOneLayot (word string) bool {
	prevLayot := -1;
	for idx, elem := range word{
		symb := string(elem)
		if (idx == 0) {prevLayot = indexOfLayot(symb)} else {
			currentLayot := indexOfLayot(symb)
			if (prevLayot != -1 && currentLayot != prevLayot) {return false} else{
				prevLayot = currentLayot
			}
		}
	}
	if (prevLayot == -1) {return false}
	return true
}

//Функция принимает входной срез с данными для проверки и возвращает результат - срез,
//содержащий только те слова, которые могут быть напечатаны символами из одного шаблона
func giveResultSlice (inputSlice [] string) [] string{
var outputSlice [] string
	for _,elem :=range inputSlice{
	if (isWordPrintableByOneLayot(elem)) {outputSlice = append(outputSlice, elem)}
}
return outputSlice
}

func main() {
//Входные данные
example1 := [] string {"фара", "лампочка", "свет"}
example2 := [] string {"вода", "лед", "пламя"}

fmt.Println(giveResultSlice(example1))
fmt.Println(giveResultSlice(example2))
}
