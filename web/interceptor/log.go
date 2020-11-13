package interceptor

import (
	"net/http"
	"time"

	"workWechat_api/pkg/logging"
)

func TimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		// next handler

		next.ServeHTTP(wr, r)
		timeElapsed := time.Since(timeStart)
		logging.Info("%dms %s %s %s %s %s",
			timeElapsed.Milliseconds(),
			r.Method ,
			r.URL.String(),
			r.Header["User-Agent"][0],
			r.Host,
			r.RemoteAddr)

	})
}
