package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type CommandResponse struct {
	ModelName               string
	FirmwareNumber          string
	RatingVoltage           string
	RatingPowerWatt         string
	State                   string
	PowerSupplyBy           string
	UtilityVoltage          float64
	OutputVoltage           float64
	BatteryCapacityPercent  float64
	RemainingRuntimeSeconds float64
	LoadWatt                float64
	LoadPercent             float64
	LineInteraction         string
	TestResult              string
	LastPowerEvent          string
}

func parseFloat(s string) float64 {
	float, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Panicln("Unable to parse string to float64", s, err)
	}
	return float
}

func splitNumber(s string, sep string) float64 {
	split := strings.Split(s, sep)[0]

	return parseFloat(split)
}

func getDuration(s string) float64 {
	filtered := strings.ReplaceAll(s, ".", "")
	filtered = strings.ReplaceAll(filtered, " ", "")
	filtered = strings.ReplaceAll(filtered, "hour", "hr")
	filtered = strings.ReplaceAll(filtered, "min", "m")
	filtered = strings.ReplaceAll(filtered, "sec", "s")

	remainingRuntime, err := time.ParseDuration(filtered)
	if err != nil {
		fmt.Println("Warning: Unable to get remaining battery duration from", s)
	}
	return remainingRuntime.Seconds()
}

func ParseStatus(status PowerstatStatus) CommandResponse {
	ratingVoltage := strings.Split(status.RatingVoltage, " V")[0]
	ratingPower := strings.Split(status.RatingPower, " Watt")[0]

	utilityVoltage := splitNumber(status.UtilityVoltage, " V")
	outputVoltage := splitNumber(status.OutputVoltage, " V")
	batteryCapacity := splitNumber(status.BatteryCapacity, " %")
	remainingRuntimeSeconds := getDuration(status.RemainingRuntime)
	loadSplit := strings.Split(status.Load, " Watt(")
	loadWatt := parseFloat(loadSplit[0])
	loadPercent := splitNumber(loadSplit[1], " %)")

	return CommandResponse{
		ModelName:               status.ModelName,
		FirmwareNumber:          status.FirmwareNumber,
		RatingVoltage:           ratingVoltage,
		RatingPowerWatt:         ratingPower,
		State:                   status.State,
		PowerSupplyBy:           status.PowerSupplyBy,
		UtilityVoltage:          utilityVoltage,
		OutputVoltage:           outputVoltage,
		BatteryCapacityPercent:  batteryCapacity,
		RemainingRuntimeSeconds: remainingRuntimeSeconds,
		LoadWatt:                loadWatt,
		LoadPercent:             loadPercent,
		LineInteraction:         status.LineInteraction,
		TestResult:              status.TestResult,
		LastPowerEvent:          status.LastPowerEvent,
	}
}
