package scheduler

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestScheduleJob(t *testing.T) {
	scheduler, err := New()
	assert.NoError(t, err)

	t.Run("should run scheduler", func(t *testing.T) {
		counter := 0
		scheduler.ScheduleJob("* * * * * *", func() {
			counter = counter + 1
		})
		scheduler.Start()
		time.Sleep(2 * time.Second)
		scheduler.Shutdown()
		assert.Equal(t, counter, 2)
	})
	t.Run("should return error if there is an issue while running scheduler", func(t *testing.T) {
		_, err := scheduler.ScheduleJob("wrong_cron * * * * *", func() {
			fmt.Println("Hello")
		})
		assert.Error(t, err)
	})
}
