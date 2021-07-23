package web

import (
	"github.com/gin-gonic/gin"
	"github.com/leungyauming/api/common/responses"
	"golang.org/x/time/rate"
	"sync"
)

type IPRateLimiter struct {
	eps      int
	limiters map[string]*rate.Limiter
	mu       sync.RWMutex
}

func NewIPRateLimiter(eps int) *IPRateLimiter {
	return &IPRateLimiter{
		eps:      eps,
		limiters: make(map[string]*rate.Limiter),
	}
}

func (irl *IPRateLimiter) PutLimiter(ip string) *rate.Limiter {
	lim := rate.NewLimiter(rate.Limit(irl.eps), irl.eps)

	irl.mu.Lock()
	irl.limiters[ip] = lim
	irl.mu.Unlock()

	return lim
}

func (irl *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	irl.mu.RLock()
	lim, found := irl.limiters[ip]
	irl.mu.RUnlock()
	if !found {
		return irl.PutLimiter(ip)
	}
	return lim
}

func (irl *IPRateLimiter) Middleware(ctx *gin.Context) {
	lim := irl.GetLimiter(ctx.ClientIP())
	if !lim.Allow() {
		responses.SendTooManyRequestsResponse(ctx)
		ctx.Abort()
		return
	}

	ctx.Next()
}
