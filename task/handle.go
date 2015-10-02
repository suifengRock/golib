package task

import (
	"errors"
	"time"
)

var taskHandle = make([]*TaskTicker, 0)

var taskChannel = make(chan string)

func TaskHandle(taskIndex string, f logicalFunc, after, interval time.Duration, allowCount uint64, log *ExecLog) {
	ticker := NewTaskTicker(f, after, interval, allowCount, log)
	taskHandle[taskIndex] = ticker
	taskChannel <- taskIndex
}

func PlanTaskExec() {
	for {
		select {
		case index, ok := <-taskChannel:
			if ok {
				go taskHandle[index].TaskFire()
			}
		}
	}
}

func TaskStop(index string) (bool, error) {
	ticker, ok := taskHandle[index]
	if !ok {
		logger.Error("  has no task for this index:%s ", index)
		return false, errors.New(" has no task for this index")
	}
	ticker.taskStop()
	delete(taskHandle, index)
	return true, nil
}

// execTimestr:  格式为"15:04:05", 如果该时间当天已过，则在明天该时间执行
// intervalSesond: 为0 表示只执行一次，allowCount无效
func TaskHandleFmt(taskindex string, f logicalFunc, execTimeStr string, intervalSecond uint, allowCount uint64, log *execLog) {
	nowInterval := myTime.GetNowInterval(execTimeStr)
	after := nowInterval * time.Second
	interval := intervalSecond * time.Second
	TaskHandle(taskIndex, f, after, interval, allowCount, log)
}
