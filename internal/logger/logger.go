package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logger interface {
	Info(message string)
	Error(message string)
}

type FileLogger struct {
	file *os.File
}

// NewFileLogger Создаем лог файл на текущий день, куда будем записывать
// всю информацию за день
func NewFileLogger() (*FileLogger, error) {
	today := time.DateOnly
	filePath := fmt.Sprintf("logs/%s.log", today)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &FileLogger{file: file}, nil
}

func (l *FileLogger) Info(message string) {
	log.SetOutput(l.file)
	log.Println(time.TimeOnly, " [INFO]: ", message)
}

func (l *FileLogger) Error(message string) {
	log.SetOutput(l.file)
	log.Println(time.TimeOnly, " [ERROR]: ", message)
}

func (l *FileLogger) Close() error {
	return l.file.Close()
}
