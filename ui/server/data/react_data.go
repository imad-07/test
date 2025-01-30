package data

import (
	"database/sql"

	"forum/server/shareddata"
)

type ReactionDB struct {
	DB *sql.DB
}

func (db *ReactionDB) CheckPostReaction(user_id, post_id int) (bool, error) {
	var exists bool
	err := db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM postReact WHERE user_Id = ? AND post_Id = ?)", user_id, post_id).Scan(&exists)
	return exists, err
}

func (db *ReactionDB) CheckCommentReaction(user_id, comment_id int) (bool, error) {
	var exists bool
	err := db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM commentReact WHERE user_Id = ? AND comment_Id = ?)", user_id, comment_id).Scan(&exists)
	return exists, err
}

func (db *ReactionDB) DeleteReactionPost(user_id, post_id int) error {
	_, err := db.DB.Exec("DELETE FROM postReact WHERE post_Id = ? AND user_Id = ?", post_id, user_id)
	return err
}

func (db *ReactionDB) DeleteReactionComment(user_id, post_id int) error {
	_, err := db.DB.Exec("DELETE FROM commentReact WHERE comment_Id = ? AND user_Id = ?", post_id, user_id)
	return err
}

func (db *ReactionDB) GetReactionTypePost(user_id, post_id int) (int, error) {
	var isLiked int
	err := db.DB.QueryRow("SELECT is_liked FROM postReact WHERE user_id = ? AND post_Id = ?", user_id, post_id).Scan(&isLiked)
	return isLiked, err
}

func (db *ReactionDB) GetReactionTypeComment(user_id, post_id int) (int, error) {
	var isLiked int
	err := db.DB.QueryRow("SELECT is_liked FROM commentReact WHERE user_id = ? AND comment_Id = ?", user_id, post_id).Scan(&isLiked)
	return isLiked, err
}

func (db *ReactionDB) InsertReactPost(user_id, post_id, like_type int) error {
	_, err := db.DB.Exec("INSERT INTO postReact (post_Id, user_Id, is_liked) VALUES (?,?,?)", post_id, user_id, like_type)
	return err
}

func (db *ReactionDB) InsertReactComment(user_id, post_id, like_type int) error {
	_, err := db.DB.Exec("INSERT INTO commentReact (comment_Id, user_Id, is_liked) VALUES (?,?,?)", post_id, user_id, like_type)
	return err
}

func (db *ReactionDB) CountPostLikes(post_id int) (int, int, error) {
	var likes, dislikes int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM postReact WHERE is_liked = 1 AND post_Id = ?", post_id).Scan(&likes)
	if err != nil {
		return 0, 0, err
	}

	err = db.DB.QueryRow("SELECT COUNT(*) FROM postReact WHERE is_liked = 2 AND post_Id = ?", post_id).Scan(&dislikes)
	if err != nil {
		return 0, 0, err
	}

	return likes, dislikes, nil
}

func (db *ReactionDB) CountCommentLikes(post_id int) (int, int, error) {
	var likes, dislikes int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM commentReact WHERE is_liked = 1 AND comment_Id = ?", post_id).Scan(&likes)
	if err != nil {
		return 0, 0, err
	}

	err = db.DB.QueryRow("SELECT COUNT(*) FROM commentReact WHERE is_liked = 2 AND comment_Id = ?", post_id).Scan(&dislikes)
	if err != nil {
		return 0, 0, err
	}

	return likes, dislikes, nil
}

func (d *ReactionDB) CheckIfLikedPost(post_id, user_id int) (bool, bool) {
	isLiked := 0
	d.DB.QueryRow("SELECT is_liked FROM postReact WHERE post_id = ? AND user_id = ?", post_id, user_id).Scan(&isLiked)
	switch isLiked {
	case 2:
		return false, true
	case 1:
		return true, false
	default:
		return false, false
	}
}

func (d *ReactionDB) CheckIfLikedComment(post_id, user_id int) (bool, bool) {
	isLiked := 0
	d.DB.QueryRow("SELECT is_liked FROM commentReact WHERE comment_id = ? AND user_id = ?", post_id, user_id).Scan(&isLiked)
	switch isLiked {
	case 2:
		return false, true
	case 1:
		return true, false
	default:
		return false, false
	}
}

func (d *ReactionDB) CheckIfThreadExists(react shareddata.React) bool {
	exist := false
	if react.Thread_type == "post" {
		d.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM post WHERE id = ?)", react.Thread_id).Scan(&exist)
		return exist
	} else {
		d.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM comment WHERE id = ?)", react.Thread_id).Scan(&exist)
		return exist
	}
}
