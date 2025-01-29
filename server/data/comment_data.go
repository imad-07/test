package data

import (
	"database/sql"

	"forum/server/shareddata"
)

type CommentData struct {
	DB      *sql.DB
	Reactdb ReactionDB
}

var CommentsPerPage = 20

// Insert a comment into the comment table in the database
func (database *CommentData) InsertComment(comment shareddata.Comment) error {
	_, err := database.DB.Exec(
		"INSERT INTO comment (user_id, post_id, content) VALUES (?, ?, ?)",
		comment.UserId, comment.PostId, comment.Content)

	return err
}

// Check if the post is exist using the id
func (database *CommentData) CheckPostExist(id int) bool {
	err := database.DB.QueryRow("SELECT id FROM post WHERE id = ?", id).Scan(&id)
	return err == nil
}

// Get Comments from a specific comment row (like from 1 and get the 100 in front of it)
func (database *CommentData) GetCommentsFrom(from, postId, userId int) ([]shareddata.ShowComment, error) {
	rows, err := database.DB.Query(
		`SELECT comment_id, comment_author, comment_content, comment_date, comment_likes, comment_dislikes
		FROM single_comment
		WHERE post_id = ? ORDER BY comment_date DESC LIMIT ? OFFSET ?`,
		postId, CommentsPerPage, from)
	if err != nil {
		return nil, err
	}

	var comments []shareddata.ShowComment
	for rows.Next() {
		var comment shareddata.ShowComment
		rows.Scan(&comment.Id, &comment.Author, &comment.Content, &comment.Date, &comment.Likes, &comment.Dislikes)
		comment.IsLiked, comment.IsDisliked = database.Reactdb.CheckIfLikedComment(comment.Id, userId)
		comments = append(comments, comment)
	}

	return comments, err
}

// Get comments count
func (database *CommentData) GetCommentsCount(postId int) (int, error) {
	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM comment WHERE post_id = ?", postId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
