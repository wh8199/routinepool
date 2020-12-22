package routinepool

type RoutinePoolConfig struct {
	MaxWorkperNumber int64 `json:"maxWorkerNumber"`
}

func DefaultRouterPoolConfig() *RoutinePoolConfig {
	return &RoutinePoolConfig{
		MaxWorkperNumber: 1024,
	}
}

func (r *RoutinePoolConfig) WithMaxWorkerNumber(maxWorkerNumber int64) *RoutinePoolConfig {
	r.MaxWorkperNumber = maxWorkerNumber

	return r
}
