package influx

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"time"
)

func NewPoint(measurement string, tags map[string]string, fields map[string]interface{}, ts time.Time) *write.Point {
	return influxdb2.NewPoint(measurement, tags, fields, ts)
}
