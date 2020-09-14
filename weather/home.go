package weather

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type obj struct {
	GeneralSituation  string
	TcInfo            string
	FireDangerWarning string
	ForecastDesc      string
	Outlook           string
	UpdateTime        string
}

//Home is a homepage which show today's weather info
func Home(c *gin.Context) {
	var obj obj
	url := "https://data.weather.gov.hk/weatherAPI/opendata/weather.php?dataType=flw"
	url2 := "https://github.com/jacc-cyc"
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

	err3 := json.Unmarshal([]byte(data), &obj)
	if err3 != nil {
		log.Println(err3)
		os.Exit(1)
	}

	message := "Welcome to this weather app. Below is some weather info of today\n\n"
	message += "Weather Situation: " + obj.GeneralSituation
	message += "\n\nForecast Description: " + obj.ForecastDesc
	message += "\n\nOutlook: " + obj.Outlook
	message += "\n\nLast Update Time: " + obj.UpdateTime
	message += "\n\nPls visit " + url2 + " for more info about this app."

	c.String(http.StatusOK, message)
}
