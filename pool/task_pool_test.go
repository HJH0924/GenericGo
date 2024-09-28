// Package pool
/**
* @Project : GenericGo
* @File    : task_pool_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/9/28 17:12
**/

package pool

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewOnDemandBlockTaskPool(t *testing.T) {
	tests := []struct {
		name             string
		initGo           int
		taskQueueSize    int
		queueBacklogRate float64
		coreGo           int
		maxGo            int
		maxIdleTime      time.Duration
		wantErr          error
	}{
		{
			name:             "all params valid",
			initGo:           1,
			taskQueueSize:    10,
			queueBacklogRate: 0.5,
			coreGo:           5,
			maxGo:            10,
			maxIdleTime:      5 * time.Second,
		},
		{
			name:             "initGo invalid",
			initGo:           0,
			taskQueueSize:    10,
			queueBacklogRate: 0.5,
			coreGo:           5,
			maxGo:            10,
			maxIdleTime:      5 * time.Second,
			wantErr:          NewErrInitGoInvalid,
		},
		{
			name:             "taskQueueSize invalid",
			initGo:           1,
			taskQueueSize:    -1,
			queueBacklogRate: 0.5,
			coreGo:           5,
			maxGo:            10,
			maxIdleTime:      5 * time.Second,
			wantErr:          NewErrTaskQueueSizeInvalid,
		},
		{
			name:             "queueBacklogRate invalid",
			initGo:           1,
			taskQueueSize:    10,
			queueBacklogRate: -1,
			coreGo:           5,
			maxGo:            10,
			maxIdleTime:      5 * time.Second,
			wantErr:          NewErrQueueBacklogRateInvalid,
		},
		{
			name:             "initGo, coreGo, maxGo invalid",
			initGo:           1,
			taskQueueSize:    10,
			queueBacklogRate: 0.5,
			coreGo:           -5,
			maxGo:            10,
			maxIdleTime:      5 * time.Second,
			wantErr:          NewErrGoroutineConditionNotMet,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pool, err := NewOnDemandBlockTaskPool(tt.initGo, tt.taskQueueSize,
				WithQueueBacklogRate(tt.queueBacklogRate),
				WithCoreGo(tt.coreGo),
				WithMaxGo(tt.maxGo),
				WithMaxIdleTime(tt.maxIdleTime))
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, int32(tt.initGo), pool.initGo)
				assert.Equal(t, tt.queueBacklogRate, pool.queueBacklogRate)
				assert.Equal(t, int32(tt.coreGo), pool.coreGo)
				assert.Equal(t, int32(tt.maxGo), pool.maxGo)
				assert.Equal(t, tt.maxIdleTime, pool.maxIdleTime)
			}
		})
	}
}
