package main

import (
	"log"
	"log/slog"
	"lucasriboli/hrdwstatus/config"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	RootPath = "/sys/class/hwmon/"
)

func main() {
	logger := slog.Default()

	yamlData, err := os.ReadFile("./config.yaml")
	if err != nil {
		logger.Error(err.Error())
	}

	var configFile config.ConfigFile

	err = yaml.Unmarshal(yamlData, &configFile)
	if err != nil {
		logger.Error(err.Error())
	}

	c, err := os.ReadDir(RootPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range c {
		dat, err := os.ReadFile(RootPath + entry.Name() + "/name")

		if err != nil {
			logger.Error(err.Error())
		}

		var cpu = strings.TrimSuffix(string(dat), "\n")
		if cpu == string("k10temp") {
			temp, err := os.ReadFile(RootPath + entry.Name() + "/temp1_input")

			if err != nil {
				logger.Error(err.Error())
			}

			temp_int, err := strconv.Atoi(strings.TrimSuffix(string(temp), "\n"))

			if err != nil {
				logger.Error(err.Error())
			}

			final_val := temp_int / 1000

			if final_val <= configFile.Thresholds.Low {
				logger.Info("LOW TEMPERATURE " + strconv.Itoa(final_val) + "°C")
			} else if configFile.Thresholds.Low < final_val && final_val < configFile.Thresholds.High {
				logger.Info("OK TEMPERATURE " + strconv.Itoa(final_val) + "°C")
			} else if final_val >= configFile.Thresholds.High {
				logger.Info("HIGH TEMPERATURE " + strconv.Itoa(final_val) + "°C")
			}
		}
	}

}
