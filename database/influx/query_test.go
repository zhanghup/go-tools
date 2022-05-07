package influx_test

import (
	"github.com/zhanghup/go-tools/database/influx"
	"testing"
)

func TestQuery(t *testing.T) {
	influx.InitDefault()

	influx.Query("zander").
		Range1("-30m").
		Measurement("data").
		FilterEqual("tag", "000").
		FilterEqual("_field", "v2").
		Find()
}
