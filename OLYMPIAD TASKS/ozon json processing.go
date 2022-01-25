/* 
OLYMPIAD TASK FROM OZON.RU FROM ROUTE 256 CONTEST

Counting number of unique elements in JSON
There is a JSON consisting of an array of objects, in each object of which there is a data field. The data field contains an object that is guaranteed to have the key` string key. Array length 0 ≤ x ≤ 105.

It is required to count the number of unique values ​​among all keys in all data.

Input data format
The input is the number 0 < N ≤ 105, followed by N lines containing the JSON file. It is guaranteed that the length of each string is l < 105 characters.

Output format
You need to send a single number to the output - the number of unique keys in all data.

Input data:
twenty
[
  {
    data: {
      "key": "one"
    }
  },
  {
    data: {
      "key": "two"
    }
  },
  {
    data: {
      "key": "two"
    },
    "meta": {
      "key": "hello"
    }
  }
]

Output:
2
*/


/*
ОЛИМПИАДНАЯ ЗАДАЧА ОТ OZON.RU С КОНТЕСТА ROUTE 256

Подсчет количества уникальных элементов в JSON
Есть JSON, состоящий из массива объектов, в каждом объекте которого есть поле data. В поле data лежит объект, в котором гарантируется наличие строкового ключа key`. Длина массива 0 ≤ x ≤ 105.

Требуется посчитать количество уникальных значений среди всех key во всех data.

Формат входных данных
На вход подаётся число 0 < N ≤ 105, далее идут N строк, содержащие в себе JSON файл. Гарантируется, что длина каждой строки l < 105 символов.

Формат выходных данных
На выход требуется отправить единственное число – количество уникальных ключей во всех data.

Входные данные:
20
[
  {
    "data": {
      "key": "one"
    }
  },
  {
    "data": {
      "key": "two"
    }
  },
  {
    "data": {
      "key": "two"
    },
    "meta": {
      "key": "hello"
    }
  }
]

Выходные данные:
2
*/


package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var N int
var str string
var input_strings = [] string {}
var uniqueKeysMap = map [string] bool {}

//Вспомогательная функция извлекает значение ключа key из полученной строки и возвращает его
func getKeyValFromString(str string) string{
	list := strings.Split(str, "\"")
	for idx,elem:= range list{
		if (elem == "key") {
			if (idx+2<len(list)){
				return list[idx+2]
			}
		}
	}
	return ""
}

//Функция получает на вход строку, содержащую ключ key, определяеь его значение с помощью вызова функции
//getKeyValFromString() и проверяет, есть ли он в mapе уникальных ключей. Если нет, то добавляет его туда.
func processKeyString(str string){
keyVal:=getKeyValFromString(str)
if (keyVal != ""){
	if _, found := uniqueKeysMap[keyVal]; found{} else{
		uniqueKeysMap[keyVal]=true
	}
}
}

//Функция получает на вход номер строки - начала блока data и прочесывает следующие строки, пока не встретит ключ key
//после чего передает строку с ключом key функции processKeyString() для обработки
func processDataBlock(start int) {
 for i:=start+1; i<len(input_strings); i++{
 	 str = input_strings[i]
	 if (strings.Contains(str, "key")) {processKeyString(str); return}
 }
}

func main() {
	fmt.Scan(&N)
	inputReader := bufio.NewReader(os.Stdin)
	var input string
	for i:=0;i<N;i++ {
		input, _ = inputReader.ReadString('\n')
		input_strings = append(input_strings,input)
	}
//Цикл прочесывает все строки, пока не найдет сроку начала блока data, после чего вызывает фнкцию обработки блоков data
//processDataBlock(), передавая ей номер строки начала блока
	for idx,str := range input_strings{
		if (strings.Contains(str, "\"data\"")) {processDataBlock(idx)}
	}
	//После обработки всех строк map uniqueKeysMap будет содержать все уникальные ключи key из блоков data.
	//Ее длина и есть искомое количество.
	fmt.Println(len(uniqueKeysMap))
}


