package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	entity "main.go/files/entity"
)

func GetTokenByID(id int) string {
	logrus.Info(entity.CONST_API_COIN_TICKER + strconv.Itoa(id))
	response, err := http.Get(entity.CONST_API_COIN_TICKER + strconv.Itoa(id))
	if err != nil {
		logrus.Error("[crypto.getTokenByID] Error from GET: ", err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Error("[crypto.getTokenByID] Error reading response: ", err)
	}

	var arr []entity.CryptoResponse
	err = json.Unmarshal(responseData, &arr)
	if err != nil {
		logrus.Error("[crypto.getTokenByID] Error unmarshaling: ", err)
	}

	res := arr[0]

	data := []string{res.Symbol, res.Name, res.Price, res.PercentChangeDay, res.PercentChangeWeek, res.MarketCap}

	str := &strings.Builder{}
	table := tablewriter.NewWriter(str)
	table.SetHeader([]string{"Symbol", "Name", "Price", "24h", "7d", "Mkt Cap"})
	table.Append(data)
	table.Render()
	out := "```" + str.String() + "```"
	logrus.Info("size of table " + strconv.Itoa(len(out)))
	if len(out) > 2000 {
		logrus.Error("[crypto.getTokenByID] Error: length limit exceeded")
	}

	return out
}
