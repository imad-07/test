package service

import (
	"database/sql"
	"errors"
	"fmt"
	"html"
	"math"
	"net/http"
	"strings"

	"forum/server/data"
	"forum/server/shareddata"
)

type PostService struct {
	PostData data.PostData
}

func NewPostService(PostData *data.PostData) *PostService {
	return &PostService{PostData: *PostData}
}

// validate if the user exists using data layer getuser
func (a *PostService) ValidatePostInput(post shareddata.Post) error {
	if len(strings.TrimSpace(post.Title)) == 0 ||len(post.Title) > 500 {
		return fmt.Errorf(shareddata.PostErrors.TitleLength)
	}
	if len(strings.TrimSpace(post.Content)) == 0 || len(post.Content) > 5000 {
		return fmt.Errorf(shareddata.PostErrors.ContentLength)
	}
	if post.UserID == 0 {
		return fmt.Errorf(shareddata.UserErrors.UserNotExist)
	}
	return nil
}

// Add Posts Service
func (a *PostService) CreatePost(post shareddata.Post, uuid string) error {
	// validate post data and add user id to the post
	post.UserID = a.PostData.GetUser(uuid)
	post.Categories = removeDuplicate(post.Categories)
	err := a.ValidatePostInput(post)

	// fix html
	post.Title = html.EscapeString(post.Title)
	post.Content = html.EscapeString(post.Content)

	if err != nil {
		return err
	}
	postId, err := a.PostData.InsertPost(post)
	if err != nil {
		return err
	}
	post.ID = postId

	// Add Categories
	/*err = a.AddCategoriesToPost(post)
	if err != nil {
		if errDB := a.PostData.DeletePost(postId); errDB != nil {
			return errDB
		}
		return err
	}*/

	return nil
}

// Add categories to the posts
func (a *PostService) AddCategoriesToPost(post shareddata.Post) error {
	for _, category_name := range post.Categories {
		// check if the category is exist and get its id
		category_id, err := a.PostData.GetCategoryId(category_name)
		if err != nil {
			return err
		}
		if category_id == 0 {
			return errors.New(shareddata.PostErrors.CategoryDoesntExist)
		}

		// add the category to the post using category_post table
		err = a.PostData.AddCategory(post.ID, category_id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *PostService) GetPost(num, userID int) ([]shareddata.Post, error) {
	start := (num * data.PostsPerPage)
	total := 0
	err := a.PostData.Tablelen("post", &total)
	if err != nil {
		return nil, err
	}
	if start > total {
		return nil, sql.ErrNoRows
	}
	row, err := a.PostData.ExtractPosts(start)
	if err != nil {
		return nil, err
	}
	var posts []shareddata.Post
	for row.Next() {
		var post shareddata.Post

		err := row.Scan(&post.ID, &post.Title, &post.Content, &post.Created, &post.Author, &post.Likes, &post.Dislikes, &post.CommentsCount)
		if err != nil {
			return nil, err
		}
		post.IsLiked, post.IsDisliked = a.PostData.Reactdb.CheckIfLikedPost(post.ID, userID)
		// Get categories
		categories, err := a.PostData.GetPostCategories(post.ID)
		if err != nil {
			return nil, err
		}

		post.Categories = categories

		posts = append(posts, post)
	}
	if err := row.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func (a *PostService) GetPostMetaData() (shareddata.PostMetaData, error) {
	var totalPosts int
	err := a.PostData.Tablelen("post", &totalPosts)
	if err != nil {
		return shareddata.PostMetaData{}, err
	}

	postsPages := math.Ceil(float64(totalPosts) / float64(data.PostsPerPage))

	return shareddata.PostMetaData{PostsCount: totalPosts, PostsPages: int(postsPages), StandardCount: data.PostsPerPage}, nil
}

// Get categories

func (a *PostService) GetSinglePost(postId, id int) (shareddata.Post, error) {
	post, err := a.PostData.GetPost(postId, id)
	if err != nil {
		return shareddata.Post{}, err
	}
	categories, err := a.PostData.GetPostCategories(postId)
	if err != nil {
		return shareddata.Post{}, err
	}
	post.Categories = categories

	return post, nil
}

// page number * posts to start from with limit of posts
// check if the total tables len is less than total
func (p *PostService) FilterPosts(num int, posts []shareddata.Post, r *http.Request, id int) ([]shareddata.Post, error) {
	var err error
	query := r.URL.Query()
	number := num * data.PostsPerPage
	filterByCat := len(query["categorie"]) != 0 && query["categorie"][0] != ""
	filterByPosts := len(query["posts"]) != 0 && query["posts"][0] != ""

	if filterByCat && !filterByPosts {
		posts, err = p.PostData.GetCategoriesPost(query["categorie"][0], posts, number, id)
		if err != nil {
			return nil, err
		}

	}

	if filterByPosts && !filterByCat {
		posts, err = p.PostData.GetOwnPostsOrLikedPosts(query["posts"][0], posts, id, number)
		if err != nil {
			return nil, err
		}
	}

	if filterByCat && filterByPosts {
		posts, err = p.PostData.GetBothFilters(id, query["posts"][0], query["categorie"][0], posts, num)
		if err != nil {
			return nil, err
		}
	}

	return posts, nil
}
// helpers
func removeDuplicate(categories []string) []string {
	categoryMap := map[string]bool{}

	newCategories := []string{}
	for _, category := range categories {
		if !categoryMap[category] {
			newCategories = append(newCategories, category)
		}
		categoryMap[category] = true
	}

	return newCategories
}

