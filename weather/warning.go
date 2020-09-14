package weather

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

//WarningInfo is to show the latest weather warning info
func WarningInfo(c *gin.Context) {
	url := "https://data.weather.gov.hk/weatherAPI/opendata/weather.php?dataType=warningInfo"

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

	message := (gjson.Get(string(data), "details.#.contents")).String()
	if message == "" {
		message += "No weather warning is available now."
	}

	c.String(http.StatusOK, message)
}
