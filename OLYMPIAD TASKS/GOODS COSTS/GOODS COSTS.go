package main

import (
	"fmt"
	"time"
)

var N,W,C,M,a,d,p int

var StocksDeliveryCosts = map[int]int{}
var GoodsDeliveryCosts = map[int]map[int]int{}
var GoodsCosts = map[int]int{}
var Goods = [] int{}

func getLowestDeliveryCostAndStock(SDC map[int]int) (int,int){
	var minStock, minCost int
		i:=0
		for stock, delCost := range SDC{
			i++
			if (i==1){
				minStock = stock; minCost = delCost; continue
			}
			if (delCost < minCost){
				minCost = delCost;
				minStock = stock
			}
			if (delCost == minCost && minStock > stock){
				minStock = stock
			}
		}
	return minStock, minCost
}

func getBestTotalPriceAndStock(good int) (int,int){
	var minStock, minDeliveryCost, bestTotalCost int
	if SDC, ok := GoodsDeliveryCosts[good]; !ok{
		return -1,-1
	} else {
		minStock, minDeliveryCost = getLowestDeliveryCostAndStock(SDC)
	}
	bestTotalCost = minDeliveryCost + GoodsCosts[good]
	return  minStock, bestTotalCost
}

func main() {
	fmt.Scan(&N)

	start := time.Now()

	for i:=0; i<N; i++{
		fmt.Scan(&W)
		fmt.Scan(&C)

		for i:=0; i<C; i++{
			fmt.Scan(&a)
			fmt.Scan(&d)
			SDC := map[int]int{}
			SDC[W] = d
			if SDC2, ok := GoodsDeliveryCosts[a]; !ok{
				SDC2 = SDC
				GoodsDeliveryCosts[a] = SDC2
			} else {GoodsDeliveryCosts[a][W]=d}
		}
	}
	fmt.Scan(&M)
	for i:=0; i<M; i++{
		fmt.Scan(&a)
		fmt.Scan(&p)
		GoodsCosts[a]=p
		Goods = append(Goods,a)
	}

for _,elem := range Goods{
	fmt.Println(getBestTotalPriceAndStock(elem))
}

	fmt.Println(time.Since(start))
}
