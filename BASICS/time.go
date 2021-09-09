/* 
The K-th second of the day is passing. Determine how many whole hours h and whole minutes m have passed since the beginning of the day. For example, if
k = 13257 = 3 * 3600 + 40 * 60 + 57,
then h = 3 and m = 40.
Input data
The input to the program is an integer k (0 <k <86399).
Output
Display the phrase:
It is ... hours ... minutes.
Instead of ellipsis, the program should print the values of h and m, separating them from words with exactly one space.
Sample Input:
13257
Sample Output:
It is 3 hours 40 minutes. 
*/

/*
Идёт k-я секунда суток. Определите, сколько целых часов h и целых минут m прошло с начала суток. Например, если
k=13257=3*3600+40*60+57,
то h=3 и m=40.
Входные данные
На вход программе подается целое число k (0 < k < 86399).
Выходные данные
Выведите на экран фразу:
It is ... hours ... minutes.
Вместо многоточий программа должна выводить значения h и m, отделяя их от слов ровно одним пробелом.
Sample Input:
13257
Sample Output:
It is 3 hours 40 minutes.
*/

package main
import "fmt"
func main() {
    var n,h,m int
    fmt.Scan(&n)
    h = n/3600
    m = (n%3600)/60
    fmt.Print("It is ",h," hours ",m," minutes.")
 } 
