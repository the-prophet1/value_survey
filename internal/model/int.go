package model

type Int *innerInt

type innerInt struct {
	Data       int
	Ingredient map[string]int
}
