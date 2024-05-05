package mocks

import "github.com/srinivasaleti/data-sync-manager/orchestrator/scheduler"

// SchedulerMock mocks a scheduler
type SchedulerMock struct {
	cronExpression string
	scheduledTask  scheduler.Task
	scheduleJobErr error
	scheduler.IScheduler
}

func (s *SchedulerMock) ScheduleJob(cronExpression string, task scheduler.Task) (string, error) {
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
func (s *SchedulerMock) GetScheduledTask() scheduler.Task {
	return s.scheduledTask
}

func NewMockScheduler() *SchedulerMock {
	return &SchedulerMock{}
}
