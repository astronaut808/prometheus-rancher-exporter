package main

import (
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func (e *Exporter) IsCacheExpired() bool {
	if e.cacheTTL == 0 {
		return true
	}
	if e.cacheExpiredAt.IsZero() {
		return true
	}
	return e.cacheExpiredAt.Before(time.Now())
}

func (e *Exporter) RenewCache() {
	e.cacheExpiredAt = time.Now().Add(e.cacheTTL)
}

// Resets the guageVecs back to 0
// Ensures we start from a clean sheet
func (e *Exporter) resetGaugeVecs() {
	for _, m := range e.gaugeVecs {
		m.Reset()
	}
}

// Describe describes all the metrics ever exported by the Rancher exporter
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range e.gaugeVecs {
		m.Describe(ch)
	}
}

// Collect function, called on by Prometheus Client library
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	var (
		data *Data
		err  error
	)

	e.mutex.Lock() // To protect metrics from concurrent collects.
	defer e.mutex.Unlock()

	e.resetGaugeVecs() // Clean starting point

	var endpointOfAPI []string
	if strings.HasSuffix(rancherURL, "v3") || strings.HasSuffix(rancherURL, "v3/") {
		endpointOfAPI = endpointsV3
	} else {
		endpointOfAPI = endpoints
	}

	cacheExpired := e.IsCacheExpired()

	// Range over the pre-configured endpoints array
	for _, p := range endpointOfAPI {
		if cacheExpired {
			data, err = e.gatherData(e.rancherURL, e.resourceLimit, e.accessKey, e.secretKey, p, ch)
			if err != nil {
				log.Errorf("Error getting JSON from URL %s", p)
				return
			}
			e.cache[p] = data
		} else {
			d, ok := e.cache[p]
			if !ok {
				continue
			}
			data = d
		}

		if err := e.processMetrics(data, p, e.hideSys, ch); err != nil {
			log.Errorf("Error scraping rancher url: %s", err)
			return
		}
		log.Infof("Metrics successfully processed for %s", p)
	}

	if cacheExpired {
		e.RenewCache()
	}

	for _, m := range e.gaugeVecs {
		m.Collect(ch)
	}
}
