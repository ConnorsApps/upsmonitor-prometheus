package main

import (
	"fmt"
	"net/http"

	"github.com/ConnorsApps/upsmonitor-prometheus/command"
)

func health(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("ok"))
}

func main() {
	// collector := utils.Collector()
	// prometheus.MustRegister(collector)

	status := command.Run()
	fmt.Println("status", status)

	// fmt.Println("Listening at localhost:" + utils.PORT + utils.METRICS_ENDPOINT)
	// http.Handle(utils.METRICS_ENDPOINT, promhttp.Handler())
	// http.Handle("/health", http.HandlerFunc(health))

	// log.Panicln(http.ListenAndServe(":"+utils.PORT, nil))
}
