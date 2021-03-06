package job

import (
	"fmt"

	"github.com/protoman92/mit-distributed-system/src/mapreduce/mrutil"
)

// WorkerJob represents a Map/Reduce job request.
type WorkerJob struct {
	File           string
	MapFuncName    mrutil.MapFuncName
	MapOpCount     uint
	ReduceFuncName mrutil.ReduceFuncName
	ReduceOpCount  uint

	// These properties can be changed.
	MapJobNumber   uint
	RemoteFileAddr string
	Type           mrutil.JobType
	Worker         string
}

// CheckWorkerJob checks the validity of a JobRequest.
func CheckWorkerJob(r WorkerJob) {
	if r.File == "" ||
		r.MapFuncName == "" ||
		r.MapOpCount == 0 ||
		r.ReduceFuncName == "" ||
		r.ReduceOpCount == 0 ||
		r.RemoteFileAddr == "" ||
		r.Type == mrutil.JobType(0) ||
		r.Worker == "" ||
		r.Worker == mrutil.UnassignedWorker {
		panic("Invalid parameters")
	}

	if r.Type == mrutil.Reduce {
		if r.RemoteFileAddr == mrutil.UnassignedWorker {
			panic("Invalid parameters")
		}
	}
}

// Clone clones the current args.
func (r WorkerJob) Clone() WorkerJob {
	return r
}

// Equals checks equality.
func (r WorkerJob) Equals(r1 WorkerJob) bool {
	return r.File == r1.File &&
		r.MapJobNumber == r1.MapJobNumber &&
		r.Type == r1.Type
}

// UID returns a unique ID for the current job request.
func (r WorkerJob) UID() string {
	return fmt.Sprintf("%s-M%d-%s", r.Type, r.MapJobNumber, r.File)
}
