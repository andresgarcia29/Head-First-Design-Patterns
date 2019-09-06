package main

type Observer interface {
	update(temp float64, humidity float64, pressure float64)
}

type DisplayElements interface {
	display()
}

type Subject interface {
	AddObserver(Observer)
	RemoveObserver(Observer)
	NotifyObservers()
}

func main() {
	weatherData := WeatherData{
		temp:     100,
		humidity: 100,
		pressure: 100,
	}

	_ = NewCurrentConditionDisplayer(&weatherData)
	_ = NewStatisticsDisplayer(&weatherData)

	weatherData.SetMeasurements(200, 200, 200)
	weatherData.SetMeasurements(300, 200, 200)
	weatherData.SetMeasurements(400, 400, 400)
}
