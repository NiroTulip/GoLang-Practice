/*
На стандартный ввод вы получаете 2 вещественных числа в формате CSV, где в качестве разделителя используется символ «;». В качестве результата требуется вывести частное от деления первого числа на второе с точностью до четырех знаков.
Sample Input:
1 149,6088607594936;1 179,0666666666666
Sample Output:
0.9750
 */

/*
On standard input, you get 2 real numbers in CSV format, where the character ";" is used as a separator. As a result, it is required to display the quotient of dividing the first number by the second with an accuracy of up to four digits.
Sample input:
1 149.6088607594936;1 179.0666666666666
Sample output:
0.9750
*/


package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ScanStr() string {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println("Ошибка ввода: ", err)
	}
	str = strings.TrimSuffix(str, "\n")
	return str
}

func parseNumber(str string) float64{
	str = strings.Replace(str," ", "", -1)
	str = strings.Replace(str,",", ".", -1)
	result, err := strconv.ParseFloat(str, 64)
	if err != nil{
		panic(err)
	}
	return result
}

func main() {
	str := ScanStr()
	numbers := strings.Split(str, ";")
	s_num1 := numbers[0]
	s_num2 := numbers[1]
	num1 := parseNumber(s_num1)
	num2 := parseNumber(s_num2)
	fmt.Printf("%.4f", num1/num2)
}