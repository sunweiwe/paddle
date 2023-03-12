package orm

import (
	"context"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	ormPrometheus "gorm.io/plugin/prometheus"
)

type MySQLMetricsCollector struct {
	Prefix        string
	Interval      uint32
	VariableNames []string
	status        map[string]prometheus.Gauge
}

func (m *MySQLMetricsCollector) Metrics(p *ormPrometheus.Prometheus) []prometheus.Collector {
	if m.Prefix == "" {
		m.Prefix = "gorm_status_"
	}

	if m.Interval == 0 {
		m.Interval = p.RefreshInterval
	}

	if m.status == nil {
		m.status = map[string]prometheus.Gauge{}
	}

	go func() {
		for range time.Tick(time.Duration(m.Interval) * time.Second) {

		}
	}()

	collectors := make([]prometheus.Collector, 0, len(m.status))

	for _, v := range m.status {
		collectors = append(collectors, v)
	}

	return collectors
}

func (m *MySQLMetricsCollector) collect(p *ormPrometheus.Prometheus) {
	rows, err := p.DB.Raw("SHOW STATUS").Rows()

	if err != nil {
		p.DB.Logger.Error(context.Background(), "gorm:prometheus query error: %v", err)
		return
	}

	var name, value string
	for rows.Next() {
		err = rows.Scan(&name, &value)
		if err != nil {
			p.DB.Logger.Error(context.Background(), "gorm:prometheus scan error: %v", err)
			continue
		}

		var exist = len(m.VariableNames) == 0

		for _, n := range m.VariableNames {
			if n == name {
				exist = true
				break
			}
		}

		if exist {
			value, err := strconv.ParseFloat(value, 64)
			if err != nil {
				continue
			}

			gauge, ok := m.status[name]
			if !ok {
				gauge = prometheus.NewGauge(prometheus.GaugeOpts{
					Name:        m.Prefix + name,
					ConstLabels: p.Labels,
				})

				m.status[name] = gauge
				_ = prometheus.Register(gauge)

				gauge.Set(value)
			}
		}
	}

}
