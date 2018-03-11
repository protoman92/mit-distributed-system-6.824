package worker

import (
	"fmt"
)

func (w *worker) String() string {
	return fmt.Sprintf("Worker %s", w.RPCParams.Address)
}

func (r *JobRequest) String() string {
	return fmt.Sprintf("Job request: %s", r.FilePath)
}
