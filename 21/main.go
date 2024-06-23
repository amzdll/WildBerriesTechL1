package main

import (
	"fmt"
	"log"
	"os"
)

// SomeService - the service that uses the adapter
type SomeService struct {
	logger loggerInterface
}

func NewService(logger loggerInterface) *SomeService {
	return &SomeService{
		logger: logger,
	}
}

func (s SomeService) printHello() {
	s.logger.Info("printHello()")
	fmt.Println("Hello World!")
}

// LoggerAdapter - adapter pattern implementation
type loggerInterface interface {
	Info(s string)
	Error(s string, err error)
	// ...
}

type LoggerAdapter struct {
	logger *log.Logger
}

func NewLoggerAdapter(logger *log.Logger) *LoggerAdapter {
	return &LoggerAdapter{
		logger: logger,
	}
}

func (l *LoggerAdapter) Info(s string) {
	l.logger.Println("INFO called func:", s)
}

func (l *LoggerAdapter) Error(s string, err error) {
	l.logger.Println("ERR called func:", s, err)
}

func main() {
	logger := log.New(os.Stdout, "log:\t", 0)
	loggerAdapter := NewLoggerAdapter(logger)
	service := NewService(loggerAdapter)
	service.printHello()
}
