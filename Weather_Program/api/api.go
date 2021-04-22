package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type PositionStackModel struct {
	Data []struct {
		Latitude           float64     `json:"latitude"`
		Longitude          float64     `json:"longitude"`
		Type               string      `json:"type"`
		Name               string      `json:"name"`
		Number             interface{} `json:"number"`
		PostalCode         interface{} `json:"postal_code"`
		Street             interface{} `json:"street"`
		Confidence         int         `json:"confidence"`
		Region             string      `json:"region"`
		RegionCode         interface{} `json:"region_code"`
		County             string      `json:"county"`
		Locality           string      `json:"locality"`
		AdministrativeArea string      `json:"administrative_area"`
		Neighbourhood      interface{} `json:"neighbourhood"`
		Country            string      `json:"country"`
		CountryCode        string      `json:"country_code"`
		Continent          string      `json:"continent"`
		Label              string      `json:"label"`
	} `json:"data"`
}

type OpenWeatherModel struct {
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	Timezone       string  `json:"timezone"`
	TimezoneOffset int     `json:"timezone_offset"`
	Current        struct {
		Dt         int     `json:"dt"`
		Sunrise    int     `json:"sunrise"`
		Sunset     int     `json:"sunset"`
		Temp       float64 `json:"temp"`
		FeelsLike  float64 `json:"feels_like"`
		Pressure   int     `json:"pressure"`
		Humidity   int     `json:"humidity"`
		DewPoint   float64 `json:"dew_point"`
		Uvi        int     `json:"uvi"`
		Clouds     int     `json:"clouds"`
		Visibility int     `json:"visibility"`
		WindSpeed  float64 `json:"wind_speed"`
		WindDeg    int     `json:"wind_deg"`
		Weather    []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
	} `json:"current"`
	Minutely []struct {
		Dt            int `json:"dt"`
		Precipitation int `json:"precipitation"`
	} `json:"minutely"`
	Hourly []struct {
		Dt         int     `json:"dt"`
		Temp       float64 `json:"temp"`
		FeelsLike  float64 `json:"feels_like"`
		Pressure   int     `json:"pressure"`
		Humidity   int     `json:"humidity"`
		DewPoint   float64 `json:"dew_point"`
		Uvi        int     `json:"uvi"`
		Clouds     int     `json:"clouds"`
		Visibility int     `json:"visibility"`
		WindSpeed  float64 `json:"wind_speed"`
		WindDeg    int     `json:"wind_deg"`
		Weather    []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Pop float64 `json:"pop"`
	} `json:"hourly"`
	Daily []struct {
		Dt      int `json:"dt"`
		Sunrise int `json:"sunrise"`
		Sunset  int `json:"sunset"`
		Temp    struct {
			Day   float64 `json:"day"`
			Min   float64 `json:"min"`
			Max   float64 `json:"max"`
			Night float64 `json:"night"`
			Eve   float64 `json:"eve"`
			Morn  float64 `json:"morn"`
		} `json:"temp"`
		FeelsLike struct {
			Day   float64 `json:"day"`
			Night float64 `json:"night"`
			Eve   float64 `json:"eve"`
			Morn  float64 `json:"morn"`
		} `json:"feels_like"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		DewPoint  float64 `json:"dew_point"`
		WindSpeed float64 `json:"wind_speed"`
		WindDeg   int     `json:"wind_deg"`
		Weather   []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds int     `json:"clouds"`
		Pop    float64 `json:"pop"`
		Uvi    float64 `json:"uvi"`
	} `json:"daily"`
	Alerts []struct {
		SenderName  string `json:"sender_name"`
		Event       string `json:"event"`
		Start       int    `json:"start"`
		End         int    `json:"end"`
		Description string `json:"description"`
	} `json:"alerts"`
}

func GetLocation(key, place string, returnModel chan PositionStackModel) {

	var Model PositionStackModel

	requestData := url.Values{}
	requestData.Add("access_key", key)
	requestData.Add("query", place)

	request, e := http.NewRequest("GET", "http://api.positionstack.com/v1/forward?"+requestData.Encode(), nil)
	if e != nil {
		log.Fatalf("Error: %s", e)
	}

	response, e := (&http.Client{}).Do(request)
	if e != nil {
		log.Fatalf("Error: %s", e)
	}

	defer response.Body.Close()

	BodyHandler(response.Body, &Model)

	returnModel <- Model

}

func GetWeather(key, lat, lon string, returnModel chan OpenWeatherModel) {

	var Model OpenWeatherModel

	requestData := url.Values{}
	requestData.Add("lat", lat)
	requestData.Add("lon", lon)
	requestData.Add("appid", key)
	requestData.Add("units", "metric")

	request, e := http.NewRequest("GET", "https://api.openweathermap.org/data/2.5/onecall?"+requestData.Encode(), nil)
	if e != nil {
		log.Fatalf("Error: %s", e)

	}

	response, e := (&http.Client{}).Do(request)
	if e != nil {
		log.Fatalf("Error: %s", e)
	}

	defer response.Body.Close()

	BodyHandler(response.Body, &Model)

	returnModel <- Model

}

func BodyHandler(r io.ReadCloser, model interface{}) {

	body, e := ioutil.ReadAll(r)
	if e != nil {
		log.Fatalf("Error: %s", e)
	}

	json.Unmarshal(body, model)
}
