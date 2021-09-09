/* 
For the given number n, finish the phrase "Grazing in the meadow ..." with one of the possible continuation: "n cows", "n cow", "n cows", correctly declining the word "cow".
Input data
Given number n (0 <n <100).
Output
The program should display the entered number n and one of the words (in Latin): korov, korova or korovy, for example, 1 korova, 2 korovy, 5 korov. There must be exactly one space between the number and the word.
Sample Input:
ten
Sample Output:
10 korov 
*/

/*
По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".
Входные данные
Дано число n (0<n<100).
Выходные данные
Программа должна вывести введенное число n и одно из слов (на латинице): korov, korova или korovy, например, 1 korova, 2 korovy, 5 korov. Между числом и словом должен стоять ровно один пробел.
Sample Input:
10
Sample Output:
10 korov
*/


package main
import "fmt"
func main() {
    var n,c int
    var s string
    fmt.Scan(&n)
    c=n%10
    switch {
        case c==1: s="korova"
        case 2<=c && c<=4: s="korovy"
        case c==0 || c>=5: s="korov"
    }
    if n>=11 && n<=14 {
        s="korov"
    }
    fmt.Println(n, s)
} 
