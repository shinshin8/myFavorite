package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../dto"
	"../model"
	"../utils"
)

// EditPost edits a existing post.
func EditPost(w http.ResponseWriter, r *http.Request) {
	// Get article id from URL query parameter and convert its type string to int.
	atcID := "article_id"
	articleIDStr := r.URL.Query().Get(atcID)
	articleID, _ := strconv.Atoi(articleIDStr)
	// Each form
	var (
		usrID = "user_id"
		til   = "title"
		cont  = "content"
	)
	// Get user id
	userIDStr := r.PostFormValue(usrID)
	userID, _ := strconv.Atoi(userIDStr)
	// Get title
	title := r.PostFormValue(til)
	//Get content
	content := r.PostFormValue(cont)
	// Check userID
	if !utils.IsID(userID) {
		// Invalid user id
		invalidUserID := 17
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
		// Set HTTP header and defined MIME type
		w.Header().Set(utils.ContentType, utils.ApplicationJSON)
		// Response JSON
		w.Write(res)
		return
	}

	// Check title
	if !utils.IsTitle(title) {
		// Invalid title
		invalidTitle := 18
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
		// Set HTTP header and defined MIME type
		w.Header().Set(utils.ContentType, utils.ApplicationJSON)
		// Response JSON
		w.Write(res)
		return
	}

	// Check content
	if !utils.IsContent(content) {
		// Invalid content
		invalidContent := 19
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
		// Set HTTP header and defined MIME type
		w.Header().Set(utils.ContentType, utils.ApplicationJSON)
		// Response JSON
		w.Write(res)
		return
	}

	// Execute update data to DB.
	result := model.EditPost(userID, articleID, title, content)

	// In the Model, the function returns JSON in other way.
	// So in this part, just response result.

	// convert struct to JSON
	res, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Set HTTP header and defined MIME type
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	// Response JSON
	w.Write(res)
}