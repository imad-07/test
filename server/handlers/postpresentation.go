package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"forum/server/data"
	"forum/server/helpers"
	postapp "forum/server/service"
	"forum/server/shareddata"

	"github.com/mattn/go-sqlite3"
)

type PostHandler struct {
	PostService postapp.PostService
}

type PostResponse struct {
	Posts    []shareddata.Post
	MetaData shareddata.PostMetaData
}

func NewPostHandler(Db *sql.DB) *PostHandler {
	reactData := data.ReactionDB{
		DB: Db,
	}
	postData := data.PostData{
		Reactdb: reactData,
		Db:      Db,
	}

	postService := postapp.PostService{
		PostData: postData,
	}

	return &PostHandler{PostService: postService}
}

func (p *PostHandler) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(shareddata.SessionName)
	if err != nil || !helpers.CheckExpiredCookie(cookie.Value, time.Now(), p.PostService.PostData.Db) {
		helpers.WriteJson(w, http.StatusUnauthorized, struct {
			Error string `json:"error"`
		}{Error: "Unauthorized"})
		return
	}

	var post shareddata.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		helpers.WriteJson(w, http.StatusBadRequest, struct {
			Error string `json:"error"`
		}{Error: err.Error()})
		return
	}

	err = p.PostService.CreatePost(post, cookie.Value)
	if err != nil {
		switch err.Error() {
		case shareddata.PostErrors.ContentLength:
			fmt.Println(1)
			helpers.WriteJson(w, http.StatusBadRequest, struct {
				Error string `json:"error"`
			}{Error: err.Error()})
			return
		case shareddata.PostErrors.TitleLength:
			fmt.Println(2)
			helpers.WriteJson(w, http.StatusBadRequest, struct {
				Error string `json:"error"`
			}{Error: err.Error()})
			return
		case shareddata.PostErrors.CategoryDoesntExist:
			fmt.Println(3)
			helpers.WriteJson(w, http.StatusBadRequest, struct {
				Error string `json:"error"`
			}{Error: err.Error()})
			return
		case shareddata.UserErrors.UserNotExist:
			fmt.Println(4)
			helpers.WriteJson(w, http.StatusBadRequest, struct {
				Error string `json:"error"`
			}{Error: err.Error()})
			return
		case sql.ErrNoRows.Error():
			fmt.Println(err)
			helpers.WriteJson(w, http.StatusBadRequest, struct {
				Error string `json:"error"`
			}{Error: err.Error()})
			return
		}

		log.Println("Unexpected error", err)
		helpers.WriteJson(w, http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (p *PostHandler) GetPostHandler(w http.ResponseWriter, r *http.Request) {
	num, _ := strconv.Atoi(r.URL.Query().Get("page-number"))

	postsMetaData, err := p.PostService.GetPostMetaData()
	if err != nil {
		if err == sqlite3.ErrLocked {
			helpers.WriteJson(w, http.StatusLocked, struct {
				Error string `json:"error"`
			}{Error: "Database Locked"})
			return
		}
	}

	var posts []shareddata.Post

	cookie, err := r.Cookie(shareddata.SessionName)
	id := 0
	if err != http.ErrNoCookie && helpers.CheckExpiredCookie(cookie.Value, time.Now(), p.PostService.PostData.Db) {
		_, id = postapp.GetUser(p.PostService.PostData.Db, cookie.Value)
	}

	query := r.URL.Query()

	if (len(query["categorie"]) != 0 && query["categorie"][0] != "") || (len(query["posts"]) != 0 && query["posts"][0] != "") {
		posts, err := p.PostService.FilterPosts(num, posts, r, id)
		if err != nil {
			if err == sql.ErrNoRows {
				helpers.WriteJson(w, http.StatusOK, PostResponse{MetaData: postsMetaData, Posts: []shareddata.Post{}})
				return
			}
			helpers.WriteJson(w, http.StatusInternalServerError, PostResponse{MetaData: postsMetaData, Posts: []shareddata.Post{}})
			return
		}
		helpers.WriteJson(w, http.StatusOK, PostResponse{Posts: posts, MetaData: postsMetaData})
	} else {
		posts, err := p.PostService.GetPost(num, id)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				helpers.WriteJson(w, http.StatusOK, PostResponse{MetaData: postsMetaData, Posts: []shareddata.Post{}})
				return
			case sqlite3.ErrLocked:
				helpers.WriteJson(w, http.StatusLocked, struct {
					Error string `json:"error"`
				}{Error: "Database Locked"})
				return
			}

			log.Println("Unexpected error", err)
			helpers.WriteJson(w, http.StatusInternalServerError, struct {
				Error string `json:"error"`
			}{Error: err.Error()})
			return
		}
		helpers.WriteJson(w, http.StatusOK, PostResponse{Posts: posts, MetaData: postsMetaData})
	}
}

func (p *PostHandler) GetPostByIdHandler(w http.ResponseWriter, r *http.Request) {
	splitedUrl := strings.Split(r.URL.Path, "/")

	postID, err := strconv.Atoi(splitedUrl[3])
	if err != nil {
		helpers.WriteJson(w, http.StatusNotFound, struct {
			Error string `json:"error"`
		}{Error: "page not found"})
		return
	}
	cookie, err := r.Cookie(shareddata.SessionName)
	id := 0
	if err != http.ErrNoCookie && helpers.CheckExpiredCookie(cookie.Value, time.Now(), p.PostService.PostData.Db) {
		_, id = postapp.GetUser(p.PostService.PostData.Db, cookie.Value)
	}

	post, err := p.PostService.GetSinglePost(postID, id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			helpers.WriteJson(w, http.StatusOK, post)
			return
		case sqlite3.ErrLocked:
			helpers.WriteJson(w, http.StatusLocked, struct {
				Error string `json:"error"`
			}{Error: "Database Locked"})
			return
		}

		helpers.WriteJson(w, http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{Error: "Internal Server Error"})
		return
	}

	helpers.WriteJson(w, http.StatusOK, post)
}

/*func ServePostPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../client/templates/post.html")
	if err != nil {
		//ErrorHandler(w, http.StatusInternalServerError, "inernal Server Error", "Error While Parsing index.html")
		log.Println("Unexpected error", err)
		return
	}
	t.Execute(w, nil)
}*/
