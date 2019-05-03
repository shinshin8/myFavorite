package controller

import (
	"encoding/json"
	"net/http"

	"../dto"
	"../model"
	"../utils"
)

// PostList is controller file for get all post with JSON format.
func PostList(w http.ResponseWriter, r *http.Request) {

	if r.Method == utils.Get {
		successfulCode := 0
		// DB result array
		dbResultArray := model.GetPosts()

		resStruct := dto.PostList{
			Status:    http.StatusOK,
			ErrorCode: successfulCode,
			Posts:     dbResultArray,
		}

		res, err := json.Marshal(resStruct)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Set HTTP header and defined MIME type
		w.Header().Set(utils.ContentType, utils.ApplicationJSON)
		// Response JSON
		w.Write(res)

	} else {
		// return simple json
		// Wrong HTTP request method
		wrongHTTPMethod := 9
		// Set values into the struct
		resStruct := dto.SimpleResutlJSON{
			Status:    http.StatusOK,
			ErrorCode: wrongHTTPMethod,
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
	}
}