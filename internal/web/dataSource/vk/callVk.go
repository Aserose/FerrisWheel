package vk

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"strings"
	"time"
)

func (v *apiVk) Request() (map[string][]interface{}, error) {

	v.logger.Info("vk: sending request to VK")

	if err := v.Vk.RequestUnmarshal("photos.search", &v.Photo, v.Params); err != nil {
		v.logger.Errorf("vk: error receiving photos: %s", err.Error())
	}
	result := v.OutgoingPointsman()

	return result, nil
}

func (v *apiVk) IncomingPointsman(str string) map[string][]interface{} {
	v.logger.Info("vk: checking request")
	if str == "new request" {
		result, err := v.Request()
		if err != nil {
			v.logger.Errorf("vk: request sending error: %s", err.Error())
		}
		return result
	}
	return nil
}

func (v *apiVk) OutgoingPointsman() map[string][]interface{} {
	v.logger.Info("vk: sorting photo")
	result := map[string][]interface{}{}

	for _, a := range v.Photo.Items {
		for _, b := range a.Sizes {
			if b.Type == "z" {
				result["url"] = append(result["url"], b.URL)
				result["end_time"] = append(result["end_time"], a.Date)
				result["date"] = append(result["date"], strings.Trim(fmt.Sprintf("%s\n", time.Unix(int64(a.Date), 0)), "+0300 MSK\\x0a"))
				result["id"] = append(result["id"], func() string {
					if strings.Contains(fmt.Sprint(a.OwnerID), "-") {
						return strings.Replace(fmt.Sprintf(" vk.com/%d\n", a.OwnerID), "-", "public", 1)
					}
					return fmt.Sprintf(" vk.com/id%d\n", a.OwnerID)
				}())
			}
		}
	}

	return result
}

func (v *apiVk) SetParams(params map[string]interface{}) {
	v.Params = api.Params{
		"lat":        params["lat"],
		"long":       params["long"],
		"start_time": params["start_time"],
		"end_time":   params["end_time"],
		"count":      params["count"],
		"radius":     params["radius"],
	}
}
