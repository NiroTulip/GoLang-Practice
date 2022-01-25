/* 
OLYMPIAD TASK FROM OZON.RU FROM ROUTE 256 CONTEST

Comparison of Roman numbers

Conditions

The Roman numeral system contains the symbols: I, V, X, L, C, D and M.

The following digits of the Roman numeral system correspond to the characters in the decimal system: I = 1; V = 5; x=10; L=50; C=100; D=500; M = 1000.

Numbers in the Roman system are usually written from largest to smallest, from left to right, for example, seven can be written as V + II = VII, and twenty-seven as XX + V + II = XXVII.

However, the number 4 is not written as IIII, but as IV. Those. To obtain some numbers in the Roman numeral system, the principle of subtraction is used.

The following rules specify the conditions under which the subtraction principle should be used:

    1. I can precede V (5) and X (10) to get 4th and 9th, respectively;
    2. X can precede L (50) and C (100) to get 40-a and 90-a respectively;
    3. C may precede D (500) and M (1000) to get 400-a and 900-a respectively;

But not everything is so simple. According to classical Roman rules, the maximum number that can be represented in this number system is 3999. However, there are ways to get around this limitation. We will use the vinculum modification. In the basic version, this allows you to divide the number into two blocks of digits, the first of which is multiplied by 1000, i.e. IV|DCXXVII = 4 627, while the second block is not more than a thousand, so that there are no collisions when writing. We will go a little further and introduce the following rules:

    1. There can be as many blocks as you like, which are separated by the symbol |
    2. Each of the blocks, except for the first one, contains a number less than 1000
    3. The first block can be any Roman number.
    4. It is assumed that we write all numbers up to and including 3999 as before, so that there are no collisions.
    5. An empty bit is replaced by the symbol N (nihil, nothing). N is used only to write an empty bit. XXN is not a valid entry.

Example entries:

    1. X|N|X = 10 000 010
    2. X|N = 10,000
    3. DCXXVII|X|N|X = 627 010 000 010

Input data format

The input is a pair of Roman numbers a and b, written in Roman form with the above modification.

Numbers are given in one line, separated by a space. Total number length <= 10^5
Output format

You need to write a comparison program that returns -1 if a is less than b, 0 if they are equal, 1 if a is greater than b.
Examples
Input data:

MMLXII MMLX

Output:

1

Input data:

XX|XX XX|XXI

Output:

-1
*/


/*
ОЛИМПИАДНАЯ ЗАДАЧА ОТ OZON.RU С КОНТЕСТА ROUTE 256

Сравнение римских чисел

Условия

Римская система счисления содержит символы: I, V, X, L, C, D и M.

Следующие цифры римской системы счисления соответствуют символам в десятичной системе: I = 1; V = 5; X = 10; L = 50; C = 100; D = 500; M = 1000.

Числа в римской системе обычно записываются от большего к меньшему, слева направо, например, семь можно записать как V + II = VII, а двадцать семь как XX + V + II = XXVII.

Однако, число 4 записывается не как IIII, а как IV. Т.е. Для получения некоторых цифр в римской системе счисления используется принцип вычитания.

Следующие правила задают условия, при которых следует использовать принцип вычитания:

    1. I может предшествовать V (5) и X (10) для получения 4-х и 9-и соответственно;
    2. X может предшествовать L (50) и С (100) для получения 40-а и 90-а соответственно;
    3. С может предшествовать D (500) и M (1000) для получения 400-а и 900-а соответственно;

Но не всё так просто. Согласно классическим римским правилам, максимальное число, которое может быть представлено в этой системе счисления, это 3999. Однако, есть способы, которые позволяют обойти это ограничение. Мы будем использовать модификацию vinculum. В базовой версии это позволяет разделить число на два блока цифр, первый из которых умножается на 1000, т.е. IV|DCXXVII = 4 627, при этом второй блок не больше тысячи, чтобы не было коллизий при записи. Мы же пойдём чуть дальше и введём следующие правила:

    1. Может быть сколько угодно много блоков, которые разделены символом |
    2. Каждый из блоков, кроме первого, содержит в себе число меньше 1000
    3. Первый блок может быть каким угодно римским числом.
    4. Предполагается, что все числа до 3999 включительно мы пишем как раньше, чтобы не было коллизий.
    5. Пустой разряд заменяется символом N (nihil, ничто). N используется только для записи пустого разряда. XXN не валидная запись.

Примеры записей:

    1. X|N|X = 10 000 010
    2. X|N = 10 000
    3. DCXXVII|X|N|X = 627 010 000 010

Формат входных данных

На вход передается пара римских чисел a и b, записанных в римской форме с указанной выше модификацией. 

Числа подаются в одной строке, через пробел. Общая длина чисел <= 10^5
Формат выходных данных

Необходимо написать программу сравнения, которая возвращает -1, если a меньше b, 0, если равны, 1, если a больше b.
Примеры
Входные данные:

MMLXII MMLX

Выходные данные:

1

Входные данные:

XX|XX XX|XXI

Выходные данные:

-1
*/

package main

import (
	"fmt"
	"strings"
)

var a, b string
var LettersValues = map[string] uint64{"N":0, "I":1, "V":5, "X":10, "L":50, "C":100, "D":500, "M":1000, "Q":4, "W":9, "E":40, "R":90, "T":400, "Y":900}
var Substitution = map[string] string {"IV":"Q", "IX":"W", "XL":"E", "XC":"R", "CD":"T", "CM":"Y"}

//Функция осуществляет замену биграмм (IV, IX, XL, XC, CD, CM) из исходной строки их условными буквенными подменами,
//т.к. дальнейшая обработка будет проводиться по буквам
func getSubstitudedString(str string) string{
	for key, val := range Substitution{
		if ((strings.Index(str, key) != -1)){
			str = strings.Replace(str, key, val, -1)
		}
	}
	return str
}

//Функция преобразует исходное число, выраженное строкой, в его численный эквивалент числового типа и возвращает его
func convertOneStrToNum(str string) uint64{
	var summ uint64
	summ = 0
	if (str == "N") {return (0)}
	str = getSubstitudedString(str)
	for _,char := range str{
		summ += LettersValues[string(char)]
	}
	return summ
}

//Функция разделяет исходное число, выраженное строкой, на блоки, разделенные разделителем | и помещает
//отдельные блоки в срез, который возвращает
func getSliceOfBlocks(str string) [] string{
	return strings.Split(str, "|")
}

//Основная функция, которая делает основную работу - сравнивает 2 исходных числа, записанные в строках str1 и str2
func compareStrings(str1, str2 string) int8{
	block1 := getSliceOfBlocks(str1)
	block2 := getSliceOfBlocks(str2)
	//Следующие 2 if определяют победителя досрочно путем сравнения длин конкурентов. У кого больше длина (т.е. больше
	//количество блоков, разделенных /), у того и числовое значение больше, дальнейшие расчеты и сравнения не требуются
	if (len(block1) < len(block2)) {return -1}else{
		if (len(block1) > len(block2)) {return 1}else{
			//Если длины конкурентов равны, то начинается раунд сравнения.
			//Попарно сравниваются блоки одного ранга слева направо, начиная с самого первого (левого).
			//Как только один из блоков больше другого, это и определяет победителя, дальнейшее сравнение блоков низших
			//рангов (которые идут правее) не требуется,т.к. они на порядок меньше текущих блоков.
			for i,str := range block1{
				if (convertOneStrToNum(str) < convertOneStrToNum(block2[i])) {return -1} else {
					if (convertOneStrToNum(str) > convertOneStrToNum(block2[i])) {return 1} else {continue}
				}

			}
		}
	}
//Если после всех проверок победитель не выявлен, значит числа равны, возвращается 0
return 0
}

func main() {
fmt.Scan(&a, &b)
fmt.Println(compareStrings(a,b))
}


