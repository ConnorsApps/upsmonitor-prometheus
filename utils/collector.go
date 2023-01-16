package utils

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type MetricCollector struct {
	properties              *prometheus.GaugeVec
	state                   *prometheus.GaugeVec
	utilityVoltage          *prometheus.GaugeVec
	remainingRuntimeSeconds *prometheus.GaugeVec
	outputVoltage           *prometheus.GaugeVec
	batteryCapacityPercent  *prometheus.GaugeVec
	loadWatt                *prometheus.GaugeVec
	loadPercent             *prometheus.GaugeVec
}

func (collector *MetricCollector) Collect(ch chan<- prometheus.Metric) {
	start := time.Now()

	commandOutput := RunCommand()
	res := ParseStatus(commandOutput)

	collector.properties.WithLabelValues(res.ModelName, res.FirmwareNumber, res.RatingVoltage, res.RatingPowerWatt).Set(1)

	isNormal := res.State == "Normal"
	stateNumber := float64(1)
	if !isNormal {
		stateNumber = 0
	}

	collector.state.WithLabelValues(res.State, res.PowerSupplyBy, res.LineInteraction, res.TestResult, res.LastPowerEvent).Set(stateNumber)

	collector.utilityVoltage.WithLabelValues().Set(res.UtilityVoltage)
	collector.remainingRuntimeSeconds.WithLabelValues().Set(res.RemainingRuntimeSeconds)
	collector.outputVoltage.WithLabelValues().Set(res.OutputVoltage)
	collector.batteryCapacityPercent.WithLabelValues().Set(res.BatteryCapacityPercent)
	collector.loadWatt.WithLabelValues().Set(res.LoadWatt)
	collector.loadPercent.WithLabelValues().Set(res.LoadPercent)

	collector.properties.Collect(ch)
	collector.state.Collect(ch)
	collector.utilityVoltage.Collect(ch)
	collector.remainingRuntimeSeconds.Collect(ch)
	collector.outputVoltage.Collect(ch)
	collector.loadWatt.Collect(ch)
	collector.loadPercent.Collect(ch)

	elapsed := time.Since(start)
	fmt.Printf("Collected in %s\n", elapsed)
}

func (collector *MetricCollector) Describe(ch chan<- *prometheus.Desc) {
	collector.properties.Describe(ch)

}
func Collector() *MetricCollector {
	namespace := "hass_health"

	properties := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "properties",
		Help:      "properties",
	}, []string{"ModelName", "FirmwareNumber", "RatingVoltage", "RatingPowerWatt"})

	state := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "state",
		Help:      "state",
	}, []string{"State", "PowerSupplyBy", "LineInteraction", "TestResult", "LastPowerEvent"})

	utilityVoltage := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "utilityVoltage",
		Help:      "utilityVoltage",
	}, []string{})

	remainingRuntimeSeconds := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "remainingRuntimeSeconds",
		Help:      "remainingRuntimeSeconds",
	}, []string{})

	outputVoltage := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "outputVoltage",
		Help:      "outputVoltage",
	}, []string{})

	batteryCapacityPercent := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "batteryCapacityPercent",
		Help:      "batteryCapacityPercent",
	}, []string{})

	loadWatt := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "loadWatt",
		Help:      "loadWatt",
	}, []string{})

	loadPercent := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "loadPercent",
		Help:      "loadPercent",
	}, []string{})

	return &MetricCollector{
		properties,
		state,
		utilityVoltage,
		remainingRuntimeSeconds,
		outputVoltage,
		batteryCapacityPercent,
		loadWatt,
		loadPercent,
	}
}
