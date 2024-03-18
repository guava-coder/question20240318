package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func ProcessingMeat(ep Emploee, meatChannel chan Meat) {
	ep.Meat = <-meatChannel
	m := ep.Meat
	fmt.Printf("%s 在 %s 取得%s\n", ep.Id, time.Now().Format(time.DateTime), m.Name)

	// 處理時間
	time.Sleep(time.Duration(m.ProcessingTimeSeconds) * time.Second)

	fmt.Printf("%s 在 %s 處理完%s\n", ep.Id, time.Now().Format(time.DateTime), m.Name)
}

func TestProcessingMeats(t *testing.T) {
	meats := []Meat{
		{
			Id:                    "beef-1",
			Name:                  "牛肉",
			ProcessingTimeSeconds: 1,
		},
		{
			Id:                    "pork-1",
			Name:                  "豬肉",
			ProcessingTimeSeconds: 2,
		},
	}

	var wg sync.WaitGroup
	wg.Add(len(meats))

	meatChannel := make(chan Meat, len(meats))

	for i := 0; i < len(meats); i++ {
		meatChannel <- meats[i]
		ProcessingMeat(Emploee{Id: "A"}, meatChannel)
		wg.Done()
	}

	wg.Wait()
}
