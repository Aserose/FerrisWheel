package dataSource

import (
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/imdario/mergo"
	"strings"
)

const (
	inputTime   = "setting the search time interval: "
	errorToUser = " incorrect format"
)

func (dS *dataSource) GetData(derive string) map[string][]interface{} {
	dS.logger.Info("dataSource: sending request to sources")
	result := dS.Vk.IncomingPointsman(derive)
	if len(result) > 0 {
		dS.Params["end_time"] = result["end_time"][len(result["end_time"])-1]
	} else {
		return nil
	}
	return result
}

func (dS *dataSource) SetRequestParametrs(derive string) (string, bool) {
	if !strings.Contains(dS.CheckStatus(), "already") {
		return "you are not authorized.\nEnter \"/start\" for authorization", false
	}

	if derive == "new request" {
		dS.Params = nil
		return "enter address", false
	}

	dS.defaultParams()

	status, ok := dS.validateRequestParams(derive)
	if !ok {
		return status, ok
	}

	dS.logger.Print(status)
	dS.Vk.SetParams(dS.Params)
	return status, ok
}

func (dS *dataSource) validateRequestParams(derive string) (string, bool) {
	queue := []string{"lat", "long", "end_time", "start_time"}

	for _, place := range queue {
		if dS.Params[place] == nil || dS.Params[place] == 0 {
			switch place {
			case "lat", "long":
				coordinates, status := dS.Geocode.GetCoordinates(derive)
				if status == "address not found" {
					return status, false
				}
				if err := mergo.Merge(&dS.Params, coordinates); err != nil {
					dS.logger.Errorf("dataSource: error occurred when setting up coordinates: %s", err.Error())
				}
				return fmt.Sprintf("address: %s\n\n%s\nenter the end date (mm/dd/yy)\n\nor send request", derive, inputTime), false
			case "end_time":
				if derive == "send request" {
					return " ", true
				}
				endTime, status := dS.convertTime(derive)
				if status == inputTime {
					return fmt.Sprint(status, errorToUser), false
				}
				if err := mergo.Merge(&dS.Params, map[string]interface{}{"end_time": endTime}); err != nil {
					dS.logger.Errorf("dataSource: error occurred when setting up end_time: %s", err.Error())
				}
				return fmt.Sprintf("end date: %s\n\n%s\nenter the start date (mm/dd/yy) \n\nor send request", derive, inputTime), false
			case "start_time":
				if derive == "send request" {
					return " ", true
				}
				startTime, status := dS.convertTime(derive)
				if status == inputTime {
					return fmt.Sprint(status, errorToUser), false
				}
				if err := mergo.Merge(&dS.Params, map[string]interface{}{"start_time": startTime}); err != nil {
					dS.logger.Errorf("dataSource: error occurred when setting up start_time: %s", err.Error())
				}
				return " ", true
			}
		}
	}

	return " ", true
}
func (dS *dataSource) convertTime(time string) (int64, string) {
	timeToUnix, err := dateparse.ParseAny(time)
	if err != nil {
		dS.logger.Errorf("dataSource: error occurred when setting up time: %s", err.Error())
		return 0, fmt.Sprint(inputTime)
	}

	return timeToUnix.Unix(), "failed"
}

func (dS *dataSource) defaultParams() {
	if dS.Params == nil {
		mergo.Merge(&dS.Params, map[string]interface{}{"start_time": 343547804})
		mergo.Merge(&dS.Params, map[string]interface{}{"count": 30})
		mergo.Merge(&dS.Params, map[string]interface{}{"radius": 100})
	} else {
		dS.Params["start_time"] = 343547804
		dS.Params["count"] = 30
		dS.Params["radius"] = 100
	}

}
