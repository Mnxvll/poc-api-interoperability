package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Response struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`

	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Name string `json:"name"`
}

// CheckConnection tests the connection to OpenWeatherMap
func CheckConnection() {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	city := os.Getenv("WEATHER_CITY")

	if apiKey == "" || city == "" {
		fmt.Println("Missing OPENWEATHER_API_KEY or WEATHER_CITY environment variables")
		return
	}
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("request error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("the server responded with code: %d \n", resp.StatusCode)
		return
	}

	var data Response
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("error parsing JSON: %v\n", err)
		return
	}

	celsius := data.Main.Temp - 273.15

	fmt.Printf("city: %s\nweather: %s (%s)\ntemperature (Celsius): %.2f\n",
		data.Name, data.Weather[0].Main, data.Weather[0].Description, celsius)
}
