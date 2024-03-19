package main

import (
	"sync"
	"testing"
)

func TestProcessingMeats(t *testing.T) {
	meats := GetRawMeat(2, 2, 1)

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

	wg.Wait()
}
