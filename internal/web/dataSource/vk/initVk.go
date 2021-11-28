package vk

import (
	"errors"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
)

func (v *apiVk) InitVk(token string) string {
	var AuthorizationStatus string
	v.logger.Info("vk: initializing")
	v.Vk = api.NewVK(token)

	_, err := v.Vk.PhotosSearch(nil)
	if err != nil {
		v.logger.Info(errors.New(err.Error()))
		AuthorizationStatus = "initializing failed"
	} else {
		AuthorizationStatus = "initializing ok"
	}

	v.logger.Info("vk: status: ", AuthorizationStatus)
	return fmt.Sprint(AuthorizationStatus)
}
