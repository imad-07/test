package service

import (
	"errors"

	"forum/server/data"
	"forum/server/shareddata"
)

type ReactService struct {
	ReactData data.ReactionDB
}

func (data *ReactService) LikesTotal(thread_type string, thread_id int) (shareddata.ReactResponse, error) {
	response := shareddata.ReactResponse{}
	err := errors.New("")
	if thread_type == "post" {
		response.Like, response.Dislike, err = data.ReactData.CountPostLikes(thread_id)
	} else {
		response.Like, response.Dislike, err = data.ReactData.CountCommentLikes(thread_id)
	}
	if err != nil {
		return response, err
	}
	return response, nil
}

func (data *ReactService) GetLikedThread(thread_type string, thread_id, user_id int) (bool, bool) {
	isLiked, isDisliked := false, false
	if thread_type == "post" {
		isLiked, isDisliked = data.ReactData.CheckIfLikedPost(thread_id, user_id)
	} else {
		isLiked, isDisliked = data.ReactData.CheckIfLikedComment(thread_id, user_id)
	}
	return isLiked, isDisliked
}

func (data *ReactService) CheckReactInput(react shareddata.React) error {
	if (react.React != 2 && react.React != 1) || (react.Thread_type != "post" && react.Thread_type != "comment") || react.Thread_id < 0 || !data.ReactData.CheckIfThreadExists(react) {
		return errors.New("bad request")
	}
	return nil
}

func (data *ReactService) ReactionService(react shareddata.React, user_id int) error {
	if react.Thread_type == "post" {
		err := data.postReaction(react.Thread_id, user_id, react.React)
		if err != nil {
			return err
		}
		return nil
	} else {
		err := data.commentReaction(react.Thread_id, user_id, react.React)
		if err != nil {
			return err
		}
		return nil
	}
}

func (data *ReactService) postReaction(post_id, user_id, react int) error {
	var exists bool
	exists, err := data.ReactData.CheckPostReaction(user_id, post_id)
	if err != nil {
		return err
	}
	if !exists {
		err := data.ReactData.InsertReactPost(user_id, post_id, react)
		if err != nil {
			return err
		}
	} else {
		var like_type int
		like_type, err := data.ReactData.GetReactionTypePost(user_id, post_id)
		if err != nil {
			return err
		}
		if like_type == react {
			err := data.ReactData.DeleteReactionPost(user_id, post_id)
			if err != nil {
				return err
			}
		} else {
			err := data.ReactData.DeleteReactionPost(user_id, post_id)
			if err != nil {
				return err
			}
			err = data.ReactData.InsertReactPost(user_id, post_id, react)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (data *ReactService) commentReaction(comment_id, user_id, react int) error {
	var exists bool
	exists, err := data.ReactData.CheckCommentReaction(user_id, comment_id)
	if err != nil {
		return err
	}
	if !exists {
		err := data.ReactData.InsertReactComment(user_id, comment_id, react)
		if err != nil {
			return err
		}
	} else {
		var isLiked int
		isLiked, err := data.ReactData.GetReactionTypeComment(user_id, comment_id)
		if err != nil {
			return err
		}
		if isLiked == react {
			err := data.ReactData.DeleteReactionComment(user_id, comment_id)
			if err != nil {
				return err
			}
		} else {
			err := data.ReactData.DeleteReactionComment(user_id, comment_id)
			if err != nil {
				return err
			}
			err = data.ReactData.InsertReactComment(user_id, comment_id, react)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
