package linehook

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/gogap/config"
	"github.com/gogap/logrus_mate"
	"github.com/sirupsen/logrus"
)

type DefaultFieldsHook struct {
	Field  string
	Skip   int
	levels []logrus.Level
}

func (df *DefaultFieldsHook) Fire(entry *logrus.Entry) error {
	entry.Data[df.Field] = findCaller(df.Skip)
	return nil
}

func (df *DefaultFieldsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

type LineHookConfig struct {
	Address string
}

func init() {
	logrus_mate.RegisterHook("linehook", NewLineHook)
}

func NewLineHook(config config.Configuration) (hook logrus.Hook, err error) {
	hook = &DefaultFieldsHook{
		Field: "line",
		Skip:  5,
	}
	return
}

func findCaller(skip int) string {
	file := ""
	line := 0
	for i := 0; i < 10; i++ {
		file, line = getCaller(skip + i)
		if !strings.HasPrefix(file, "logrus") {
			break
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func getCaller(skip int) (string, int) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", 0
	}
	n := 0
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	return file, line
}
