package core

import "go.uber.org/zap"

var logger, _ = zap.NewProduction()
var Log = logger.Sugar()
