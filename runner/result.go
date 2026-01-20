package runner

import "time"

type Result struct {
	Log        string
	TimeTaken  time.Duration
	MemoryUsed int
}
