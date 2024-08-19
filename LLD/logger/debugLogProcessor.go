package logger

import "fmt"

type DebugLogProcessor struct {
	NextProcessor Logger
}

func newDebugLogProcessor(nextLogProcessor Logger) DebugLogProcessor {
	return DebugLogProcessor{
		NextProcessor: nextLogProcessor,
	}
}

func (debug *DebugLogProcessor) Log(logLevel LOG_LEVEL, message string) {
	if logLevel == DEBUG {
		fmt.Println("debug logger: ", message)
		return
	}

	debug.NextProcessor.Log(logLevel, message)
}
