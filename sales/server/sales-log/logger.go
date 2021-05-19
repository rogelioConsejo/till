package sales_log

type Logger struct {
	persistence Persistence
}

func SalesLog(persistence Persistence) *Logger {
	logger := new(Logger)
	logger.persistence = persistence
	return logger
}

func (l *Logger) Log(sale *Sale) (id string, err error) {
	return l.persistence.Save(*sale)
}

func (l *Logger) Sales() ([]Sale, error) {
	return l.persistence.Sales()
}

func (l *Logger) Sale(id string) (Sale, error) {
	return l.persistence.Sale(id)
}