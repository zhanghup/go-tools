package main

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/influx"
	"math/rand"
	"time"
)

func main() {

	tags := make([]map[string]string, 0)
	for i := 0; i < 1000; i++ {
		tags = append(tags, map[string]string{
			"tag": fmt.Sprintf("%03d", i),
			"id":  tools.UUID(),
		})
	}

	point := func() float64 {
		return rand.Float64() * 1000
	}

	t := time.Now().AddDate(-1, 0, 0)

	e := influx.InitEngine()

	for t.Before(time.Now()) {
		for _, tag := range tags {
			e.Write(influx.NewPoint("data", tag, map[string]interface{}{"price": point()}, t))
		}
		t.Add(time.Second)
	}
}
