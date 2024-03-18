package main

type Meat struct {
	Id                    string
	Name                  string
	ProcessingTimeSeconds int
}

type Emploee struct {
	Id string
	Meat
}
