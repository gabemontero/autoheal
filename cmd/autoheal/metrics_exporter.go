package main

import (
	"net/http"

	monitoring "github.com/openshift/autoheal/pkg/apis/monitoring/v1alpha1"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	actionsInitiated = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "autoheal_actions_initiated_total",
			Help: "Number of initiated healing actions",
		},
		[]string{"type", "template", "rule"},
	)
)

func (h *Healer) metricsHandler() http.Handler {
	return promhttp.Handler()
}

func (h *Healer) initExportedMetrics() {
	prometheus.MustRegister(actionsInitiated)
}

func (h *Healer) incrementAwxActions(
	action *monitoring.AWXJobAction,
	ruleName string,
) {
	actionsInitiated.With(
		map[string]string{
			"type":     "awxJob",
			"template": action.Template,
			"rule":     ruleName,
		},
	).Inc()
}
