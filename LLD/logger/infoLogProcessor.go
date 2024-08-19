package logger

import "fmt"

type InfoLogProcessor struct {
	NextProcessor Logger
}

func newInfoLogProcessor(nextLogProcessor Logger) InfoLogProcessor {
	return InfoLogProcessor{
		NextProcessor: nextLogProcessor,
	}
}

func (info InfoLogProcessor) Log(logLevel LOG_LEVEL, message string) {
	if logLevel == INFO {
		fmt.Println("info logger: ", message)
		return
	}

	info.NextProcessor.Log(logLevel, message)
}
