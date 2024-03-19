package main

import (
	"fmt"
	"sync"
	"time"
)

//	員工處理肉，並輸出處理的資訊
//
// ep Emploee
func processingMeat(ep Emploee) {
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

// 員工開始根據肉的數量來處理肉，直到處理完所有肉為止
//
// count int 肉的數量
func (w Work) ProcessingMeats(count int) {
	for i := 0; i < count; i++ {
		w.Emploee.Meat = <-w.meatChannel
		processingMeat(w.Emploee)
		w.wg.Done()
	}
}

// 處理肉的產線，由獨立作業的員工各自處理
// Parameter meats is a slice of Meat struct.
func ProductionLine(emploees []Emploee, meats []Meat) {
	// 根據肉的數量設定等待次數
	var wg sync.WaitGroup
	wg.Add(len(meats))

	// 存放肉的 channel
	meatChannel := make(chan Meat, len(meats))
	for i := 0; i < len(meats); i++ {
		meatChannel <- meats[i]
	}

	// 根據員工數量產生工作
	works := make([]Work, 0)
	for i := 0; i < len(emploees); i++ {
		works = append(works, NewWork(emploees[i], meatChannel, &wg))
	}

	// 開始處理肉
	for i := 0; i < len(works); i++ {
		go works[i].ProcessingMeats(len(meats))
	}

	// 等待所有肉都處理完
	wg.Wait()
}
