package server

import (
	"github.com/labstack/echo/v4"
	"runtime"
	"sync"
	"time"
)

type sysStats struct {
	CurrentMemAlloc uint64 `json:"current_memory"`
	SysReserved     uint64 `json:"sys_reserved"`
	HeapLiveObjects uint64 `json:"heap_live_objects"`
	GCRuns          uint32 `json:"gc_runs"`
	LastGC          uint64 `json:"last_gc_in_ns"`
}

type stats struct {
	System   sysStats  `json:"system"`
	BootTime time.Time `json:"serving_since"`
	Uptime   string    `json:"uptime"`
	Requests int       `json:"processed_requests"`
	mtx      sync.RWMutex
}

func newStats() *stats {
	return &stats{
		BootTime: time.Now(),
	}
}

func (s *stats) midProcess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		s.mtx.Lock()
		defer s.mtx.Unlock()
		s.Requests++
		return nil
	}
}

func (s *stats) current() stats {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	st := sysStats{
		CurrentMemAlloc: ms.Alloc,
		SysReserved:     ms.Sys,
		HeapLiveObjects: ms.Mallocs - ms.Frees,
		GCRuns:          ms.NumGC,
		LastGC:          ms.LastGC,
	}
	return stats{
		System:   st,
		BootTime: s.BootTime,
		Requests: s.Requests,
		Uptime:   time.Now().Sub(s.BootTime).String(),
	}
}

func (s *stats) HandleStats(c echo.Context) error {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return c.JSON(200, s.current())
}

// -----

func midContactHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Contact-Me", "nicovillanueva@pm.me")
		return next(c)
	}
}
