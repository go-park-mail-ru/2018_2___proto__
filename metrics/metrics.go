package metrics

import "github.com/prometheus/client_golang/prometheus"

type Profiler struct {
	// Hits        prometheus.Counter
	HitsStats   prometheus.CounterVec
	ActiveRooms prometheus.Gauge
	Routines    prometheus.Gauge
	CPUUsage    prometheus.Gauge
	MemUsage    prometheus.Gauge
	DiskUsage   prometheus.Gauge
}

func NewProfiler() *Profiler {
	p := Profiler{
		// Hits: prometheus.NewCounter(prometheus.CounterOpts{
		// 	Name: "hits_stat_uses_total",
		// 	Help: "Number of hits and statuses.",
		// }),

		HitsStats: *prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "hits_stat_uses_total",
			Help: "Number of hits and statuses.",
		},
			[]string{"status"}),

		ActiveRooms: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "active_rooms",
			Help: "Number of active rooms.",
		}),

		CPUUsage: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "cpu_usage",
			Help: "Info on CPU usage.",
		}),

		MemUsage: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "mem_usage",
			Help: "Info on mem usage.",
		}),

		DiskUsage: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "disk_usage",
			Help: "Info on disk usage.",
		}),
	}

	prometheus.MustRegister(p.HitsStats)
	prometheus.MustRegister(p.ActiveRooms)
	prometheus.MustRegister(p.CPUUsage)
	prometheus.MustRegister(p.MemUsage)
	prometheus.MustRegister(p.DiskUsage)

	return &p
}
