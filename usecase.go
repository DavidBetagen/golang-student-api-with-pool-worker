type studentUsecase struct {
	repo       domain.StudentRepository
	workerPool *WorkerPool
	jobManager *JobManager
}

func NewStudentUsecase(r domain.StudentRepository, pool *WorkerPool, jm *JobManager) domain.StudentUsecase {
	return &studentUsecase{
		repo:       r,
		workerPool: pool,
		jobManager: jm,
	}
}

func (uc *studentUsecase) CreateAsync(student *domain.Student) (string, error) {
	jobID := uuid.New().String()
	uc.workerPool.SubmitStudentWithTracking(student, jobID, uc.jobManager)
	return jobID, nil
}

func (uc *studentUsecase) GetJobStatus(id string) (*JobInfo, bool) {
	return uc.jobManager.GetJob(id)
}
