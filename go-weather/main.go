package main

import (
    "fmt"
    "strings"
)

func main() {
    weatherData := map[string]string{
        "tokyo":   "晴れ",
        "osaka":   "曇り",
        "sapporo": "雪",
        "fukuoka": "雨",
    }

    var city string
    fmt.Print("都市名を入力してください: ")
    fmt.Scanln(&city)

    // 入力を小文字に変換
    city = strings.ToLower(city)

    weather, ok := weatherData[city]
    if ok {
        // 表示用に頭文字だけ大文字に変換
        displayCity := strings.Title(city)
        fmt.Printf("%sの天気: %s\n", displayCity, weather)
    } else {
        fmt.Println("その都市のデータはありません")
    }
}
