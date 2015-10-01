package task

import (
	"golib/log"
)

var logger = New(os.Stderr, "[PLAN_TASK] ", true, log.Lshortfile|log.LstdFlags)
