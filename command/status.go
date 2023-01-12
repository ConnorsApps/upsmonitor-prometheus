package command

import (
	"os"
	"strings"
)

type PowerstatStatus struct {
	ModelName        string
	FirmwareNumber   string
	RatingVoltage    string
	RatingPower      string
	State            string
	PowerSupplyBy    string
	UtilityVoltage   string
	OutputVoltage    string
	BatteryCapacity  string
	RemainingRuntime string
	Load             string
	LineInteraction  string
	TestResult       string
	LastPowerEvent   string
}

// Model Name................... ST Series
func labelValue(output string, label string) string {
	afterLabel := strings.Split(output, label)[1]
	beforeNewline := strings.Split(afterLabel, "\n")[0]
	dotsRemoved := strings.Split(beforeNewline, ".... ")[1]
	return dotsRemoved
}

func command() string {
	// out, err := exec.Command("pwrstat", "-status").Output()

	// if err != nil {
	// 	log.Panicln("Unable to run command", err)
	// }
	out, err := os.ReadFile("./example/command-output")
	if err != nil {
		panic(err)
	}
	return string(out)
}

func pwrstatStatus() PowerstatStatus {
	output := command()

	split := strings.Split(output, "Current UPS status")

	propertiesText := split[0]
	upsStatus := split[1]

	model := labelValue(propertiesText, "Model Name")
	firmware := labelValue(propertiesText, "Firmware Number")
	ratingVoltage := labelValue(propertiesText, "Voltage")
	ratingPower := labelValue(propertiesText, "Rating Power")

	state := labelValue(upsStatus, "State")
	powerSupplyBy := labelValue(upsStatus, "Power Supply by")
	utilityVoltage := labelValue(upsStatus, "Utility Voltage")
	outputVoltage := labelValue(upsStatus, "Output Voltage")
	batteryCapacity := labelValue(upsStatus, "Battery Capacity")
	remainingRuntime := labelValue(upsStatus, "Remaining Runtime")
	load := labelValue(upsStatus, "Load")
	lineInteraction := labelValue(upsStatus, "Line Interaction")
	testResult := labelValue(upsStatus, "Test Result")
	lastPowerEvent := labelValue(upsStatus, "Last Power Event")

	return PowerstatStatus{
		ModelName:        model,
		FirmwareNumber:   firmware,
		RatingVoltage:    ratingVoltage,
		RatingPower:      ratingPower,
		State:            state,
		PowerSupplyBy:    powerSupplyBy,
		UtilityVoltage:   utilityVoltage,
		OutputVoltage:    outputVoltage,
		BatteryCapacity:  batteryCapacity,
		RemainingRuntime: remainingRuntime,
		Load:             load,
		LineInteraction:  lineInteraction,
		TestResult:       testResult,
		LastPowerEvent:   lastPowerEvent,
	}
}
