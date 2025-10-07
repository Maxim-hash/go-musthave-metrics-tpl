package models

import "strconv"

type MemStorage struct {
	counters map[string]int64
	gauges   map[string]float64
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		counters: make(map[string]int64),
		gauges:   make(map[string]float64),
	}
}

func (ms *MemStorage) UpdateCounter(name string, value int64) {
	ms.counters[name] += value
}

func (ms *MemStorage) UpdateGauge(name string, value float64) {
	ms.gauges[name] = value
}

func (ms *MemStorage) GetCounter(name string) (int64, bool) {
	val, ok := ms.counters[name]
	return val, ok
}

func (ms *MemStorage) GetGauge(name string) (float64, bool) {
	val, ok := ms.gauges[name]
	return val, ok
}

func (ms *MemStorage) GetAllMetrics() map[string]string {
	result := make(map[string]string)
	for k, v := range ms.counters {
		result[k] = strconv.FormatInt(v, 10)
	}
	for k, v := range ms.gauges {
		result[k] = strconv.FormatFloat(v, 'f', -1, 64)
	}
	return result
}
