package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

type Config struct {
	Path  string
	Level string
}

//	GetLogger 获取日志 logger
func GetLogger() *zap.SugaredLogger {
	return Logger
}

// NewLogger 创建日志 logger
func NewLogger(c Config) (*zap.SugaredLogger, error) {
	cfg := zap.NewProductionConfig()
	err := makeDirIfNotExist(c.Path)
	if err != nil {
		return nil, err
	}

	if c.Path == "" {
		c.Path = c.Path + "/http_server.log"
	}
	cfg.OutputPaths = []string{
		"stdout",
		c.Path + "/http_server.log",
	}

	switch c.Level {
	case "INFO":
		cfg.Level.SetLevel(zap.InfoLevel)
	case "DEBUG":
		cfg.Level.SetLevel(zap.DebugLevel)
	case "WARNING":
		cfg.Level.SetLevel(zap.WarnLevel)
	case "ERROR":
		cfg.Level.SetLevel(zap.ErrorLevel)
	default:
		cfg.Level.SetLevel(zap.InfoLevel)
	}
	l, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	defer l.Sync()
	undo := zap.ReplaceGlobals(l)
	defer undo()

	Logger = zap.S()
	return Logger, err
}

func makeDirIfNotExist(path string) error {
	if _, err := os.Stat(path); err == nil {
		fmt.Println("path exists", path)
	} else {
		fmt.Println("path not exists ", path)
		err = os.MkdirAll(path, 0711)

		if err != nil {
			fmt.Println("Error creating directory")
			return err
		}
	}
	return nil
}
