package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	RedirectsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "url_shortener_redirects_total",
			Help: "Total number of redirects by shortcode.",
		},
		[]string{"shortcode"},
	)
	ShortenRequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "url_shortener_shorten_requests_total",
			Help: "Total number of shorten requests.",
		},
	)
)

func Init() {
	prometheus.MustRegister(RedirectsTotal)
	prometheus.MustRegister(ShortenRequestsTotal)
}

func Handler() http.Handler {
	return promhttp.Handler()
}
