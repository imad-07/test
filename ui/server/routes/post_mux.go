package routes

import (
	"database/sql"
	"net/http"
	"regexp"
	"time"

	handler "forum/server/handlers"
	ratelimiter "forum/server/middlewares/rate_limiter"
)

type PostMuxType struct {
	DB *sql.DB
}

// comments url pattern
var commentsPathPattern = regexp.MustCompile(`/api/post/(\d+)/comments/(\d+)`)

// check url for post (get a specific post using id)
var specificPostPattern = regexp.MustCompile(`^/api/post/(\d+)$`)

func (pm *PostMuxType) PostMux(w http.ResponseWriter, r *http.Request) {
	// Extract path
	path := r.URL.Path

	// Post Handler Type
	postHandler := handler.NewPostHandler(pm.DB)

	if r.Method == http.MethodGet {
		if specificPostPattern.MatchString(path) {
			postHandler.GetPostByIdHandler(w, r)
		} else if path == "/api/post/" {
			postHandler.GetPostHandler(w, r)
		} else if commentsPathPattern.MatchString(path) {
			commentsHandler := handler.NewCommentHandler(pm.DB)
			commentsHandler.GetCommentsHandler(w, r)

		} else {
			http.Error(w, "Page not found", http.StatusNotFound)
			return
		}
	} else if r.Method == http.MethodPost {
		if specificPostPattern.MatchString(path) || commentsPathPattern.MatchString(path) {
			http.Error(w, "Page not found", http.StatusNotFound)
			return
		} else if path == "/api/post/" {
			ratelimiter.AddPostLimter.RateMiddleware(http.HandlerFunc(postHandler.CreatePostHandler),
				100000, 10*time.Second, pm.DB).ServeHTTP(w, r)
				return
		} else {
			http.Error(w, "Page not found", http.StatusNotFound)
			return
		}
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
