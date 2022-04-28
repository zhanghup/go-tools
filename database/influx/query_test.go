package influx_test

import (
	"fmt"
	"github.com/zhanghup/go-tools/database/influx"
	"testing"
)

func TestQuery(t *testing.T) {
	s := influx.InitEngine().Query("test").
		First().
		Columns("tag").
		Range1("-2h").
		Range2("-2h").
		Range3("1h", "2h").
		Limit1(1).
		Limit2(1).
		Limit3(1, 1).
		Filter("(r) => r.tag == '000'").
		Filter("(r) => r.tag == '000'", false).
		Find()
	fmt.Println(s)
}
