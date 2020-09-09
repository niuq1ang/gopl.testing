package license

const (
	BATCH_DEFAULT_MAX = 100
)

type LogBatch struct {
	Max    int
	buffer []LogReporter
}

func InitLogBatch() *LogBatch {
	log := &LogBatch{Max: BATCH_DEFAULT_MAX}
	log.buffer = make([]LogReporter, 0, BATCH_DEFAULT_MAX)
	return log
}

func (log *LogBatch) AddLog2LogBatch(data LogReporter) error {
	log.buffer = append(log.buffer, data)
	if len(log.buffer) >= log.Max {
		return log.Flush()
	}
	return nil
}

func (log *LogBatch) Flush() error {
	if len(log.buffer) == 0 {
		return nil
	}
	err := WriteLog2File(log.buffer)
	if err != nil {
		return err
	}
	log.buffer = log.buffer[:0]
	return nil
}

func (log *LogBatch) Close() error {
	return log.Flush()
}
