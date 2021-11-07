package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Weather App")

	label1 := canvas.NewText("Weather Details", color.White)
	label1.TextStyle = fyne.TextStyle{Bold: true}

	img := canvas.NewImageFromFile("weather.png")
	img.FillMode = canvas.ImageFillOriginal

	label2 := widget.NewEntry()
	label2.TextStyle = fyne.TextStyle{Bold: true}
	label2.SetPlaceHolder("-")
	label2.Disable()

	label3 := widget.NewEntry()
	label3.TextStyle = fyne.TextStyle{Bold: true}
	label3.SetPlaceHolder("-")
	label3.Disable()

	label4 := widget.NewEntry()
	label4.TextStyle = fyne.TextStyle{Bold: true}
	label4.SetPlaceHolder("-")
	label4.Disable()

	label5 := widget.NewEntry()
	label5.TextStyle = fyne.TextStyle{Bold: true}
	label5.SetPlaceHolder("-")
	label5.Disable()

	selector := widget.NewSelect([]string{"delhi", "noida", "faridabad", "gurgaon"}, func(s string) {
		// fmt.Print(s)
		res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + s + "&appid=-------------api-key-------------")
		if err != nil {
			fmt.Print(err)
		} else {
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Print(err)
			} else {
				weather, err := UnmarshalWeather(body)
				if err != nil {
					fmt.Print(err)
				} else {
					// fmt.Print(weather)
					label2.SetText(weather.Sys.Country)
					label3.SetText(fmt.Sprintf("%.2f", weather.Wind.Speed))
					label4.SetText(fmt.Sprintf("%.2f", weather.Main.Temp))
					label5.SetText(fmt.Sprintf("%d", weather.Main.Humidity))
				}
			}
		}
	})

	myWindow.SetContent(container.NewVBox(
		container.New(
			layout.NewCenterLayout(),
			label1,
		),
		container.New(
			layout.NewCenterLayout(),
			img,
		),
		container.New(
			layout.NewGridLayout(1),
			selector,
		),
		container.New(
			layout.NewGridLayout(2),
			canvas.NewText("Country: ", color.White),
			label2,
		),
		container.New(
			layout.NewGridLayout(2),
			canvas.NewText("Wind Speed: ", color.White),
			label3,
		),
		container.New(
			layout.NewGridLayout(2),
			canvas.NewText("Temprature: ", color.White),
			label4,
		),
		container.New(
			layout.NewGridLayout(2),
			canvas.NewText("Humidity: ", color.White),
			label5,
		),
	))

	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(400, 600))
	myWindow.ShowAndRun()
}

func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Coord      Coord            `json:"coord"`
	Weather    []WeatherElement `json:"weather"`
	Base       string           `json:"base"`
	Main       Main             `json:"main"`
	Visibility int64            `json:"visibility"`
	Wind       Wind             `json:"wind"`
	Clouds     Clouds           `json:"clouds"`
	Dt         int64            `json:"dt"`
	Sys        Sys              `json:"sys"`
	Timezone   int64            `json:"timezone"`
	ID         int64            `json:"id"`
	Name       string           `json:"name"`
	Cod        int64            `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
}

type Sys struct {
	Type    int64  `json:"type"`
	ID      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type WeatherElement struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
}
