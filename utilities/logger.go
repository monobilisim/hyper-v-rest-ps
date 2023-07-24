package utilities

import (
	"runtime"
	"strconv"
	"strings"

	walog "github.com/sirupsen/logrus"
)

var Log = walog.New()

var log = Log

func SetupLogger() {

	log.SetFormatter(&walog.TextFormatter{
		ForceColors:     true, // Enable colors in the console output
		FullTimestamp:   true, // Show full timestamp with date and time
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			// Get the filename and the function name from the file path
			slash := strings.LastIndex(f.File, "/")
			filename := f.File[slash+1:]
			return "", "[" + filename + ":" + strconv.Itoa(f.Line) + "]"
		},
	})

	log.SetReportCaller(true)
	log.SetLevel(walog.TraceLevel)
}
