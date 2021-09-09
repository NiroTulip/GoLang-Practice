/* 
The hour hand has turned d degrees since the beginning of the day. Determine how many whole hours h and whole minutes m are now.
Input data
The input to the program is an integer d (0 <d <360).
Output
Display the phrase:
It is ... hours ... minutes.
Instead of ellipsis, the program should print the values of h and m, separating them from words with exactly one space.
Sample Input:
90
Sample Output:
It is 3 hours 0 minutes. 
*/


/*
Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m.
Входные данные
На вход программе подается целое число d (0 < d < 360).
Выходные данные
Выведите на экран фразу:
It is ... hours ... minutes.
Вместо многоточий программа должна выводить значения h и m, отделяя их от слов ровно одним пробелом.
Sample Input:
90
Sample Output:
It is 3 hours 0 minutes.
 */

package main
import "fmt"
func main() {
var d int 
fmt.Scan(&d)
var h int = d/30
var m int = (d%30)*2
fmt.Println("It is", h, "hours", m, "minutes.")  
}
