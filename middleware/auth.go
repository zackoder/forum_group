package middleware

import (
	"net/http"
	"sync"
	"time"

	"forum/utils"
)

var (
	rateLimitData = make(map[string]*RateLimitInfo)
	mu            sync.Mutex
	rateLimit     = 10
	rateInterval  = time.Second
)

type RateLimitInfo struct {
	RequestCount int
	LastRequest  time.Time
}

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		token, err := r.Cookie("token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		if !allowRequest(token.Value) {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		query := `SELECT user_id FROM sessions WHERE token=?;`
		var userID int
		err = utils.DB.QueryRow(query, token.Value).Scan(&userID)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next(w, r)
	}
}

func allowRequest(token string) bool {
	mu.Lock()
	defer mu.Unlock()

	now := time.Now()

	// Retrieve or initialize rate limit data for the token
	info, exists := rateLimitData[token]
	if !exists {
		rateLimitData[token] = &RateLimitInfo{
			RequestCount: 1,
			LastRequest:  now,
		}
		return true
	}

	// Reset count if the time window has passed
	if now.Sub(info.LastRequest) > rateInterval {
		info.RequestCount = 1
		info.LastRequest = now
		return true
	}
	if info.RequestCount < rateLimit {
		info.RequestCount++
		info.LastRequest = now
		return true
	}

	return false
}
