package model

type Ingredient struct {
	Quantity float64
	Unit     string
	Name     string
	Category string
}

type Instruction struct {
	Content string
}

type Recipe struct {
	Title        string
	Author       string
	Category     string
	Ingredients  []Ingredient
	Instructions []Instruction
}
