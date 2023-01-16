package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ConnorsApps/upsmonitor-prometheus/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func health(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("ok"))
}

func main() {
	collector := utils.Collector()
	prometheus.MustRegister(collector)

	fmt.Println("Listening at localhost:" + utils.PORT + utils.METRICS_ENDPOINT)
	http.Handle(utils.METRICS_ENDPOINT, promhttp.Handler())
	http.Handle("/health", http.HandlerFunc(health))

	log.Panicln(http.ListenAndServe(":"+utils.PORT, nil))
}
