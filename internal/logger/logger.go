package logger

import (
	"log"
	"os"
	"path/filepath"
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
	today := time.Now().Format("2006-01-02")
	dir, _ := os.Getwd()
	logDir := filepath.Join(dir, "logs")
	filePath := filepath.Join(logDir, today+".log")
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &FileLogger{file: file}, nil
}

func (l *FileLogger) Info(message string) {
	log.SetOutput(l.file)
	log.Println("[INFO]: ", message)
}

func (l *FileLogger) Error(message string) {
	log.SetOutput(l.file)
	log.Println("[ERROR]: ", message)
}

func (l *FileLogger) Close() {
	log.SetOutput(l.file)
	log.Println("[INFO]: [SERVER] Server closed")
	_ = l.file.Close()
}
