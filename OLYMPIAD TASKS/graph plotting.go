/* 
OLYMPIAD TASK FROM OZON.RU FROM ROUTE 256 CONTEST

Plotting

The database stores the history of chart values ​​by day. Each value is either an integer x ≥ 0 or the special value -1.

It is required to display the values ​​of the graph points according to the data from the database. The following rules apply:

    1. if there is no value for day X in the data or it is equal to -1, then the value of the nearest previous day for which there is a value is considered current;
    2. if there are no values ​​before day X, then on this day we give a special value -1, which indicates that there is no data;
    3. for any day in the future we give a special value -1;

Input data format

The input on the first line is the following numbers:

    1. number of the first day of the chart (inclusive) - a ≥ 0;
    2. number of the last day (inclusive) b ≥ 0, while b > a;
    3. today's number c;
    4. number of records in the database 0 < d < 105.

This is followed by data in the format of pairs (day number, value), one pair in each row, sorted in reverse order (i.e., the last day comes first, then the rest). The number of pairs may not be equivalent to the number of days that will be in the chart.
Output format

As a result, all days from the first to the last inclusive must be present in direct order.
Examples
Input data:

11 17 15 2
14 11
12 10

Output:

11-1
12 10
13 10
14 11
15 11
16-1
17-1
*/


/*
ОЛИМПИАДНАЯ ЗАДАЧА ОТ OZON.RU С КОНТЕСТА ROUTE 256

Построение графика

В базе данных хранится история значений графика по дням. Каждое значение – это целое число x ≥ 0, либо специальное значение -1.

Требуется вывести значения точек графика по данным из базы. При этом действуют следующие правила:

    1. если в данных отсутствует значение для дня X или оно равно -1, то текущим считается значение ближайшего предыдущего дня, для которого есть значение;
    2. если до дня X нет никаких значений, то в этот день отдаем специальное значение -1, которое говорит о том, что данных нет;
    3. для любого дня в будущем отдаем специальное значение -1;

Формат входных данных

На вход в первой строке подаются следующие числа:

    1. номер первого дня графика (включительно) - a ≥ 0;
    2. номер последнего дня (включительно) b ≥ 0, при этом b > a;
    3. номер сегодняшнего дня c;
    4. количество записей в БД 0 < d < 105.

Далее следуют данные в формате пар (номер дня, значение), по одной паре в каждой строке, отсортированные в обратном порядке (т.е. сначала идет последний день, потом остальные). Количество пар может быть не эквивалентно количеству дней, которые будут в графике.
Формат выходных данных

В результате должны присутствовать все дни от первого до последнего включительно в прямом порядке.
Примеры
Входные данные:

11 17 15 2
14 11
12 10

Выходные данные:

11 -1
12 10
13 10
14 11
15 11
16 -1
17 -1 
*/

package main

import "fmt"

var days = [] int {}
var data = map[int] int {}
var start, end, today, entries, day, val, currNoDataVal int

func main() {
	fmt.Scan(&start, &end, &today, &entries)
	for i:=0; i<entries; i++{
		fmt.Scan(&day, &val)
		if (val != -1) {data[day]=val}
	}
	for i:=start; i<=end; i++{
		days=append(days,i)
	}

currNoDataVal = -1
for _, elem := range days{
	if val, found := data[elem]; found {
		currNoDataVal = val
	} else{
		if (elem>today){data[elem]=-1} else{
			data[elem] = currNoDataVal
		}
	}
}
for _, elem := range days{
	fmt.Println(elem, data[elem])
}
}
