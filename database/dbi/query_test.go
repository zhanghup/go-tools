package dbi_test

import (
	_ "embed"
	"github.com/zhanghup/go-tools/database/dbi"
	"testing"
)

//go:embed config-default.yml
var initConfigByte []byte

func TestQuery(t *testing.T) {
	dbi.Query().From("zander").
		Range1("-30h").
		Measurement("data", `_field=="price"`).
		Find()
}

func init() {
	dbi.InitDefault(initConfigByte)
}
