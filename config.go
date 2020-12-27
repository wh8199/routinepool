package routinepool

import (
	"time"

	"github.com/wh8199/log"
)

const (
	DefaultWorkerCleanInterval = time.Minute
	DefaultMaxIdleTime         = time.Minute
)

type RoutinePoolConfig struct {
	MaxWorkperNumber int64            `json:"maxWorkerNumber"`
	CleanInterval    time.Duration    `json:"cleanInterval"`
	MaxIdleTime      time.Duration    `json:"maxIdleTime"`
	LogLevel         log.LoggingLevel `json:"logLevel"`
}

func DefaultRouterPoolConfig() *RoutinePoolConfig {
	return &RoutinePoolConfig{
		MaxWorkperNumber: 1024,
		CleanInterval:    DefaultWorkerCleanInterval,
		MaxIdleTime:      DefaultMaxIdleTime,
		LogLevel:         log.ERROR_LEVEL,
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

func (r *RoutinePoolConfig) WithLogLevel(logLevel log.LoggingLevel) {
	r.LogLevel = logLevel
}
