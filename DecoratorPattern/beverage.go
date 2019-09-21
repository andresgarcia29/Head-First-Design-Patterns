package main

type Beverage interface {
	getDescription() string
	getCost() int
}

type Condiments struct {
	Beverage
}

type Expresso struct {
	description string
	cost        int
}

func (e *Expresso) getCost() int {
	return 10
}

func (e *Expresso) getDescription() string {
	return "I'm an expresso"
}

type HouseBlend struct {
	Beverage
	description string
	cost        int
}

func (h *HouseBlend) getCost() int {
	return 10
}

func (h *HouseBlend) getDescription() string {
	return "I'm an House Blend"
}

type Mocha struct {
	Condiments
}

func NewMocha(beverage Beverage) Beverage {
	mocha := new(Mocha)
	mocha.Beverage = beverage
	return mocha
}

func (m *Mocha) getDescription() string {
	return m.Beverage.getDescription() + ", Mocha"
}

func (m *Mocha) getCost() int {
	return m.Beverage.getCost() + 5
}
