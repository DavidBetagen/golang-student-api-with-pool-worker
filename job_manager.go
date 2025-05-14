package usecase

import (
	"sync"
	"time"
)

type JobStatus string

const (
	JobPending JobStatus = "pending"
	JobDone    JobStatus = "done"
	JobFailed  JobStatus = "failed"
)

type JobInfo struct {
	ID        string
	Status    JobStatus
	Error     string
	CreatedAt time.Time
}

type JobManager struct {
	jobs map[string]*JobInfo
	mu   sync.RWMutex
}

func NewJobManager() *JobManager {
	return &JobManager{
		jobs: make(map[string]*JobInfo),
	}
}

func (jm *JobManager) AddJob(id string) {
	jm.mu.Lock()
	defer jm.mu.Unlock()
	jm.jobs[id] = &JobInfo{
		ID:        id,
		Status:    JobPending,
		CreatedAt: time.Now(),
	}
}

func (jm *JobManager) UpdateStatus(id string, status JobStatus, errMsg string) {
	jm.mu.Lock()
	defer jm.mu.Unlock()
	if job, exists := jm.jobs[id]; exists {
		job.Status = status
		job.Error = errMsg
	}
}

func (jm *JobManager) GetJob(id string) (*JobInfo, bool) {
	jm.mu.RLock()
	defer jm.mu.RUnlock()
	job, exists := jm.jobs[id]
	return job, exists
}
