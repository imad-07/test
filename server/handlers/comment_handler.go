package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"forum/server/data"
	"forum/server/helpers"
	"forum/server/service"
	"forum/server/shareddata"

	"github.com/mattn/go-sqlite3"
)

type CommentHandler struct {
	CommentService service.CommentService
}

// CommentsResponse type => type that the GetCommentsHandler will return it
type CommentsResponse struct {
	Comments []shareddata.ShowComment
	MetaData shareddata.CommentMetaData
}

// New Comment Handler function
func NewCommentHandler(Db *sql.DB) CommentHandler {
	reactData := data.ReactionDB{
		DB: Db,
	}
	commentData := data.CommentData{
		Reactdb: reactData,
		DB:      Db,
	}

	userData := data.UserData{
		DB: Db,
	}

	commentService := service.CommentService{
		CommentData: commentData,
		UserData:    userData,
	}

	return CommentHandler{
		CommentService: commentService,
	}
}

func (h *CommentHandler) AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	// parse data
	comment := shareddata.Comment{}
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	cookie, err := r.Cookie(shareddata.SessionName)
	if err != nil || !helpers.CheckExpiredCookie(cookie.Value, time.Now(), h.CommentService.CommentData.DB) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Validate Inputs
	err = h.CommentService.ValidateInput(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user exist
	userExist := h.CommentService.CheckUserExist(cookie.Value)
	if !userExist {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// add userUid to the comment
	comment.UserUID = cookie.Value

	// add the comment using the AddComment from the service layer
	err = h.CommentService.AddComment(comment)
	if err != nil {
		switch err.Error() {
		case sqlite3.ErrLocked.Error():
			http.Error(w, "Database locked", http.StatusLocked)
			return
		case shareddata.PostErrors.PostNotExist:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		default:
			log.Printf("Unexpected Error when we add comment %s", err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	// return a success response
	helpers.WriteJson(w, http.StatusCreated, struct{ Message string }{
		Message: "Your comment added successfuly",
	})
}

func (h *CommentHandler) GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	// extract post id and comment page number from the path
	postId, pageNumber, err := extractPostAndPage(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get Comments Metadata
	commentsMetaData, err := h.CommentService.GetCommentMetaData(postId)
	if err != nil {
		if err == sqlite3.ErrLocked {
			http.Error(w, "Database locked", http.StatusLocked)
			return
		}

		log.Printf("Unexpected Error when we get meta data for comments %s", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	cookie, err := r.Cookie(shareddata.SessionName)
	id := 0
	if err != http.ErrNoCookie {
		_, id = service.GetUser(h.CommentService.CommentData.DB, cookie.Value)
	}
	// Get Comments
	comments, err := h.CommentService.GetComments(postId, pageNumber, id)
	if err != nil {
		switch err.Error() {
		case sqlite3.ErrLocked.Error():
			http.Error(w, "Database locked", http.StatusLocked)
			return
		case shareddata.CommentErrors.InvalidPage:
			// Send Empty Array of Comments To the user
			helpers.WriteJson(w, http.StatusOK, CommentsResponse{MetaData: commentsMetaData, Comments: []shareddata.ShowComment{}})
			return
		default:
			log.Printf("Unexpected Error when we get comment %s", err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	// Send Comments To the user
	helpers.WriteJson(w, http.StatusOK,
		CommentsResponse{Comments: comments, MetaData: commentsMetaData})
}

// take a path with this form /api/post/2/comment/1
// where 2 is the post id and 1 is the comment page number
// and return the post id and the comment page number
func extractPostAndPage(path string) (int, int, error) {
	splitedPath := strings.Split(path, "/")

	if len(splitedPath) != 6 {
		return 0, 0, errors.New("invalid path")
	}

	postId, err := strconv.Atoi(splitedPath[3]) // 3 = post id
	if err != nil {
		return 0, 0, errors.New("post id doesn't valid")
	}

	pageNumber, err := strconv.Atoi(splitedPath[5]) // 5 = comment page
	if err != nil {
		return 0, 0, errors.New("page number doesn't valid")
	}

	return postId, pageNumber, nil
}
