package mylog

import (
	"fmt"
	seelog "github.com/cihub/seelog"
)

func InitLog(conf_file string) (err error) {
	logger, err := seelog.LoggerFromConfigAsFile(conf_file)

	if err != nil {
		seelog.Critical("err parsing config log file:", err)
		return
	}
	err = seelog.ReplaceLogger(logger)
	if err != nil {
		seelog.Critical("err ReplaceLogger config log file:", err)
		return
	}
	return nil
}

func LOG_DEBUG(v ...interface{}) {
	seelog.Debug(v...)
}

func LOG_INFO(v ...interface{}) {
	seelog.Info(v...)
}

func LOG_ERROR(v ...interface{}) {
	seelog.Error(v...)
}

func LOG_DEBUG_F(format string, params ...interface{}) {
	seelog.Debugf(format, params...)
}

func LOG_INFO_F(format string, params ...interface{}) {
	seelog.Infof(format, params...)
}

func LOG_ERROR_F(format string, params ...interface{}) {
	seelog.Errorf(format, params...)
}

func Flush() {
	seelog.Flush()
}

func InitLogInfo() {
	fmt.Println("InitLogInfo!")
}
