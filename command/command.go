package command

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type CommandResponse struct {
	ModelName              string
	FirmwareNumber         string
	RatingVoltage          float64
	RatingPowerWatt        float64
	State                  string
	PowerSupplyBy          string
	UtilityVoltage         float64
	OutputVoltage          float64
	BatteryCapacityPercent float64
	RemainingRuntime       time.Duration
	LoadWatt               float64
	LoadPercent            float64
	LineInteraction        string
	TestResult             string
	LastPowerEvent         string
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

func Run() CommandResponse {
	status := pwrstatStatus()
	ratingVoltage := splitNumber(status.RatingVoltage, " V")
	ratingPower := splitNumber(status.RatingPower, " Watt")

	utilityVoltage := splitNumber(status.UtilityVoltage, " V")
	outputVoltage := splitNumber(status.OutputVoltage, " V")
	batteryCapacity := splitNumber(status.BatteryCapacity, " %")
	remainingRuntime, err := time.ParseDuration(status.RemainingRuntime)
	if err != nil {
		fmt.Println("Warning: Unable to get remaining battery duration from", status.RemainingRuntime)
	}
	loadSplit := strings.Split(status.Load, " Watt(")
	loadWatt := parseFloat(loadSplit[0])
	loadPercent := splitNumber(loadSplit[1], " %)")

	return CommandResponse{
		ModelName:              status.ModelName,
		FirmwareNumber:         status.FirmwareNumber,
		RatingVoltage:          ratingVoltage,
		RatingPowerWatt:        ratingPower,
		State:                  status.State,
		PowerSupplyBy:          status.PowerSupplyBy,
		UtilityVoltage:         utilityVoltage,
		OutputVoltage:          outputVoltage,
		BatteryCapacityPercent: batteryCapacity,
		RemainingRuntime:       remainingRuntime,
		LoadWatt:               loadWatt,
		LoadPercent:            loadPercent,
		LineInteraction:        status.LineInteraction,
		TestResult:             status.TestResult,
		LastPowerEvent:         status.LastPowerEvent,
	}
}
