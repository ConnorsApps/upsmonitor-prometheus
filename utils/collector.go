package utils

// type MetricCollector struct {
// 	properties              *prometheus.GaugeVec
// 	state                   *prometheus.GaugeVec
// 	utilityVoltage          *prometheus.GaugeVec
// 	outputVoltage           *prometheus.GaugeVec
// 	batteryCapacity         *prometheus.GaugeVec
// 	remainingRuntimeSeconds *prometheus.GaugeVec
// 	load                    *prometheus.GaugeVec
// 	lineInteraction         *prometheus.GaugeVec
// 	testResult              *prometheus.GaugeVec
// 	lastPowerEvent          *prometheus.GaugeVec
// }

// func (collector *MetricCollector) Collect(ch chan<- prometheus.Metric) {
// 	start := time.Now()
// 	elapsed := time.Since(start)

// 	// collector.running.WithLabelValues(state).Set(isRunning)

// 	collector.running.Collect(ch)

// 	fmt.Printf("Collected in %s\n", elapsed)
// }

// func (collector *MetricCollector) Describe(ch chan<- *prometheus.Desc) {
// 	collector.running.Describe(ch)

// }
// func Collector() *MetricCollector {
// 	namespace := "hass_health"

// 	running := prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 		Namespace: namespace,
// 		Name:      "running",
// 		Help:      "something",
// 	}, []string{"state"})

// 	return &MetricCollector{
// 		running,
// 	}
// }
