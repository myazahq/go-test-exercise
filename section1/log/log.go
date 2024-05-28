package log

import (
	"fmt"
	"log"
	"os"
)

type Logger interface {
	Log(message string)
}

type fileLogger struct{
	file *os.File
}

func NewFileLogger(filePath string) *fileLogger {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
	return &fileLogger{
		file: file,
	}
}

func (fl *fileLogger) Log(message string){
	//write message to file
	fl.file.WriteString(message)
}

type ConsoleLogger struct{}

func (cl *ConsoleLogger) Log(message string){
	//write message to console
	fmt.Println(message)
}