package logger

import "fmt"

type ErrorLogProcessor struct {
	NextProcessor Logger
}

func newErrorLogProcessor(nextLogProcessor Logger) ErrorLogProcessor {
	return ErrorLogProcessor{
		NextProcessor: nextLogProcessor,
	}
}

func (error *ErrorLogProcessor) Log(logLevel LOG_LEVEL, message string) {
	if logLevel == ERROR {
		fmt.Println("error logger: ", message)
		return
	}
	if error.NextProcessor == nil {
		fmt.Println("undefined log level: ", message)
	}
}
