package core

import (
	"bytes"
	"fmt"
	"go-blog/global"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

// 这是一个LogFormatter的累，需要实现Formatter接口，此接口需要有Format方法
type LogFormatter struct{}

// 配置日志 输出格式等
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	// 根据不用的日志级别设置颜色显示
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	// 获取日志的配置信息
	log := global.Config.Logger
	// 时间戳的格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	// 应该是显示函数的信息，在那个文件在那一行
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", log.Prefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s\n", log.Prefix, timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger() *logrus.Logger {
	mLog := logrus.New()
	// mLog.SetOutput(os.Stdout)
	// mLog.SetReportCaller(global.Config.Logger.ShowLine)
	// mLog.SetFormatter(&LogFormatter{})
	// leval, err := logrus.ParseLevel(global.Config.Logger.Level)
	// if err != nil {
	// 	leval = logrus.InfoLevel
	// }
	// mLog.SetLevel(leval)
	setLogger(mLog)
	global.Log = mLog
	return mLog
}

func setLogger(logs *logrus.Logger) {
	logs.SetOutput(os.Stdout)
	logs.SetReportCaller(global.Config.Logger.ShowLine)
	logs.SetFormatter(&LogFormatter{})
	leval, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		leval = logrus.InfoLevel
	}
	logs.SetLevel(leval)
}

func InitDefaultLogger() {
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(global.Config.Logger.ShowLine)
	logrus.SetFormatter(&LogFormatter{})
	leval, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		leval = logrus.InfoLevel
	}
	logrus.SetLevel(leval)
}
