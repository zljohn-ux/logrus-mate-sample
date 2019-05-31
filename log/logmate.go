package log

import (
	"github.com/gogap/logrus_mate"
	"github.com/sirupsen/logrus"

	_ "github.com/MrsJohn/logrus-mate-sample/log/linehook"
	_ "github.com/gogap/logrus_mate/hooks/expander"
	_ "github.com/gogap/logrus_mate/hooks/file"
	_ "github.com/gogap/logrus_mate/hooks/lfshook"
	_ "github.com/gogap/logrus_mate/writers/rotatelogs"
)

//InitLog 初始化日志
func InitLogursmate() {
	mate, _ := logrus_mate.NewLogrusMate(
		logrus_mate.ConfigFile(
			"./conf/mate.conf", // { mike {formatter.name = "text"} }
		),
	)
	if err := mate.Hijack(logrus.StandardLogger(),
		"debug",
	); err != nil {
		panic(err)
	}
}
