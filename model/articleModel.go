package model

import (
	"log"

	"github.com/shinshin8/myFavorite/dto"
	"github.com/shinshin8/myFavorite/utils"
)

// GetPosts is a function that returns an array which includes DB results with JSON format.
func GetPosts() []dto.Posts {
	// Initalize DB Connection
	sql := utils.DBInit()
	// Close DB connection at the end.
	defer sql.Close()
	// SQL syntax
	getPosts := `SELECT 
					article_table.article_id, 
					user_table.user_name, 
					article_table.title, 
					article_table.content,
					article_table.created_time, 
					article_table.modified_time 
				FROM 
					article_table 
				INNER JOIN 
					user_table 
				ON 
					article_table.user_id = user_table.user_id 
				ORDER BY 
					article_table.created_time DESC`

	row, err := sql.Query(getPosts)

	if err != nil {
		log.Fatal(err)
	}

	// Prepare an array which save JSON results.
	var postArray []dto.Posts

	for row.Next() {
		posts := dto.Posts{}
		if err := row.Scan(&posts.ArticleID, &posts.UserName, &posts.Title, &posts.Content, &posts.CreatedTime, &posts.ModifiedTime); err != nil {
			log.Fatal(err)
		}

		// Appending JSON in array.
		postArray = append(postArray, posts)
	}

	return postArray
}

// UserPostsList gets a specific user's posts list from DB and convert its result into JSON.
// At the parameter, user id will be put in and its type is int.
func UserPostsList(userID int) []dto.Posts {
	// Initalize DB Connection
	sql := utils.DBInit()
	// Close DB connection at the end.
	defer sql.Close()
	// SQL syntax
	getPosts := `SELECT 
					article_table.article_id, 
					user_table.user_name, 
					article_table.title, 
					article_table.content,
					article_table.created_time, 
					article_table.modified_time 
				FROM 
					article_table 
				INNER JOIN 
					user_table 
				ON 
					article_table.user_id = user_table.user_id
				WHERE
					user_table.user_id = ?
				ORDER BY 
					article_table.created_time DESC`

	row, err := sql.Query(getPosts, userID)

	if err != nil {
		log.Fatal(err)
	}

	// Prepare an array which save JSON results.
	var postArray []dto.Posts

	for row.Next() {
		posts := dto.Posts{}
		if err := row.Scan(&posts.ArticleID, &posts.UserName, &posts.Title, &posts.Content, &posts.CreatedTime, &posts.ModifiedTime); err != nil {
			log.Fatal(err)
		}

		// Appending JSON in array.
		postArray = append(postArray, posts)
	}

	return postArray
}

// LikedOrNot returns the result if a post selected is liked.
// At the first parameter, user id will be set with int type.
// At the second parameter, article id will be set with int type.
func LikedOrNot(userID int, articleID int) bool {
	// Initalize DB Connection
	sql := utils.DBInit()
	// Close DB connection at the end.
	defer sql.Close()
	// SQL syntax
	checkLikedOrNot := `SELECT IF
							(COUNT(*),'true','false') 
						AS 
							liked_flg 
						FROM 
							liked_table 
						WHERE 
							user_id = ? AND article_id = ?`

	var likedFlg bool
	err := sql.QueryRow(checkLikedOrNot, userID, articleID).Scan(&likedFlg)
	if err != nil {
		log.Fatal(err)
	}
	return likedFlg
}

// FavoriteOrNot returns the result if a post selected is liked.
// At the first parameter, user id will be set with int type.
// At the second parameter, article id will be set with int type.
func FavoriteOrNot(userID int, articleID int) bool {
	// Initalize DB Connection
	sql := utils.DBInit()
	// Close DB connection at the end.
	defer sql.Close()
	// SQL syntax
	checkFavoriteOrNot := `SELECT IF
							(COUNT(*),'true','false') 
						AS 
							favorite_flg 
						FROM 
							favorite_table 
						WHERE 
							user_id = ? AND article_id = ?`

	var favoriteFlg bool
	err := sql.QueryRow(checkFavoriteOrNot, userID, articleID).Scan(&favoriteFlg)
	if err != nil {
		log.Fatal(err)
	}
	return favoriteFlg
}

// SinglePost returns the result of a single post in JSON.
// At the first parameter, user id will be set with int type.
// At the second parameter, article id will be set with int type.
func SinglePost(userID int, articleID int) dto.Posts {
	// Initalize DB Connection
	sql := utils.DBInit()
	// Close DB connection at the end.
	defer sql.Close()
	// SQL syntax
	signlePost := `SELECT 
						article_table.article_id, 
						user_table.user_name, 
						article_table.title, 
						article_table.content, 
						article_table.created_time, 
						article_table.modified_time 
					FROM 
						article_table 
					INNER JOIN 
						user_table 
					ON 
						article_table.user_id = user_table.user_id 
					WHERE 
						article_table.user_id = ? AND article_table.article_id = ?`

	var post dto.Posts

	err := sql.QueryRow(signlePost, userID, articleID).Scan(&post.ArticleID, &post.UserName, &post.Title, &post.Content, &post.CreatedTime, &post.ModifiedTime)

	if err != nil {
		log.Fatal(err)
	}

	return post
}

// CreateNewPost inserts a new post data in DB.
// At first parameter, user id is set in int type.
// At second paramter, title is set in string type.
// At third parameter, content is set in string type.
func CreateNewPost(userID int, title string, content string) dto.SimpleResutlJSON {
	// Initalize DB Connection
	sql := utils.DBInit()
	// Close DB connection at the end.
	defer sql.Close()
	// SQL syntax
	insertNewPost := `INSERT INTO 
							article_table(
								user_id, 
								title, 
								content) 
						VALUES(?,?,?);`
	rows, err := sql.Prepare(insertNewPost)
	if err != nil {
		log.Fatal(err)
		sqlErrorStatus := 8
		res := dto.SimpleResutlJSON{
			Status:    false,
			ErrorCode: sqlErrorStatus,
		}
		return res
	}
	rows.Exec(userID, title, content)
	successStatus := 0

	res := dto.SimpleResutlJSON{
		Status:    true,
		ErrorCode: successStatus,
	}
	return res
}

// EditPost updates specific post data and return the result in JSON.
// At first parameter, user id is set in int type.
// At second paramter, article id is set in int type.
// At third parameter, title is set in string type.
// At forth parameter, content is set in string type.
func EditPost(userID int, articleID int, title string, content string) dto.SimpleResutlJSON {
	// Initalize DB Connection
	sql := utils.DBInit()
	// Close DB connection at the end.
	defer sql.Close()
	// Update sql syntax
	update := `UPDATE 
					article_table 
				SET 
					title = ?, 
					content = ? 
				WHERE 
					user_id = ? 
				AND 
					article_id = ?`
	rows, err := sql.Prepare(update)
	if err != nil {
		log.Fatal(err)
		sqlErrorStatus := 8
		res := dto.SimpleResutlJSON{
			Status:    false,
			ErrorCode: sqlErrorStatus,
		}
		return res
	}
	rows.Exec(title, content, userID, articleID)
	successStatus := 0

	res := dto.SimpleResutlJSON{
		Status:    true,
		ErrorCode: successStatus,
	}
	return res
}

// DeletePost deletes specific post data and return the result in JSON.
// At first parameter, user id is set in int type.
// At second paramter, article id is set in int type.
func DeletePost(userID int, articleID int) dto.SimpleResutlJSON {
	// Initalize DB Connection
	sql := utils.DBInit()
	// Close DB connection at the end.
	defer sql.Close()
	// delete sql syntax
	deleteSQL := `DELETE FROM 
						article_table 
					WHERE 
						user_id = ? 
					AND 
						article_id = ?`
	rows, err := sql.Prepare(deleteSQL)
	if err != nil {
		log.Fatal(err)
		sqlErrorStatus := 8
		res := dto.SimpleResutlJSON{
			Status:    false,
			ErrorCode: sqlErrorStatus,
		}
		return res
	}
	rows.Exec(userID, articleID)
	successStatus := 0

	res := dto.SimpleResutlJSON{
		Status:    true,
		ErrorCode: successStatus,
	}
	return res
}
