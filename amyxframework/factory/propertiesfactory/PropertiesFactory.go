package propertiesfactory

import (
	"../../environment"
	"../../utils/httputils"
	"../../utils/jsonutils"
	"../../utils/logger"
	"../../utils/stringutils"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func BuildProperties(propertiesFile string) {
	logger.Info("Load Properties: %s", propertiesFile)
	//从配置文件获取配置项
	configs := loadPropertieFromFile(propertiesFile)
	for key, value := range configs {
		environment.SetProperties(key, value)
	}
	//如果启用配置中心，则继续从配置中心获取配置项
	if environment.GetPropertyAsBool("config.center.enable") == true {
		monitorConfigCenter()
	}
}

func loadPropertieFromFile(path string) map[string]string {
	config := make(map[string]string)

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
	return config
}

func monitorConfigCenter() {
	logger.Info("Load Config Center")
	applicationName := environment.GetProperty("server.name")
	logger.Info("Application Name: %s", applicationName)
	configCenterUrlValue := environment.GetProperty("config.center.url")
	configCenterProfileValue := environment.GetProperty("config.center.profile")
	configCenterLabelValue := environment.GetProperty("config.center.label")
	if stringutils.IsEmpty(configCenterUrlValue) ||
		stringutils.IsEmpty(configCenterProfileValue) ||
		stringutils.IsEmpty(configCenterLabelValue) {
		logger.Error("Invalid Config Center Property")
		return
	}
	logger.Info("Config Center: %s", configCenterUrlValue, configCenterProfileValue, configCenterLabelValue)
	configs := loadPropertieFromConfigCenter(configCenterUrlValue, "public", configCenterProfileValue, configCenterLabelValue)
	if configs != nil {
		for key, value := range configs {
			environment.SetProperties(key, value)
		}
	}
	configs = loadPropertieFromConfigCenter(configCenterUrlValue, applicationName, configCenterProfileValue, configCenterLabelValue)
	if configs != nil {
		for key, value := range configs {
			environment.SetProperties(key, value)
		}
	}
}

type ConfigCenterResult struct {
	Data map[string]string
}

func loadPropertieFromConfigCenter(url string, application string, profile string, label string) map[string]string {
	fullUrl := fmt.Sprintf("%s/configs?profile=%s&application=%s&label=%s", url, profile, application, label)
	result, err := httputils.Get(fullUrl, nil)
	if err != nil {
		logger.Error("Load Config Center Failed: %v", err)
		return nil
	}
	logger.Info("Read Configuration: %s", result)
	var configCenterResult ConfigCenterResult
	jsonutils.Json2Obj(result, &configCenterResult)
	return configCenterResult.Data
}
