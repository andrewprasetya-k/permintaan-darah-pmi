package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

// RateLimiter returns a rate limiting middleware
// rate format: "100-H" = 100 requests per hour
// "10-M" = 10 requests per minute
// "1-S" = 1 request per second
func RateLimiter(rate string) gin.HandlerFunc {
    // Parse rate format dengan benar
    parsedRate, err := limiter.NewRateFromFormatted(rate)
    if err != nil {
        panic("Invalid rate format: " + rate) // fail fast saat startup
    }

    store := memory.NewStore()
    instance := limiter.New(store, parsedRate) // ← pakai parsedRate

    return func(c *gin.Context) {
        key := c.ClientIP()
        limiterCtx, err := instance.Get(c.Request.Context(), key)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Internal server error",
            })
            c.Abort()
            return
        }

        c.Header("X-RateLimit-Limit", strconv.FormatInt(limiterCtx.Limit, 10))
        c.Header("X-RateLimit-Remaining", strconv.FormatInt(limiterCtx.Remaining, 10))
        c.Header("X-RateLimit-Reset", strconv.FormatInt(limiterCtx.Reset, 10))

        if limiterCtx.Reached {
            resetTime := time.Unix(limiterCtx.Reset, 0)
            retryAfterSeconds := int64(time.Until(resetTime).Seconds())
            if retryAfterSeconds < 1 {
                retryAfterSeconds = 1
            }

            c.Header("Retry-After", strconv.FormatInt(retryAfterSeconds, 10))
            c.JSON(http.StatusTooManyRequests, gin.H{
                "success":           false,
                "message":           "Rate limit exceeded. Please try again in " + strconv.FormatInt(retryAfterSeconds, 10) + " seconds.",
                "retryAfterSeconds": retryAfterSeconds,
            })
            c.Abort()
            return
        }

        c.Next()
    }
}

// LoginRateLimiter applies stricter rate limiting for login endpoints
// 5 attempts per minute per IP
func LoginRateLimiter() gin.HandlerFunc {
	return RateLimiter("5-M")
}

// GeneralAPIRateLimiter applies standard rate limiting for general endpoints
// 50 requests per minute per IP
func GeneralAPIRateLimiter() gin.HandlerFunc {
	return RateLimiter("100-M")
}

// StrictAPIRateLimiter applies strict rate limiting for sensitive endpoints
// 30 requests per minute per IP
func StrictAPIRateLimiter() gin.HandlerFunc {
	return RateLimiter("30-M")
}
