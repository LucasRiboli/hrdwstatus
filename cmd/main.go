package main

import (
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

const (
	RootPath = "/sys/class/hwmon/"
)

func main() {
	logger := slog.Default()

	c, err := os.ReadDir(RootPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range c {
		dat, _ := os.ReadFile(RootPath + entry.Name() + "/name")
		var cpu = strings.TrimSuffix(string(dat), "\n")
		if cpu == string("k10temp") {
			temp, _ := os.ReadFile(RootPath + entry.Name() + "/temp1_input")
			temp_int, _ := strconv.Atoi(strings.TrimSuffix(string(temp), "\n"))
			final_val := temp_int / 1000

			if final_val <= 25 {
				logger.Info("LOW TEMPERATURE " + strconv.Itoa(final_val) + "°C")
			} else if final_val < 60 {
				logger.Info("OK TEMPERATURE " + strconv.Itoa(final_val) + "°C")
			} else if final_val >= 60 {
				logger.Info("HIGH TEMPERATURE " + strconv.Itoa(final_val) + "°C")
			}
		}
	}

}
