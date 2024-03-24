package main

import (
	"sync"
	"testing"
)

func TestProductionLine(t *testing.T) {
	t.Run("test processing meats", func(t *testing.T) {
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
	})
	t.Run("test production line", func(t *testing.T) {
		ProductionLine(GetEmploees([]string{"A", "B", "C", "D", "E"}), GetRawMeat(5, 3, 1))
	})
}
