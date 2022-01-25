package main

import (
	"fmt"
	"time"
)

var N,W,C,M,a,d,p int
var StoreStocks = map[int]map[int]int{}
var GoodsCosts = map[int]int{}
var Goods = [] int{}

func getMinCostAndStock(mapForSorting map[int]int) (int,int){
	var mincostDelivery, minstockNumber, i int
	i=0
	for key, value := range mapForSorting{
		i++
		if (i==1){mincostDelivery = value; minstockNumber = key; continue }
		if (value < mincostDelivery) {mincostDelivery = value; minstockNumber = key; continue}
		if (value == mincostDelivery) {if (key<minstockNumber) {minstockNumber = key}}
	}
	return minstockNumber, mincostDelivery
}

func getStockWithLowestPrice(good int) (int,int) {
var DeliveryCostsForStocks = map[int]int{}
var  minstockNumber, totalCostDelivery int

	for key, value := range StoreStocks {
		if delCost, found := value[good]; found {
			DeliveryCostsForStocks[key] = delCost
		}
	}

	if (len(DeliveryCostsForStocks) == 0) {return -1,-1} else {
		minstockNumber, totalCostDelivery = getMinCostAndStock(DeliveryCostsForStocks)
		totalCostDelivery += GoodsCosts[good]
	}
	return minstockNumber,totalCostDelivery
}

func main() {
fmt.Scan(&N)

start := time.Now()

for i:=0; i<N; i++{
	fmt.Scan(&W)
	fmt.Scan(&C)
	DeliveryCost := map[int]int{}
	for i:=0; i<C; i++{
		fmt.Scan(&a)
		fmt.Scan(&d)
		DeliveryCost[a]=d
	}
	StoreStocks[W]=DeliveryCost
}
fmt.Scan(&M)
for i:=0; i<M; i++{
	fmt.Scan(&a)
	fmt.Scan(&p)
	GoodsCosts[a]=p
	Goods = append(Goods,a)
}

for _,elem := range Goods{
	fmt.Println(getStockWithLowestPrice(elem))
}

fmt.Println(time.Since(start))

}
