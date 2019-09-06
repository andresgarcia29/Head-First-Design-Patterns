package main

import (
	"fmt"
)

type CurrentConditionDisplayer struct {
	subject  Subject
	temp     float64
	humidity float64
}

func NewCurrentConditionDisplayer(s Subject) *CurrentConditionDisplayer {
	currentConditionDisplayer := &CurrentConditionDisplayer{}
	currentConditionDisplayer.subject = s
	s.AddObserver(currentConditionDisplayer)
	return currentConditionDisplayer
}

func (c *CurrentConditionDisplayer) update(temp float64, humidity float64, pressure float64) {
	c.temp = temp
	c.humidity = humidity
	c.display()
}

func (c *CurrentConditionDisplayer) display() {
	fmt.Printf("CurrentConditionDisplayer:\tThe current condition is %f with %f of humidity\n", c.temp, c.humidity)
}

type StatisticsDisplayer struct {
	subject  Subject
	temp     float64
	humidity float64
	pressure float64
}

func NewStatisticsDisplayer(s Subject) *StatisticsDisplayer {
	statisticsDisplayer := &StatisticsDisplayer{}
	statisticsDisplayer.subject = s
	s.AddObserver(statisticsDisplayer)
	return statisticsDisplayer
}

func (s *StatisticsDisplayer) update(temp float64, humidity float64, pressure float64) {
	s.temp = temp
	s.humidity = humidity
	s.pressure = pressure
	s.display()
}

func (s *StatisticsDisplayer) display() {
	fmt.Printf("StatisticsDisplayer:\t\tThe current mean of all the measurements is %f\n", (s.temp+s.humidity+s.pressure)/3)
}
