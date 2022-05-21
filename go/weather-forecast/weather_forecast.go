// Package weather - forecasting the weather.
package weather

// CurrentCondition stores current weather conditions.
var CurrentCondition string
// CurrentLocation stores the current location.
var CurrentLocation string

// Forecast will return the current weather condition for the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
