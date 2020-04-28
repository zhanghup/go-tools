package tog

import (
	"github.com/zhanghup/go-tools/tog/logger"
)

var toginfo = logger.NewLogger(logger.OptionStdout())
var togerr = logger.NewLogger(logger.OptionStderr())
