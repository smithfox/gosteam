package log

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Logger struct {
	logDir        string
	level         int
	rollingDay    int
	logfile       *os.File
	lastLogFile   *os.File
	lock          *sync.Mutex
	log           *log.Logger
	outputConsole bool
	name          string
}

const (
	LOG_DEBUG = iota
	LOG_INFO
	LOG_WARN
	LOG_ERROR
)

func NewLog(logDir string, logname string) *Logger {
	logger := &Logger{}
	logger.logDir = logDir
	logger.level = LOG_DEBUG
	logger.rollingDay = 0
	logger.lock = new(sync.Mutex)
	logger.name = logname
	return logger
}

func NewConsoleLog() *Logger {
	logger := &Logger{}
	logger.lock = new(sync.Mutex)
	logger.rollingDay = -1
	logger.log = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	return logger
}

//print log to console?
func (logger *Logger) ConsoleOutput(f bool) {
	logger.outputConsole = f
}

//LOG_DEBUG,LOG_INFO,LOG_WARN,LOG_ERROR
func (logger *Logger) SetLevel(level int) {
	logger.level = level
}

func (logger *Logger) Debug(v ...interface{}) {
	logger.rolling()
	logger.log.Println(v)
	if logger.outputConsole {
		fmt.Println(v)
	}
}

func (logger *Logger) Debugf(format string, v ...interface{}) {
	logger.rolling()
	logger.log.Printf(format, v...)
	if logger.outputConsole {
		fmt.Printf(format, v...)
	}
}

func (logger *Logger) Info(v ...interface{}) {
	if logger.level > LOG_INFO {
		return
	}
	logger.rolling()
	logger.log.Println(v)
	if logger.outputConsole {
		fmt.Println(v)
	}
}

func (logger *Logger) Infof(format string, v ...interface{}) {
	if logger.level > LOG_INFO {
		return
	}
	logger.rolling()
	logger.log.Printf(format, v...)
	if logger.outputConsole {
		fmt.Printf(format, v...)
	}
}

func (logger *Logger) Warn(v ...interface{}) {
	if logger.level > LOG_WARN {
		return
	}
	logger.rolling()
	logger.log.Println(v)
	if logger.outputConsole {
		fmt.Println(v)
	}
}

func (logger *Logger) Warnf(format string, v ...interface{}) {
	if logger.level > LOG_WARN {
		return
	}
	logger.rolling()
	logger.log.Printf(format, v...)
	if logger.outputConsole {
		fmt.Printf(format, v...)
	}
}

func (logger *Logger) Error(v ...interface{}) {
	if logger.level > LOG_ERROR {
		return
	}
	logger.rolling()
	logger.log.Println(v)
	if logger.outputConsole {
		fmt.Println(v)
	}
}

func (logger *Logger) Errorf(format string, v ...interface{}) {
	if logger.level > LOG_ERROR {
		return
	}
	logger.rolling()
	logger.log.Printf(format, v...)
	if logger.outputConsole {
		fmt.Printf(format, v...)
	}
}

func (logger *Logger) Close() {
	if logger.logfile != nil {
		logger.logfile.Close()
		logger.logfile = nil
	}
}

func (logger *Logger) rolling() {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	logger._rolling()
}

func (logger *Logger) _rolling() {
	now := time.Now()
	if logger.rollingDay < 0 {
		if logger.lastLogFile != nil {
			logger.lastLogFile.Close()
			logger.lastLogFile = nil
		}
		return
	} else {
		if logger.logfile != nil {
			logger.logfile.Sync()
		}
		//延迟关闭上一个log文件, 优化锁粒度: 无需锁文件写, 只需锁rolling本身
		logger.lastLogFile = logger.logfile
		logger.rollingDay = now.YearDay()

		logName := logger.name + "_" + now.Format("20060102") + ".log"
		logPath := filepath.Join(logger.logDir, logName)

		var err error
		logger.logfile, err = os.OpenFile(logPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			fmt.Printf("fail to create log file %v err=%v\n", logPath, err)
		}

		logger.log = log.New(logger.logfile, "", log.Ldate|log.Ltime)
	}

}
