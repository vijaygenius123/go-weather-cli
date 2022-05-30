package main

import (
	"encoding/json"
	"flag"
	fmt "fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

type Response struct {
	Weather []struct {
		ID          string `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float32 `json:"temp"`
	} `json:"main"`
}

func getWeather(city string) *Response {
	godotenv.Load()
	OPEN_WEATHER_API_KEY := os.Getenv("OPEN_WEATHER_API_KEY")
	fmt.Println("Getting weather data for ", city)
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, OPEN_WEATHER_API_KEY)
	resp, _ := http.Get(url)
	weatherResp := new(Response)
	json.NewDecoder(resp.Body).Decode(&weatherResp)
	return weatherResp

}

func main() {

	city := flag.String("city", "Edinburgh", "the city you want to retrieve weather data for")
	flag.Parse()

	weatherResp := getWeather(*city)
	fmt.Println(weatherResp.Weather[0].Main, weatherResp.Main.Temp)

}
