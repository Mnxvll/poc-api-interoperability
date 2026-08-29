package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type WeatherResponse struct {
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

func main() {

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	connectToWeatherAPI()

	r := gin.Default()

	// Simple ping route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

// Connection test to OpenWeatherMap
func connectToWeatherAPI() {
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

	var data WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("error parsing JSON: %v\n", err)
		return
	}

	fmt.Printf("city: %s\nweather: %s (%s)\ntemperature (Kelvin): %.2f\n",
		data.Name, data.Weather[0].Main, data.Weather[0].Description, data.Main.Temp)
}
