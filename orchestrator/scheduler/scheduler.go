package scheduler

import (
	"github.com/go-co-op/gocron/v2"
)

type Scheduler struct {
	scheduler gocron.Scheduler
}

type IScheduler interface {
	ScheduleJob(cronExpression string, task Task) (string, error)
	Start()
	Shutdown() error
}

type Task func()

func (s *Scheduler) ScheduleJob(cronExpression string, task Task) (string, error) {
	job, err := s.scheduler.NewJob(
		gocron.CronJob(cronExpression, true),
		gocron.NewTask(task),
	)
	if err != nil {
		return "", err
	}
	return job.ID().String(), err
}

func (s *Scheduler) Start() {
	s.scheduler.Start()
}

func (s *Scheduler) Shutdown() error {
	return s.scheduler.Shutdown()
}

func New() (*Scheduler, error) {
	goCronScheduler, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}
	return &Scheduler{
		scheduler: goCronScheduler,
	}, nil
}
