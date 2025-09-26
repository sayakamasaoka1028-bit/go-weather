package main

import (
    "fmt"
    "strings"
)

func main() {
    weatherData := map[string]string{
        "tokyo":   "晴れ ☀️",
        "osaka":   "曇り ☁️",
        "sapporo": "雪 ❄️",
        "fukuoka": "雨 🌧️",
    }

    fmt.Print("都市名をカンマ区切りで入力してください: ")
    var input string
    fmt.Scanln(&input)

    cities := strings.Split(input, ",")

    for _, city := range cities {
        city = strings.TrimSpace(strings.ToLower(city))
        weather, ok := weatherData[city]
        displayCity := strings.Title(city)
        if ok {
            fmt.Printf("%sの天気: %s\n", displayCity, weather)
        } else {
            fmt.Printf("%sのデータはありません\n", displayCity)
        }
    }
}
