package main

import (
	"context"
	"sync"
)

// bully algorithm
// compare id
// pros simple
// cons may not yield the best leader

// raft

// lock-based lesder election

// to avoid long running jobs, we dispatch the work to another go routine
// two go routines communicate through channels
type Job interface {
	Run(context.Context)
}

type JobManager interface {
	StartJob()
	StopJob()
}

// yellow routine
type JMImpl struct {
	lock        sync.Mutex
	JobCancel   context.CancelFunc
	wg          sync.WaitGroup
	JobStopChan chan struct{}
}

// black rountine
type leaderElectionImpl struct{}

func (le *leaderElectionImpl) StartJob() {}

// ...

// heartbeat

// lock
// / who;s the leader
// / acquire/ release lock

// // acquire lock
// /
// intsert into le (resource name, leaser_od , expire_at)
// value ...
// on conflic
// do update set
// where
// returning

// // release
// //// update expire time to current time
// / heartbeat
// /// lease-based leader elction
// //// db upsert

func waitForCrtlC() {}
func main() {

}

/// things to consider
/// Hob panic
//// be aware users that runs job will overwrite your context

/// testing
/// using interface and mockJob for testing
/// for IO / network / stange hobs
