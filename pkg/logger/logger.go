package logger

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	//set log  logout by json
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetReportCaller(false)
}

func Write(msg string, filename string) {
	setOutPutFile(logrus.InfoLevel, filename)
	logrus.Info(msg)
}

func Debug(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.DebugLevel, "debug")
	logrus.WithFields(fields).Debug(args)
}

func Info(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.DebugLevel, "info")
	logrus.WithFields(fields).Info(args)
}

func Warn(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.DebugLevel, "warn")
	logrus.WithFields(fields).Warn(args)
}

func Fatal(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.DebugLevel, "fatal")
	logrus.WithFields(fields).Fatal(args)
}

func Error(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.DebugLevel, "error")
	logrus.WithFields(fields).Error(args)
}

func Panic(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.DebugLevel, "panic")
	logrus.WithFields(fields).Trace(args)
}

func Trace(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.DebugLevel, "trace")
	logrus.WithFields(fields).Trace(args)
}

func setOutPutFile(level logrus.Level, logName string) {
	if _, err := os.Stat("./runtime/log"); os.IsNotExist(err) {
		err := os.Mkdir("./runtime/log", 0777)
		if err != nil {
			panic(fmt.Errorf("create log dir failed, err: %v", err))
		}
	}

	timeStr := time.Now().Format("2006-01-02")
	fileName := path.Join("./runtime/log", logName+"_"+timeStr+".log")

	var err error
	os.Stderr, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
	}

	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(level)
	return
}

func LoggerToFile() gin.LoggerConfig {

	if _, err := os.Stat("./runtime/log"); os.IsNotExist(err) {
		err := os.Mkdir("./runtime/log", 0777)
		if err != nil {
			panic(fmt.Errorf("create log dir failed, err: %v", err))
		}
	}

	timeStr := time.Now().Format("2006-01-02")
	fileName := path.Join("./runtime/log", "gin_"+timeStr+".log")

	os.Stderr, _ = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	var conf = gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("%s -%s \"%s %s %s %d %s \"%s\" %s\"\n",
				param.TimeStamp.Format("2006-01-02 15:04:05"),
				param.ClientIP,
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		},
		Output: io.MultiWriter(os.Stdout, os.Stderr),
	}

	return conf
}

func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			if _, errDir := os.Stat("./runtime/log"); os.IsNotExist(errDir) {
				err := os.Mkdir("./runtime/log", 0777)
				if err != nil {
					panic(fmt.Errorf("create log dir failed, err: %v", err))
				}
			}

			timeStr := time.Now().Format("2006-01-02")
			fileName := path.Join("./runtime/log", "error_"+timeStr+".log")

			f, errFile := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			if errFile != nil {
				panic(fmt.Errorf("create log file failed, err: %v", errFile))
			}

			timeFileStr := time.Now().Format("2006-01-02 15:04:05")
			f.WriteString("panic error time:" + timeFileStr + "\n")
			f.WriteString(fmt.Sprintf("%v", err))
			f.WriteString("Stacktrce from panic:" + string(debug.Stack()) + "\n")
			f.Close()
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  fmt.Sprintf("%v", err),
			})
			c.Abort()

		}
	}()
	c.Next()
}
