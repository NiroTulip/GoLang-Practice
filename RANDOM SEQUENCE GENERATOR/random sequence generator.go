package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

func generator(n,max int) [] int {
	rand.Seed(time.Now().UnixNano())
	randmap := make(map[int]bool)
	for len(randmap) < n {
		randmap[rand.Intn(max)] = true
	}
	result := make([]int,0)
	for i, _ := range randmap {
		result = append(result, i)
	}
	sort.Ints(result)
	return result
}

func getStrFromSlice(slice [] int) string{
	var result = ""
	for _,e := range slice{
		result=result + strconv.Itoa(e) + ","
	}
	result = strings.TrimRight(result, ",")
	return result
}

func main() {
fmt.Println(getStrFromSlice(generator(10, 100)))
}
