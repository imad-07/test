package data

import (
	"database/sql"
	"fmt"

	"forum/server/shareddata"
)

var PostsPerPage = 15 // Number of  posts to display

type PostData struct {
	Reactdb ReactionDB
	Db      *sql.DB
}

func NewDAL(Db *sql.DB) *PostData {
	return &PostData{Db: Db}
}

func (d *PostData) GetUser(uid string) int {
	var id int

	err := d.Db.QueryRow("SELECT id FROM user_profile WHERE uid = ?", uid).Scan(&id)
	if err != nil {
		return 0
	}

	return id
}

func (d *PostData) InsertPost(post shareddata.Post) (int, error) {
	query := "INSERT INTO post (title, content, user_id) VALUES (?, ?, ?)"
	rowResult, err := d.Db.Exec(query, post.Title, post.Content, post.UserID)
	if err != nil {
		fmt.Println("chob")
		return 0, err
	}

	rowId, err := rowResult.LastInsertId()
	if err != nil {
		//return 0, err
	}

	return int(rowId), nil
}

func (d *PostData) Tablelen(table string, total *int) error {
	err := d.Db.QueryRow("SELECT COUNT(*) FROM " + table).Scan(total)
	return err
}

func (d *PostData) ExtractPosts(start int) (*sql.Rows, error) {
	rows, err := d.Db.Query(`SELECT post_id, post_title, post_content, post_date, post_author, post_likes, post_dislikes, post_comments_count
	FROM single_post
   ORDER BY post_date DESC LIMIT ? OFFSET ?`, PostsPerPage, start)
	if err != nil {
		return nil, err
	}
	return rows, err
}

func (d *PostData) GetPost(id, user_id int) (shareddata.Post, error) {
	var post shareddata.Post

	row := d.Db.QueryRow(`SELECT  post_id, post_title, post_content, post_date, post_author, post_likes, post_dislikes, post_comments_count, joined_at
	FROM single_post
	WHERE post_id = ?`, id)

	if row.Err() != nil {
		return shareddata.Post{}, row.Err()
	}
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.Created, &post.Author, &post.Likes, &post.Dislikes, &post.CommentsCount, &post.Joined_at)
	if err != nil {
		return shareddata.Post{}, row.Err()
	}
	if id != 0 {
		post.IsLiked, post.IsDisliked = d.Reactdb.CheckIfLikedPost(post.ID, user_id)
	}
	return post, row.Err()
}

func (d *PostData) GetCategoryId(name string) (int, error) {
	var id int
	err := d.Db.QueryRow("SELECT id FROM categories WHERE category_name = ?", name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (d *PostData) AddCategory(postId, categoryId int) error {
	_, err := d.Db.Exec("INSERT INTO post_category (post_id, category_id) VALUES(?, ?)", postId, categoryId)
	if err != nil {
		return err
	}

	return nil
}

func (d *PostData) DeletePost(post_id int) error {
	row, err := d.Db.Prepare("DELETE FROM post WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = row.Exec(post_id)
	if err != nil {
		return err
	}
	return nil
}

func (a *PostData) GetOwnPostsOrLikedPosts(params string, posts []shareddata.Post, userid int, num int) ([]shareddata.Post, error) {
	if params == "liked" {
		rows, err := a.Db.Query("SELECT sp.post_id, sp.post_title, sp.post_content, sp.post_date, sp.post_author, sp.post_likes, sp.post_dislikes, sp.post_comments_count FROM single_post sp INNER JOIN postReact pr ON sp.post_id = pr.post_id AND is_liked = 1 AND user_id = ?  LIMIT ? OFFSET ?", userid, PostsPerPage, num)
		if err != nil {
			return nil, err
		}
		return a.GetPostbyRows(rows, posts, userid)
	}
	return a.GetPostsByUserid(userid, "", posts, num)
}

func (a *PostData) GetPostsByUserid(userId int, categorie string, posts []shareddata.Post, num int) ([]shareddata.Post, error) {
	if categorie == "" {
		rows, err := a.Db.Query("SELECT post_id, post_title, post_content, post_date, post_author, post_likes, post_dislikes, post_comments_count FROM single_post WHERE post_author_id = ? LIMIT ? OFFSET ?", userId, PostsPerPage, num)
		if err != nil {
			return nil, err
		}

		posts, err = a.GetPostbyRows(rows, posts, userId)
		if err != nil {
			return nil, err
		}
	} else {
		rows, err := a.Db.Query(`
    SELECT 
        sp.post_id, 
        sp.post_title, 
        sp.post_content, 
        sp.post_date, 
        sp.post_author, 
        sp.post_likes, 
        sp.post_dislikes, 
        sp.post_comments_count 
    FROM 
        single_post sp 
    INNER JOIN 
        post_category pc ON sp.post_id = pc.post_id 
    INNER JOIN 
        categories c ON pc.category_id = c.id 
    WHERE 
        c.category_name = ? AND sp.post_author_id = ?
    LIMIT ? OFFSET ?`,
			categorie, userId, PostsPerPage, num)
		if err != nil {
			return nil, err
		}

		posts, err = a.GetPostbyRows(rows, posts, userId)
		if err != nil {
			return nil, err
		}
	}
	return posts, nil
}

func (a *PostData) GetBothFilters(user_id int, params, categories string, posts []shareddata.Post, num int) ([]shareddata.Post, error) {
	if params == "liked" {
		rows, err := a.Db.Query(`
			SELECT 
				sp.post_id, 
				sp.post_title, 
				sp.post_content, 
				sp.post_date, 
				sp.post_author, 
				sp.post_likes, 
				sp.post_dislikes, 
				sp.post_comments_count 
			FROM 
				single_post sp 
			INNER JOIN 
				postReact pr ON sp.post_id = pr.post_id
			INNER JOIN 
				post_category pc ON sp.post_id = pc.post_id 
			INNER JOIN 
				categories c ON pc.category_id = c.id 
			WHERE 
				pr.is_liked = 1 AND pr.user_id = ? AND c.category_name = ?
			LIMIT ? OFFSET ?`,
			user_id, categories, PostsPerPage, num)
		if err != nil {
			return nil, err
		}
		return a.GetPostbyRows(rows, posts, user_id)
	}
	return a.GetPostsByUserid(user_id, categories, posts, num)
}

func (a *PostData) GetCategoriesPost(categorie string, posts []shareddata.Post, num int, userid int) ([]shareddata.Post, error) {
	rows, err := a.Db.Query("SELECT sp.post_id, sp.post_title, sp.post_content, sp.post_date, sp.post_author, sp.post_likes, sp.post_dislikes, sp.post_comments_count FROM single_post sp INNER JOIN post_category pc ON sp.post_id = pc.post_id INNER JOIN categories c ON pc.category_id = c.id WHERE c.category_name = ? LIMIT ? OFFSET ?", categorie, PostsPerPage, num)
	if err != nil {
		return nil, err
	}
	return a.GetPostbyRows(rows, posts, userid)
}

func (a *PostData) GetPostbyRows(rows *sql.Rows, posts []shareddata.Post, userId int) ([]shareddata.Post, error) {
	var post shareddata.Post
	for rows.Next() {
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Created, &post.Author, &post.Likes, &post.Dislikes, &post.CommentsCount)
		if err != nil {
			return nil, err
		}
		post.IsLiked, post.IsDisliked = a.Reactdb.CheckIfLikedPost(post.ID, userId)
		categories, err := a.GetPostCategories(post.ID)
		if err != nil {
			return nil, err
		}
		post.Categories = categories
		posts = append(posts, post)
	}
	err := rows.Err()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (a *PostData) GetPostCategories(postId int) ([]string, error) {
	// Get Categories Ids
	var names []string
	rows, err := a.Db.Query(`
	SELECT 
    c.category_name
FROM 
    post p
JOIN 
    post_category pc ON pc.post_id = p.id
JOIN
    categories c ON c.id = pc.category_id
WHERE p.id = ?`, postId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var name string
		rows.Scan(&name)
		names = append(names, name)
	}

	return names, nil
}
