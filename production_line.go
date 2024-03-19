package main

import (
	"fmt"
	"sync"
	"time"
)

// 輸出員工處理肉的資訊
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

// 處理肉的產線，有五位獨立作業的員工
// Parameter meats is a slice of Meat struct.
func ProductionLine(meats []Meat) {
	var wg sync.WaitGroup
	wg.Add(len(meats))

	meatChannel := make(chan Meat, len(meats))
	for i := 0; i < len(meats); i++ {
		meatChannel <- meats[i]
	}

	workA := NewWork(Emploee{Id: "A"}, meatChannel, &wg)
	go workA.ProcessingMeats(len(meats))

	workB := NewWork(Emploee{Id: "B"}, meatChannel, &wg)
	go workB.ProcessingMeats(len(meats))

	workC := NewWork(Emploee{Id: "C"}, meatChannel, &wg)
	go workC.ProcessingMeats(len(meats))

	workD := NewWork(Emploee{Id: "D"}, meatChannel, &wg)
	go workD.ProcessingMeats(len(meats))

	workE := NewWork(Emploee{Id: "E"}, meatChannel, &wg)
	go workE.ProcessingMeats(len(meats))

	wg.Wait()
}
