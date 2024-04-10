package controllers

import (
	"log"
	"time"

	"surlit/internal/logic/interfaces"
)

type HealthCheckService struct {
	checks     []interfaces.HealthCheck
	timeout    time.Duration
	timePeriod time.Duration
}

var _ interfaces.HealthCheckService = (*HealthCheckService)(nil)

func NewHealthCheckService() *HealthCheckService {
	return &HealthCheckService{
		checks:     nil,
		timeout:    time.Minute,
		timePeriod: time.Minute * 5,
	}
}

func (h *HealthCheckService) RegisterChecks(checks ...interfaces.HealthCheck) interfaces.HealthCheckService {
	h.checks = append(h.checks, checks...)
	return h
}

func (h *HealthCheckService) SetTimeout(timeout time.Duration) interfaces.HealthCheckService {
	h.timeout = timeout
	return h
}

func (h *HealthCheckService) SetTimePeriod(timePeriod time.Duration) interfaces.HealthCheckService {
	h.timePeriod = timePeriod
	return h
}

func (h *HealthCheckService) Check(check interfaces.HealthCheck) {
	if _, err := check.Exec(); err != nil { // TODO think about details and context
		log.Printf("error checking \"%s\"\n", check.Name())
	} else {
		log.Printf("all ok checking \"%s\"\n", check.Name())
	}
}

func (h *HealthCheckService) StartService() {
	ticker := time.NewTicker(h.timePeriod)
	go func() {
		for range ticker.C {
			for _, check := range h.checks {
				h.Check(check)
			}
		}
	}()
}
