package log

import (
	"context"
	"fmt"
	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gmimo/common/util"
	"gopkg.in/sohlich/elogrus.v7"
	"io"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var (
	myLog      *MyLogger
	defaultKey = "applicationId"
)

type LogConfig struct {
	FileName    string
	Level       string
	EsServer    string
	StashServer string
	Hooks       []string
	Outputs     []string
}

type MyLogger struct {
	log           *logrus.Logger
	*logrus.Entry // 用于自定义字段
}

func InitLogger(conf LogConfig, appname string) (err error) {
	logr := logrus.New()

	// 添加默认字段
	ent := logr.WithFields(logrus.Fields{defaultKey: appname}) // *logrus.Entry

	// 解析level
	lv, err := logrus.ParseLevel(conf.Level)
	if err != nil {
		return err
	}
	logr.SetLevel(lv)
	logr.SetFormatter(&logrus.JSONFormatter{})
	logr.SetNoLock() // 关闭logrus互斥锁

	// 多种日志输出
	if len(conf.Outputs) > 0 {
		logr.SetOutput(multiWriter(conf.Outputs, conf.FileName))
	}

	myLog = &MyLogger{logr, ent}

	// 添加多个Hook
	myLog.AddHooks(conf, appname)

	return nil
}

// ----------------------util------------------------

// 多种写入日志方法
func multiWriter(outputs []string, fn string) io.Writer {
	var (
		writers = make([]io.Writer, 0)
	)
	for _, wr := range outputs {

		switch strings.ToLower(wr) {
		case "stdout":
			writers = append(writers, os.Stdout)

		case "file":
			f, err := os.OpenFile(fn, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
			if err != nil {
				fmt.Println("日志文件写入失败:", err)
			}
			writers = append(writers, f)
		}
	}
	return io.MultiWriter(writers...)
}

// 添加多个Hook
func (l *MyLogger) AddHooks(conf LogConfig, appname string) {
	for _, t := range conf.Hooks {
		switch t {
		case "es":
			if hook := NewEsHook(conf.EsServer, strings.ToLower(appname), l.log.GetLevel()); hook != nil {
				l.Logger.AddHook(hook)
			}
		case "stash":
			if hook := NewLogstashHook(conf.StashServer); hook != nil {
				l.Logger.AddHook(hook)
			}
		}
	}
}

// 1. es hook
func NewEsHook(esUrl string, name string, level logrus.Level) *elogrus.ElasticHook {
	host := util.GetLocalIP()
	client, err := elastic.NewClient(elastic.SetURL(esUrl))
	if err != nil {
		fmt.Println("connect es server error:", err)
		return nil
	}

	// 根据日期，定义生成的ES索引名字
	indexFunc := func() string {
		return name + "-" + time.Now().Format("20060102")
	}

	hook, err := elogrus.NewAsyncElasticHookWithFunc(client, host, level, indexFunc)
	if err != nil {
		fmt.Println("add es hook error:", err)
		return nil
	}
	fmt.Println("已加载ES Hook:", esUrl, name, level)
	return hook
}

// 2. logstatsh hook
func NewLogstashHook(stashUrl string) *logrustash.Hook {
	hook, err := logrustash.NewHook("tcp", stashUrl, "")
	if err != nil {
		return nil
	}
	fmt.Println("已加载LogStash Hook:", stashUrl)
	return hook
}

// ----------------------api------------------------
// DEBUG
func Debug(args ...interface{}) {
	myLog.Debug(args...)
}
func Debugf(format string, args ...interface{}) {
	myLog.Debugf(format, args...)
}

func Debugc(c context.Context, args ...interface{}) {
	myLog.withRequestId(c).Debug(args...)
}
func Debugcf(c context.Context, format string, args ...interface{}) {
	myLog.withRequestId(c).Debugf(format, args...)
}

// INFO
func Info(args ...interface{}) {
	myLog.Info(args...)
}
func Infof(format string, args ...interface{}) {
	myLog.Infof(format, args...)
}

func Infoc(c context.Context, args ...interface{}) {
	myLog.withRequestId(c).Info(args...)
}
func Infocf(c context.Context, format string, args ...interface{}) {
	myLog.withRequestId(c).Infof(format, args...)
}

// WARN
func Warn(args ...interface{}) {
	myLog.Warn(args...)
}
func Warnf(format string, args ...interface{}) {
	myLog.Warnf(format, args...)
}

func Warnc(c context.Context, args ...interface{}) {
	myLog.withRequestId(c).Warn(args...)
}
func Warncf(c context.Context, format string, args ...interface{}) {
	myLog.withRequestId(c).Warnf(format, args...)
}

// ERROR
func Error(args ...interface{}) {
	myLog.withTrace(context.Background(), string(debug.Stack())).Error(args...)
}
func Errorf(format string, args ...interface{}) {
	myLog.withTrace(context.Background(), string(debug.Stack())).Errorf(format, args...)
}

func Errorc(c context.Context, args ...interface{}) {
	myLog.withTrace(c, string(debug.Stack())).Error(args...)
}
func Errorcf(c context.Context, format string, args ...interface{}) {
	myLog.withTrace(c, string(debug.Stack())).Errorf(format, args...)
}

// withRequestId: 用于打印context中的reqid值
// withTrace: 报错时打印报错堆栈信息

func (l *MyLogger) withRequestId(c context.Context) *logrus.Entry {
	return l.WithFields(logrus.Fields{
		"reqid": c.Value("reqid")})
}

func (l *MyLogger) withTrace(c context.Context, emsg string) *logrus.Entry {
	return l.WithFields(logrus.Fields{
		"reqid": c.Value("reqid"),
		"trace": emsg})
}
