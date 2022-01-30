package covid

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
	"main.go/files/entity"
)

func GetCountryInfo(country string) string {
	response, err := http.Get("https://disease.sh/v3/covid-19/countries/" + country + "?yesterday=true&strict=true")
	if err != nil {
		logrus.Error("[covid.GetCountryInfo] Error from GET: ", err)
		return "Invalid country"
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Error("[covid.GetCountryInfo] Error reading response: ", err)
	}

	var resp entity.CovidResponse
	err = json.Unmarshal(responseData, &resp)
	if err != nil {
		logrus.Error("[covid.GetCountryInfo] Error unmarshaling: ", err)
	}

	result := `Covid Information for ` + country + ` as of yesterday: 
	
Total cases			: ` + strconv.Itoa(resp.Cases) + `
New cases today		: ` + strconv.Itoa(resp.TodayCases) + `

Total deaths		: ` + strconv.Itoa(resp.Deaths) + `
New deaths today	: ` + strconv.Itoa(resp.TodayDeaths) + `

Total recovered		: ` + strconv.Itoa(resp.Recovered) + `
New recoveries today: ` + strconv.Itoa(resp.TodayRecovered) + `

Active cases		: ` + strconv.Itoa(resp.Active) + `
Critical			: ` + strconv.Itoa(resp.Critical) + `

Data provided via disease.sh API. Sourced from Worldometers.`

	return result
}
