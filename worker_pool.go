package usecase

import (
	"log"
	"student-api/internal/student/domain"

	"github.com/panjf2000/ants/v2"
)

type WorkerPool struct {
	pool *ants.Pool
	repo domain.StudentRepository
}

func NewWorkerPool(repo domain.StudentRepository, maxWorkers int) *WorkerPool {
	p, err := ants.NewPool(maxWorkers)
	if err != nil {
		log.Fatalf("Failed to create worker pool: %v", err)
	}
	return &WorkerPool{
		pool: p,
		repo: repo,
	}
}

func (wp *WorkerPool) SubmitStudent(student *domain.Student) {
	err := wp.pool.Submit(func() {
		if err := wp.repo.Create(student); err != nil {
			log.Printf("Insert error: %v", err)
		}
	})
	if err != nil {
		log.Printf("Pool submit error: %v", err)
	}
}


func (wp *WorkerPool) SubmitStudentWithTracking(student *domain.Student, jobID string, jm *JobManager) {
	jm.AddJob(jobID)

	err := wp.pool.Submit(func() {
		if err := wp.repo.Create(student); err != nil {
			jm.UpdateStatus(jobID, JobFailed, err.Error())
			return
		}
		jm.UpdateStatus(jobID, JobDone, "")
	})

	if err != nil {
		jm.UpdateStatus(jobID, JobFailed, "submit to pool failed: "+err.Error())
	}
}
