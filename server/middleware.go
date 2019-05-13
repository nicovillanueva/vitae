package server

import (
	"runtime"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

type sysStats struct {
	CurrentMemAlloc uint64 `json:"current_memory"`
	SysReserved     uint64 `json:"sys_reserved"`
	HeapLiveObjects uint64 `json:"heap_live_objects"`
	GCRuns          uint32 `json:"gc_runs"`
	LastGC          uint64 `json:"last_gc_in_ns"`
}

type requestors struct {
	Requests   int      `json:"processed_requests"`
	UserAgents []string `json:"uniq_user_agents"`
}

func (r *requestors) addUserAgent(ua string) {
	for _, nua := range r.UserAgents {
		if nua == ua {
			return
		}
	}
	r.UserAgents = append(r.UserAgents, ua)
}

type stats struct {
	System     sysStats   `json:"system"`
	BootTime   time.Time  `json:"serving_since"`
	Uptime     string     `json:"uptime"`
	Requestors requestors `json:"request_stats"`
	mtx        sync.RWMutex
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

		r := c.Request()
		s.Requestors.Requests++
		ua := r.UserAgent()
		s.Requestors.addUserAgent(ua)
		if len(ua) > 50 {
			ua = ua[:50]
		}
		logger.Debugf("[%s %s] [%s...] -> %d",
			r.Method, r.URL.Path, ua, c.Response().Status)
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
		System:     st,
		BootTime:   s.BootTime,
		Requestors: s.Requestors,
		Uptime:     time.Now().Sub(s.BootTime).String(),
	}
}

// @Summary Statistics
// @Description Some statistics about this API
// @Produce json
// @Success 200 {object} server.stats "Some stats"
// @tags Others
// @Router /stats [GET]
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
