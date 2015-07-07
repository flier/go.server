package base

import (
	"fmt"
	"log"
	"log/syslog"
	"os"

	"github.com/golang/glog"
)

type LogfFunc func(format string, v ...interface{})

type LoggingAdapter struct {
	Debugf, Infof, Warnf, Errorf, Fatalf, Panicf LogfFunc
}

var (
	LogAdapter = &LoggingAdapter{
		Debugf: func(format string, v ...interface{}) { log.Printf("DEBUG: "+format, v...) },
		Infof:  func(format string, v ...interface{}) { log.Printf("INFO: "+format, v...) },
		Warnf:  func(format string, v ...interface{}) { log.Printf("WARN: "+format, v...) },
		Errorf: func(format string, v ...interface{}) { log.Printf("ERROR: "+format, v...) },
		Fatalf: func(format string, v ...interface{}) { log.Fatalf("FATAL: "+format, v...) },
		Panicf: func(format string, v ...interface{}) { log.Panicf("FATAL: "+format, v...) },
	}

	GLogAdapter = &LoggingAdapter{
		Debugf: glog.V(2).Infof,
		Infof:  glog.V(1).Infof,
		Warnf:  glog.Warningf,
		Errorf: glog.Errorf,
		Fatalf: glog.Fatalf,
		Panicf: func(format string, v ...interface{}) {
			s := fmt.Sprintf(format, v...)

			glog.ErrorDepth(1, s)

			panic(s)
		},
	}
)

func newSyslogAdapter(w *syslog.Writer) *LoggingAdapter {
	return &LoggingAdapter{
		Debugf: func(format string, v ...interface{}) { w.Debug(fmt.Sprintf(format, v...)) },
		Infof:  func(format string, v ...interface{}) { w.Info(fmt.Sprintf(format, v...)) },
		Warnf:  func(format string, v ...interface{}) { w.Warning(fmt.Sprintf(format, v...)) },
		Errorf: func(format string, v ...interface{}) { w.Err(fmt.Sprintf(format, v...)) },
		Fatalf: func(format string, v ...interface{}) {
			s := fmt.Sprintf(format, v...)

			w.Crit(s)

			os.Exit(1)
		},
		Panicf: func(format string, v ...interface{}) {
			s := fmt.Sprintf(format, v...)

			w.Emerg(s)

			panic(s)
		},
	}
}
