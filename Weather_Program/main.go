package main

import (
	"fmt"
	"github.com/paulgureghian/Go_Programs/Weather_Program/api"
)

var OWM_API_KEY string = ""
var POSITIONSTACK_KEY string = ""
var LOCATION string = "London, United Kingdom"

func main() {
	// Create our channels.
	locationReturn := make(chan api.PositionStackModel)
	weatherReturn := make(chan api.OpenWeatherModel)

	// Get location data and wait for finish.
	go api.GetLocation(POSITIONSTACK_KEY, LOCATION, locationReturn)
	// Block until we receive this.
	LocationModel := <-locationReturn

	// Get weather data.
	go api.GetWeather(OWM_API_KEY, fmt.Sprintf("%f", LocationModel.Data[0].Latitude), fmt.Sprintf("%f", LocationModel.Data[0].Longitude), weatherReturn)
	// Block until we receive this.
	WeatherModel := <-weatherReturn

	// Print current temperature.
	println()
	fmt.Printf("Current Temperature: %.1f\u00B0C\n", WeatherModel.Current.Temp)
	println()
}
