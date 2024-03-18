package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func ProcessingMeat(ep Emploee) {
	m := ep.Meat
	fmt.Printf("%s 在 %s 取得%s\n", ep.Id, time.Now().Format(time.DateTime), m.Name)

	// 處理時間
	time.Sleep(time.Duration(m.ProcessingTimeSeconds) * time.Second)

	fmt.Printf("%s 在 %s 處理完%s\n", ep.Id, time.Now().Format(time.DateTime), m.Name)
}

type Work struct {
	Emploee
	meatChannel chan Meat
	wg          *sync.WaitGroup
}

func NewWork(emploee Emploee, meatChannel chan Meat, wg *sync.WaitGroup) Work {
	return Work{
		Emploee:     emploee,
		meatChannel: meatChannel,
		wg:          wg,
	}
}

func (w Work) ProcessingMeats(count int) {
	for i := 0; i < count; i++ {
		w.Emploee.Meat = <-w.meatChannel
		ProcessingMeat(w.Emploee)
		w.wg.Done()
	}
}

func TestProcessingMeats(t *testing.T) {
	meats := GetRawMeat(1, 1, 1)

	var wg sync.WaitGroup
	wg.Add(len(meats))

	meatChannel := make(chan Meat, len(meats))
	for i := 0; i < len(meats); i++ {
		meatChannel <- meats[i]
	}

	work := NewWork(Emploee{Id: "A"}, meatChannel, &wg)
	go work.ProcessingMeats(len(meats))

	wg.Wait()
}
