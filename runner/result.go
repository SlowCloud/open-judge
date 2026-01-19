package runner

import "time"

type Result struct {
	finished   bool
	Log        []string
	TimeTaken  time.Duration
	MemoryUsed int64
}
