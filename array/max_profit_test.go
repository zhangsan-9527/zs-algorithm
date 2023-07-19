package array

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{2, 10, 8, 2, 4, 5, 3} // 20 + 40 +33
	result := MaxProfit(prices)
	fmt.Println(result)
}

func TestMaxProfitUp(t *testing.T) {
	//prices := []int{2, 10, 8, 2, 4, 5, 3}
	prices := []int{1, 2, 3, 4, 5}
	result := MaxProfitUp(prices)
	fmt.Println(result)

}
