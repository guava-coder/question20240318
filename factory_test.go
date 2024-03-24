package main

import "testing"

func TestFactory(t *testing.T) {
	t.Run("get raw meats", func(t *testing.T) {
		meats := GetRawMeat(10, 7, 5)

		beefCounts, porkCounts, chickenCounts := 0, 0, 0
		for _, meat := range meats {
			switch meat.Name {
			case "牛肉":
				beefCounts++
			case "豬肉":
				porkCounts++
			case "雞肉":
				chickenCounts++
			}
		}

		if beefCounts != 10 {
			t.Errorf("beef counts = %v, want %v", beefCounts, 10)
		}
		if porkCounts != 7 {
			t.Errorf("pork counts = %v, want %v", porkCounts, 7)
		}
		if chickenCounts != 5 {
			t.Errorf("chicken counts = %v, want %v", chickenCounts, 5)
		}
	})
	t.Run("get emploees", func(t *testing.T) {
		emploees := GetEmploees([]string{"A", "B", "C", "D", "E"})
		if len(emploees) != 5 {
			t.Errorf("len(emploees) = %v, want %v", len(emploees), 5)
		}
	})
}
