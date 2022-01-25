package main

import (
	"fmt"
	"time"
)

var N,W,C,M,a,d,p int

var MapMinStockForGoods = map[int]int{}
var MapMinCostForGoods = map[int]int{}

var GoodsCosts = map[int]int{}
var Goods = [] int{}

func getLowestDeliveryCostAndStock(good int) (int,int){
	return MapMinStockForGoods[good], MapMinCostForGoods[good]
}

func getBestTotalPriceAndStock(good int) (int,int){
	var minStock, minDeliveryCost, bestTotalCost int
	if _, ok := MapMinStockForGoods[good]; !ok{
		return -1,-1
	} else{
		minStock, minDeliveryCost = getLowestDeliveryCostAndStock(good)
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
			if _, ok := MapMinStockForGoods[a]; !ok{
				MapMinStockForGoods[a]=W
				MapMinCostForGoods[a]=d
			}else{
				if (d < MapMinCostForGoods[a]) {
					MapMinStockForGoods[a]=W
					MapMinCostForGoods[a]=d
				} else if (d == MapMinCostForGoods[a]){
					if (W < MapMinStockForGoods[a]) {MapMinStockForGoods[a] = W}
				}
			}
		}
	}
	fmt.Scan(&M)
	for i:=0; i<M; i++{
		fmt.Scan(&a)
		fmt.Scan(&p)
		Goods = append(Goods, a)
		GoodsCosts[a]=p
	}

for _,good := range Goods{
	fmt.Println(getBestTotalPriceAndStock(good))
}

	fmt.Println(time.Since(start))
}
