// Package weather provides helpers for
// creating weather forecasts.
package weather

// CurrentCondition stores the last condition for the forecast.
var CurrentCondition string

// CurrentLocation stores the last location for the forescast.
var CurrentLocation string

// Forecast takes city and condition as parameters, and returns the weather forecast for it.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
