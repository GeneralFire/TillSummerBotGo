package inmemcache

import (
	"strconv"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	deleteOpLabel = "delete"
	putOpLabel    = "put"
	getOpLabel    = "get"
	popOpLabel    = "pop"
)

type meterecCache[K comparable, V any] struct {
	cache cache[K, V]

	cacheRequestsTotal    *prometheus.CounterVec
	HitMissRate           *prometheus.CounterVec
	HistogramResponseTime *prometheus.HistogramVec
}

func NewMeteredCache[K comparable, V any](
	TTL time.Duration,
	cacheName string,
) CacheInterface[K, V] {
	var useTTL bool

	if TTL == 0 {
		useTTL = false
	} else {
		useTTL = true
	}

	return &meterecCache[K, V]{
		cache: cache[K, V]{
			useTTL:  useTTL,
			storage: make(map[K]valueWithCreationTime[V]),
			lock:    sync.RWMutex{},
			TTL:     TTL,
		},

		cacheRequestsTotal: promauto.NewCounterVec(prometheus.CounterOpts{
			Namespace: cacheName,
			Subsystem: "cache",
			Name:      "requests_total",
		},
			[]string{"request"},
		),

		HitMissRate: promauto.NewCounterVec(prometheus.CounterOpts{
			Namespace: cacheName,
			Subsystem: "cache",
			Name:      "hit_miss_rate",
		},
			[]string{"request", "status"},
		),

		HistogramResponseTime: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: cacheName,
			Subsystem: "cache",
			Name:      "histogram_response_time_seconds",
			Buckets:   prometheus.ExponentialBuckets(0.0000001, 2, 16),
		},
			[]string{"request", "status"},
		),
	}
}

func (c *meterecCache[K, V]) Put(key K, value V) bool {
	c.cacheRequestsTotal.WithLabelValues(putOpLabel).Inc()

	timeStart := time.Now()

	ok := c.cache.Put(key, value)

	elapsed := time.Since(timeStart).Seconds()

	statusAsString := strconv.FormatBool(ok)
	c.HistogramResponseTime.WithLabelValues(putOpLabel, statusAsString).Observe(elapsed)
	c.HitMissRate.WithLabelValues(putOpLabel, statusAsString).Inc()
	return ok
}

func (c *meterecCache[K, V]) Get(key K) (V, bool) {
	c.cacheRequestsTotal.WithLabelValues(getOpLabel).Inc()

	timeStart := time.Now()

	v, ok := c.cache.Get(key)

	elapsed := time.Since(timeStart).Seconds()

	statusAsString := strconv.FormatBool(ok)
	c.HistogramResponseTime.WithLabelValues(getOpLabel, statusAsString).Observe(elapsed)
	c.HitMissRate.WithLabelValues(getOpLabel, statusAsString).Inc()
	return v, ok
}

func (c *meterecCache[K, V]) Pop(key K) (V, bool) {
	c.cacheRequestsTotal.WithLabelValues(popOpLabel).Inc()

	timeStart := time.Now()

	v, ok := c.cache.Pop(key)

	elapsed := time.Since(timeStart).Seconds()

	statusAsString := strconv.FormatBool(ok)
	c.HistogramResponseTime.WithLabelValues(popOpLabel, statusAsString).Observe(elapsed)
	c.HitMissRate.WithLabelValues(popOpLabel, statusAsString).Inc()
	return v, ok
}

func (c *meterecCache[K, V]) Delete(key K) {
	c.cacheRequestsTotal.WithLabelValues(deleteOpLabel).Inc()

	timeStart := time.Now()

	c.cache.Delete(key)

	elapsed := time.Since(timeStart).Seconds()

	statusAsString := strconv.FormatBool(true)
	c.HistogramResponseTime.WithLabelValues(deleteOpLabel, statusAsString).Observe(elapsed)
	c.HitMissRate.WithLabelValues(deleteOpLabel, statusAsString).Inc()
}
