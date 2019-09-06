package main

type WeatherData struct {
	observers []Observer
	temp      float64
	humidity  float64
	pressure  float64
}

func (w *WeatherData) SetMeasurements(temp float64, humidity float64, pressure float64) {
	w.temp = temp
	w.humidity = humidity
	w.pressure = pressure
	w.NotifyObservers()
}

func (w *WeatherData) AddObserver(o Observer) {
	w.observers = append(w.observers, o)
}

func (w *WeatherData) RemoveObserver(o Observer) {
	for i, observer := range w.observers {
		if observer == o {
			w.observers = append(w.observers[:i], w.observers[:i]...)
		}
	}
}

func (w *WeatherData) NotifyObservers() {
	for _, observer := range w.observers {
		observer.update(w.temp, w.humidity, w.pressure)
	}
}
