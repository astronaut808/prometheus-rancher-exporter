package main

import (
	"regexp"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// Exporter Sets up all the runtime and metrics
type Exporter struct {
	labelsFilter   *regexp.Regexp
	rancherURL     string
	accessKey      string
	secretKey      string
	hideSys        bool
	resourceLimit  string
	mutex          sync.RWMutex
	gaugeVecs      map[string]*prometheus.GaugeVec
	cacheTTL       time.Duration
	cacheExpiredAt time.Time
	cache          map[string]*Data
}

// NewExporter creates the metrics we wish to monitor
func newExporter(rancherURL, accessKey, secretKey string, labelsFilter *regexp.Regexp, hideSys bool, resourceLimit string, cacheTTL time.Duration) *Exporter {
	gaugeVecs := addMetrics()
	return &Exporter{
		labelsFilter:  labelsFilter,
		gaugeVecs:     gaugeVecs,
		rancherURL:    rancherURL,
		accessKey:     accessKey,
		secretKey:     secretKey,
		hideSys:       hideSys,
		resourceLimit: resourceLimit,
		cacheTTL:      cacheTTL,
		cache:         make(map[string]*Data),
	}
}
