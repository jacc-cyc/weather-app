package weather

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

type obj2 struct {
	GeneralSituation string `json:"generalSituation"`

	WeatherForecast []struct {
		ForecastDate    string `json:"forecastDate"`
		Week            string `json:"week"`
		ForecastWind    string `json:"forecastWind"`
		ForecastWeather string `json:"forecastWeather"`

		ForecastMaxtemp struct {
			Value int    `json:"value"`
			Unit  string `json:"unit"`
		} `json:"forecastMaxtemp"`

		ForecastMintemp struct {
			Value int    `json:"value"`
			Unit  string `json:"unit"`
		} `json:"forecastMintemp"`

		ForecastMaxrh struct {
			Value int    `json:"value"`
			Unit  string `json:"unit"`
		} `json:"forecastMaxrh"`

		ForecastMinrh struct {
			Value int    `json:"value"`
			Unit  string `json:"unit"`
		} `json:"forecastMinrh"`
	} `json:"weatherForecast"`

	UpdateTime string `json:"updateTime"`
}

var url = "https://data.weather.gov.hk/weatherAPI/opendata/weather.php?dataType=fnd"
var lastUpdate string

//Info is to show the general weather info of coming days
func Info(c *gin.Context) {
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Println(err1)
		os.Exit(1)
	}

	data, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		log.Println(err2)
		os.Exit(1)
	}

	value := gjson.Get(string(data), "generalSituation")
	days := gjson.Get(string(data), "weatherForecast.#.forecastDate")
	lastUpdate = (gjson.Get(string(data), "updateTime")).String()

	message := "General Weather Situation of coming days:\n\n"
	message += value.String()

	message += "\n\nAvailable days for weather checking:\n\n"
	message += days.String()

	message += "\n\nLast Update: " + lastUpdate

	c.String(http.StatusOK, message)
}

//DayInfo is to show the detail weather info of a specific day
func DayInfo(c *gin.Context) {
	var obj2 obj2
	var wk, wind, weather string
	var maxtmp, mintmp, maxrh, minrh int

	date := c.Param("date")

	response, err1 := http.Get(url)
	if err1 != nil {
		log.Println(err1)
		os.Exit(1)
	}

	data, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		log.Println(err2)
		os.Exit(1)
	}

	err3 := json.Unmarshal([]byte(data), &obj2)
	if err3 != nil {
		log.Println(err3)
		os.Exit(1)
	}

	lastUpdate = (gjson.Get(string(data), "updateTime")).String()
	valid := false

	for i := 0; i < 9; i++ {
		if obj2.WeatherForecast[i].ForecastDate == date {
			wk = obj2.WeatherForecast[i].Week
			weather = obj2.WeatherForecast[i].ForecastWeather
			wind = obj2.WeatherForecast[i].ForecastWind
			maxtmp = obj2.WeatherForecast[i].ForecastMaxtemp.Value
			mintmp = obj2.WeatherForecast[i].ForecastMintemp.Value
			maxrh = obj2.WeatherForecast[i].ForecastMaxrh.Value
			minrh = obj2.WeatherForecast[i].ForecastMinrh.Value
			valid = true
			break
		}
	}

	if !valid {
		message := "Please check /info for available days."
		c.String(http.StatusOK, message)
	} else {
		message := date

		message += "\n" + wk
		message += "\n\nWind: " + wind
		message += "\nWeather: " + weather

		message += "\n\nMax. Temp: " + strconv.Itoa(maxtmp) + "°C"
		message += "\nMin. Temp: " + strconv.Itoa(mintmp) + "°C"
		message += "\n\nMax. Relative Humidity: " + strconv.Itoa(maxrh) + "%"
		message += "\nMin. Relative Humidity: " + strconv.Itoa(minrh) + "%"

		message += "\n\nLast Update: " + lastUpdate

		c.String(http.StatusOK, message)
	}
}
