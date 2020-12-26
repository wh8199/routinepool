package routinepool

import "time"

const (
	DefaultWorkerCleanInterval = time.Minute
	DefaultMaxIdleTime         = time.Minute
)

type RoutinePoolConfig struct {
	MaxWorkperNumber int64         `json:"maxWorkerNumber"`
	CleanInterval    time.Duration `json:"cleanInterval"`
	MaxIdleTime      time.Duration `json:"maxIdleTime"`
}

func DefaultRouterPoolConfig() *RoutinePoolConfig {
	return &RoutinePoolConfig{
		MaxWorkperNumber: 1024,
		CleanInterval:    DefaultWorkerCleanInterval,
		MaxIdleTime:      DefaultMaxIdleTime,
	}
}

func (r *RoutinePoolConfig) WithMaxWorkerNumber(maxWorkerNumber int64) {
	r.MaxWorkperNumber = maxWorkerNumber
}

func (r *RoutinePoolConfig) WithCleanWorkerInterval(cleanInterval string) {
	interval, err := time.ParseDuration(cleanInterval)
	if err != nil {
		r.CleanInterval = DefaultWorkerCleanInterval
		return
	}

	r.CleanInterval = interval
}

func (r *RoutinePoolConfig) WithMaxIdleTime(maxIdleTime string) {
	maxIdleTimeDuration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		r.MaxIdleTime = DefaultMaxIdleTime
		return
	}

	r.MaxIdleTime = maxIdleTimeDuration
}
