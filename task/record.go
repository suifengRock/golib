package task

type ExecLog interface {
	TickerBegin(*TaskTicker)
	TickerEnd(*TaskTicker, error)
}

type execLog struct{}

func (e *execLog) TickerBegin(t *TaskTicker) {
	logger.Info(" 开始。。。第%d次执行。。。", t.GetExecCount())
}

func (e *execLog) TickerEnd(t *TaskTicker, err error) {
	logger.Info(" >>>>执行情况>>>>>>")
	if err != nil {
		logger.Error(" err : %s", err.Error())
	} else {
		logger.Info(" success")
	}
	logger.Info(" <<<<<<<<<<<<<<<<")
	logger.Info(" 结束。。。第%d次执行。。。", t.GetExecCount())

}
