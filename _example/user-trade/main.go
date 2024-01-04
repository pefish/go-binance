package main

import (
	"context"
	"fmt"
	"github.com/pefish/go-binance/futures"
	"log"
)

func main() {
	err := do()
	if err != nil {
		log.Fatal(err)
	}
}

func do() error {
	startTime := 1704246284744
	results, err := futures.NewClient(
		"",
		"",
	).
		NewListAccountTradeService().
		StartTime(int64(startTime)). // 返回内容包含开始时间
		Limit(1000).
		Do(context.Background())
	if err != nil {
		return err
	}
	if len(results) == 0 {
		fmt.Printf("No records.\n")
		return nil
	}

	sum := 0
	for _, result := range results {
		sum++
		fmt.Println(result.OrderID, result.ID, uint64(result.Time))
	}
	fmt.Println(sum)
	return nil
}
