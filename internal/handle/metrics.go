package handle

import (
	"github.com/rcrowley/go-metrics"
	"net/http"
)

// MetricsHandle writes metrics to JSON.
func MetricsHandle(w http.ResponseWriter, _ *http.Request) {
	metrics.WriteJSONOnce(metrics.DefaultRegistry, w)
}
