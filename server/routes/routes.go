package routes

import (
	"database/sql"
	"net/http"
	"time"

	handler "forum/server/handlers"
	"forum/server/middlewares"
	ratelimiter "forum/server/middlewares/rate_limiter"
)

func Routes(db *sql.DB) *http.ServeMux {
	infoHandler := handler.NewInfoHandler(db)
	commentHandler := handler.NewCommentHandler(db)
	userHandler := handler.NewUserHandler(db)
	reactHandler := handler.NewReactHandler(db)

	postMultiplexer := PostMuxType{DB: db}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("../client/assets"))
	mux.Handle("/client/assets/", http.StripPrefix("/client/assets", fs))
	sf := http.FileServer(http.Dir("../ui"))
	mux.Handle("/ui/", http.StripPrefix("/ui", sf))

	addCommentHandler := ratelimiter.AddCommentsLimter.RateMiddleware(http.HandlerFunc(commentHandler.AddCommentHandler), 100, 2*time.Second, db)

	mux.Handle("/api/comment", middlewares.MethodMiddleware(addCommentHandler, http.MethodPost))

	// mini multeplexer (has all the enpoints of post)
	// like: getting and adding posts and getting comment on a specific post
	mux.HandleFunc("/api/post", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path += "/"
		postMultiplexer.PostMux(w, r)
	})
	mux.HandleFunc("/api/post/", postMultiplexer.PostMux)

	//mux.HandleFunc("GET /post/", handler.ServePostPage)

	// Home Endpoint
	mux.HandleFunc("/", handler.HomeHandler)

	// Info Endpoint
	mux.HandleFunc("GET /api/info", infoHandler.Info)

	// Auth Endpoint
	loginRateLimiter := ratelimiter.LoginLimiter.RateMiddlewareAuth(http.HandlerFunc(userHandler.LoginHandler), 500, time.Minute)
	signupRateLimiter := ratelimiter.SignupLimiter.RateMiddlewareAuth(http.HandlerFunc(userHandler.SignUpHandler), 5, time.Minute)
	mux.Handle("/api/login", loginRateLimiter)
	//mux.HandleFunc("/login", userHandler.ServeLoginPage)
	mux.Handle("/api/signup", signupRateLimiter)
	//mux.HandleFunc("/signup", userHandler.ServeSignUpPage)
	//mux.HandleFunc("/commingsoon", handler.ServeFeaturesPage)

	// Reactions endpoint
	reactionRateLimiter := ratelimiter.ReactionsLimiter.RateMiddleware(http.HandlerFunc(reactHandler.ReactHandler), 10, 500*time.Millisecond, db)
	mux.Handle("/api/reaction", reactionRateLimiter)

	go func() {
		for {
			time.Sleep(120 * time.Minute)
			ratelimiter.AddCommentsLimter.RemoveSleepUsers()
			ratelimiter.AddPostLimter.RemoveSleepUsers()
			ratelimiter.LoginLimiter.RemoveSleepUsers()
			ratelimiter.ReactionsLimiter.RemoveSleepUsers()
			ratelimiter.SignupLimiter.RemoveSleepUsers()
		}
	}()

	return mux
}
