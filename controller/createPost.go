package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/shinshin8/myFavorite/dto"
	"github.com/shinshin8/myFavorite/model"
	"github.com/shinshin8/myFavorite/utils"
)

var portConfig dto.PortConfig

// CreatePost creates a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	// listening port
	port := portConfig.Port.Port
	// Set CORS
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	w.Header().Set(utils.Cors, "http://localhost"+port)
	w.Header().Set(utils.ArrowHeader, utils.ContentType)
	w.Header().Set(utils.Credential, utils.True)

	// Session
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sessionToken := c.Value

	// Get user id from cache.
	userIDCache, err := utils.Cache.Do("GET", sessionToken)
	userID, _ := redis.Int(userIDCache, err)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if userIDCache == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// input form name
	var (
		til  = "title"
		cont = "content"
	)
	// Get title
	title := r.PostFormValue(til)
	//Get content
	content := r.PostFormValue(cont)

	// Check userID
	if !utils.IsID(userID) {
		// Invalid user id
		invalidUserID := 14
		// Set values into the struct
		resStruct := dto.NewPost{
			Status:    http.StatusOK,
			ErrorCode: invalidUserID,
			UserID:    userID,
			Title:     title,
			Content:   content,
		}
		// convert struct to JSON
		res, err := json.Marshal(resStruct)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Response JSON
		w.Write(res)
		return
	}

	// Check title
	if !utils.IsTitle(title) {
		// Invalid title
		invalidTitle := 15
		// Set values into the struct
		resStruct := dto.NewPost{
			Status:    http.StatusOK,
			ErrorCode: invalidTitle,
			UserID:    userID,
			Title:     title,
			Content:   content,
		}
		// convert struct to JSON
		res, err := json.Marshal(resStruct)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Response JSON
		w.Write(res)
		return
	}

	// Check content
	if !utils.IsContent(content) {
		// Invalid content
		invalidContent := 16
		// Set values into the struct
		resStruct := dto.NewPost{
			Status:    http.StatusOK,
			ErrorCode: invalidContent,
			UserID:    userID,
			Title:     title,
			Content:   content,
		}
		// convert struct to JSON
		res, err := json.Marshal(resStruct)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Response JSON
		w.Write(res)
		return
	}

	// Execute insert data to DB.
	result := model.CreateNewPost(userID, title, content)

	// In the Model, the function returns JSON in other way.
	// So in this part, just response result.

	// convert struct to JSON
	res, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Response JSON
	w.Write(res)

}
