package main

import (
	"encoding/json"
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

func main() {
	godotenv.Load()

	OPEN_WEATHER_API_KEY := os.Getenv("OPEN_WEATHER_API_KEY")
	city := ""
	fmt.Println("Enter City")
	fmt.Scanf("%s", &city)
	fmt.Println("Getting Weather For : ", city)
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, OPEN_WEATHER_API_KEY)
	fmt.Println(url)
	resp, err := http.Get(url)
	weatherResp := new(Response)
	json.NewDecoder(resp.Body).Decode(&weatherResp)

	fmt.Println(weatherResp)

	if err != nil {
		fmt.Println(err.Error())
	}

	if err != nil {
		return
	}

}
