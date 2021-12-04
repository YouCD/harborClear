package main

import (
	config "harborClear/configs"
	"harborClear/pkg/core"
	"harborClear/pkg/log"
)

func main() {
	log.Debugf("env harborClear_Projects     value: %s", config.Projects)
	log.Debugf("env harborClear_KeepMonth    value: %d", config.Month)
	log.Debugf("env harborClear_Loglevel     value: %s", config.Loglevel)
	log.Debugf("env harborClear_HarborUrl    value: %s", config.HarborUrl)
	log.Debugf("env harborClear_HarborUser   value: %s", config.HarborUser)
	log.Debugf("env harborClear_UserPassword value: %s", config.HarborUserPassword)
	log.Debugf("env harborClear_ClearFlag    value: %v", config.ClearFlag)

	core.Core()
}
