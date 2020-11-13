package logging

import (
	"fmt"
	"sync"
	"time"
)

const (
	LevelError = iota  // LevelError 错误
	LevelWarning    // LevelWarning 警告
	LevelInfo	// LevelInfo 提示
	LevelDebug	// LevelDebug 除错
	LevelPanic	//LevelPanic 极端错误
)


var (
	GloabLevel = LevelInfo

	levelPrefix = []string{"ERROR:  ","WARNING:","INFO:   ","DEBUG:  ","Panic:  "}
	mu    sync.Mutex
)

func  Println(level int, msg string) {

	mu.Lock()
	defer mu.Unlock()
	_, _ = fmt.Printf(
		"%s %s %s\n",
		levelPrefix[level],
		time.Now().Format("2006-01-02 15:04:05"),
		msg,
	)
}

// Panic 极端错误
func  Panic(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	Println(LevelPanic, msg)
	panic(msg)
}


// Error 错误
func Error(format string, v ...interface{}) {
	if LevelError > GloabLevel {
		return
	}
	msg := fmt.Sprintf(format, v...)
	Println(LevelError, msg)
}

// Warning 警告
func Warning(format string, v ...interface{}) {
	if LevelWarning > GloabLevel {
		return
	}
	msg := fmt.Sprintf(format, v...)
	Println(LevelWarning, msg)
}

// Info 信息
func Info(format string, v ...interface{}) {
	if LevelInfo >GloabLevel {
		return
	}
	msg := fmt.Sprintf(format, v...)
	Println(LevelInfo, msg)
}

// Debug 校验
func  Debug(format string, v ...interface{}) {
	if LevelDebug > GloabLevel {
		return
	}
	msg := fmt.Sprintf(format, v...)
	Println(LevelDebug, msg)
}

func SetLevel(level string) {
	intLevel := LevelInfo
	switch level {
	case "error":
		intLevel = LevelError
	case "warning":
		intLevel = LevelWarning
	case "info":
		intLevel = LevelInfo
	case "debug":
		intLevel = LevelDebug
	}
	GloabLevel = intLevel
}