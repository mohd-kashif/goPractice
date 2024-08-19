package logger

type Logger interface {
	Log(LOG_LEVEL, string)
}

type LogProcessor struct {
	NextProcessor Logger
}

func NewLogProcessor() LogProcessor {
	errorLogProcessor := newErrorLogProcessor(nil)
	debugLogProcessor := newDebugLogProcessor(&errorLogProcessor)
	infoLogger := newInfoLogProcessor(&debugLogProcessor)

	return LogProcessor{
		NextProcessor: infoLogger,
	}
}

func (log *LogProcessor) Log(logLevel LOG_LEVEL, message string) {
	if log.NextProcessor != nil {
		log.NextProcessor.Log(logLevel, message)
	}

}
