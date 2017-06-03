package log

import (
	"os"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("twalk")

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func init() {
	regular := logging.NewLogBackend(os.Stderr, "", 0)
	formatted := logging.NewLogBackend(os.Stderr, "", 0)
	formattedFormatter := logging.NewBackendFormatter(formatted, format)
	regularLeveled := logging.AddModuleLevel(regular)
	regularLeveled.SetLevel(logging.ERROR, "")
	logging.SetBackend(regularLeveled, formattedFormatter)
}

func Debug(text string) {
	log.Debug(text)
}

func Info(text string) {
	log.Info(text)
}

func Notice(text string) {
	log.Notice(text)
}

func Warning(text string) {
	log.Warning(text)
}

func Error(text string) {
	log.Error(text)
}

func Critical(text string) {
	log.Critical(text)
}
