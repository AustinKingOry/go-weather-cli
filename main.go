package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
)

const apiKey = "b7f63c36e08263b6cfac499387649412"

type WeatherResponse struct {
    Name string `json:"name"`
    Main struct {
        Temp     float64 `json:"temp"`
        Pressure int     `json:"pressure"`
        Humidity int     `json:"humidity"`
    } `json:"main"`
    Weather []struct {
        Description string `json:"description"`
    } `json:"weather"`
}

func getWeather(city string) (*WeatherResponse, error) {
    url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var weatherResponse WeatherResponse
    if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
        return nil, err
    }

    return &weatherResponse, nil
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please provide a city name")
        os.Exit(1)
    }

    city := os.Args[1]
    weather, err := getWeather(city)
    if err != nil {
        log.Fatalf("Failed to get weather: %v", err)
    }

    fmt.Printf("Weather in %s:\n", weather.Name)
    fmt.Printf("Temperature: %.2f°C\n", weather.Main.Temp)
    fmt.Printf("Pressure: %d hPa\n", weather.Main.Pressure)
    fmt.Printf("Humidity: %d%%\n", weather.Main.Humidity)
    fmt.Printf("Description: %s\n", weather.Weather[0].Description)
}
