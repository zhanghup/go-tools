package main

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/dbi"
	"github.com/zhanghup/go-tools/database/influx"
	"math/rand"
	"time"
)

func main() {

	tags := make([]map[string]string, 0)
	for i := 0; i < 1; i++ {
		tags = append(tags, map[string]string{
			"tag": fmt.Sprintf("%03d", i),
			"id":  tools.UUID(),
		})
	}

	point := func() float64 {
		return rand.Float64() * 1000
	}

	//t := time.Now().AddDate(0, 0, -1)
	t := time.Now().Add(time.Second * -6)

	for t.Before(time.Now()) {
		for _, tag := range tags {
			yy, mm, dd := t.Date()
			hh, MM, ss := t.Clock()
			dbi.Write(influx.NewPoint("data", tag, map[string]interface{}{"price": point(), "v1": point(), "v2": point()}, time.Date(yy, mm, dd, hh, MM, ss, 0, time.Local)))
		}
		t = t.Add(time.Second)
	}

	time.Sleep(time.Second * 1)
	fmt.Println(dbi.Len(), "-------------")
	for {
		if dbi.Len() == 0 {
			break
		}
	}
}

func init() {
	dbi.InitDefault()
}
