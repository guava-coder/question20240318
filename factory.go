package main

import "fmt"

func GetRawMeat() (meats []Meat) {
	var beafs, porks, chickens []Meat

	for i := 0; i < 10; i++ {
		beaf := Meat{
			Id:                    fmt.Sprintf("beaf-%d", i),
			Name:                  "牛肉",
			ProcessingTimeSeconds: 1,
		}
		beafs = append(beafs, beaf)
	}

	for i := 0; i < 7; i++ {
		pork := Meat{
			Id:                    fmt.Sprintf("pork-%d", i),
			Name:                  "豬肉",
			ProcessingTimeSeconds: 2,
		}
		porks = append(porks, pork)
	}

	for i := 0; i < 5; i++ {
		chicken := Meat{
			Id:                    fmt.Sprintf("chicken-%d", i),
			Name:                  "雞肉",
			ProcessingTimeSeconds: 3,
		}
		chickens = append(chickens, chicken)
	}

	meats = append(meats, beafs...)
	meats = append(meats, porks...)
	meats = append(meats, chickens...)

	return
}

func GetEmploees() []Emploee {
	ids := []string{
		"A", "B", "C", "D", "E",
	}
	emploees := make([]Emploee, 0)
	for i := 0; i < len(ids); i++ {
		emploees = append(emploees, Emploee{
			Id:   fmt.Sprintf(ids[i]),
			Meat: Meat{},
		})
	}
	return emploees
}
