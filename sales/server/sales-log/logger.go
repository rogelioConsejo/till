package sales_log

type Logger struct {
	persistence Persistence
}

func SalesLog(persistence Persistence) *Logger {
	logger := new(Logger)
	logger.persistence = persistence
	return logger
}

func (l *Logger) Log(sale *Sale) error {
	return l.persistence.Save(*sale)
}
