package scheduler

// SchedulerMock mocks a scheduler
type SchedulerMock struct {
	cronExpression string
	scheduledTask  Task
	scheduleJobErr error
	IScheduler
}

func (s *SchedulerMock) ScheduleJob(cronExpression string, task Task) (string, error) {
	s.cronExpression = cronExpression
	s.scheduledTask = task
	return "id", s.scheduleJobErr
}

func (s *SchedulerMock) SetScheduleJobErr(err error) {
	s.scheduleJobErr = err
}

func (s *SchedulerMock) GetLatestCronExpression() string {
	return s.cronExpression
}

func (s *SchedulerMock) Start() {}

func (s *SchedulerMock) Shutdown() error {
	return nil
}

func (s *SchedulerMock) Reset() {
	s.scheduleJobErr = nil
	s.cronExpression = ""
}
func (s *SchedulerMock) GetScheduledTask() Task {
	return s.scheduledTask
}

func NewMockScheduler() *SchedulerMock {
	return &SchedulerMock{}
}
