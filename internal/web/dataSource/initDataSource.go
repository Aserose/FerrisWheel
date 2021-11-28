package dataSource

import (
	"fmt"
	"strings"
)

func (d *dataSource) InitServices(tokens map[string]interface{}) {
	d.AuthorizationStatus = make(map[string]string)
	d.AuthorizationStatus["vk"] = d.Vk.InitVk(fmt.Sprintf("%v", tokens["vk"]))
}

func (d *dataSource) CheckStatus() string {
	statusList := "not authorized in: "
	for key, status := range d.AuthorizationStatus {
		if strings.Contains(status, "failed") {
			statusList += fmt.Sprintf("\n%s", key)
		}
	}
	if statusList == "not authorized in: " {
		return "you are already authorized"
	}
	return statusList
}
