package core

import (
	"fmt"

	"go.uber.org/zap"
)

var logger, _ = zap.NewProduction()

var Log = OpenBankingLogger{}

type OpenBankingLogger struct {
	*zap.SugaredLogger
}

func (obl *OpenBankingLogger) Info(args ...any) {
	obl.Infoln(args)
}

func (obl *OpenBankingLogger) Warn(args ...any) {
	obl.Warnln(args)
}

func (obl *OpenBankingLogger) Error(args ...any) {
	obl.Errorln(args)
}

func (obl *OpenBankingLogger) Sync() {
	if err := obl.SugaredLogger.Sync(); err != nil {
		fmt.Println(err.Error())
	}
}
