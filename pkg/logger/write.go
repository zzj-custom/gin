package logger

import (
	"fmt"
	"gopkg.in/errgo.v2/fmt/errors"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type LogWriter struct {
	logDir             string // 日志根目录地址
	module             string // 模块名
	currentFileName    string // 当前被指定的文件名
	currentUseFileName string // 当前使用的文件名
	turnCateDuration   time.Duration
	mutex              sync.RWMutex
	outFile            *os.File
}

func New(logPath string, module string, duration time.Duration) *LogWriter {
	return &LogWriter{
		logDir:           logPath,
		module:           module,
		turnCateDuration: duration,
	}
}

func (w *LogWriter) getName() string {
	base := time.Now().Truncate(w.turnCateDuration)
	return fmt.Sprintf(
		"%s/%s/%s_%s",
		w.logDir,
		base.Format("2006-01-02"),
		w.module,
		base.Format("15"),
	)
}

func (w *LogWriter) getWrite() (io.Writer, error) {
	fileName := w.currentFileName
	// 判断是否会有新的文件，会出现新的文件名
	useFileName := w.getName()
	if useFileName != fileName {
		fileName = useFileName
	}

	// 创建目录
	dirName := filepath.Dir(fileName)
	if err := os.MkdirAll(dirName, 0755); err != nil {
		return nil, errors.Wrap(err)
	}

	fileHandler, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	_ = w.outFile.Close()
	w.outFile = fileHandler
	w.currentUseFileName = fileName
	w.currentFileName = fileName
	return fileHandler, nil
}

func (w *LogWriter) Write(p []byte) (int, error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if out, err := w.getWrite(); err != nil {
		return 0, errors.New("failed to fetch target io.Writer")
	} else {
		return out.Write(p)
	}
}
