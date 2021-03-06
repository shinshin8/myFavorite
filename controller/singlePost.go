package controller

import (
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/shinshin8/myFavorite_backend/dto"
	"github.com/shinshin8/myFavorite_backend/model"
	"github.com/shinshin8/myFavorite_backend/utils"
)

// SinglePost returns an individual post in JSON
func SinglePost(w http.ResponseWriter, r *http.Request) {
	// Set CORS
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	w.Header().Set(utils.Cors, utils.CorsWildCard)
	w.Header().Set(utils.ArrowHeader, utils.ContentType)
	w.Header().Set(utils.ArrowMethods, utils.Methods)
	w.Header().Set(utils.Credential, utils.True)
	// Get jwt from header.
	reqToken := r.Header.Get(utils.Authorization)
	// Check if jwt is verified.
	userID := utils.VerifyToken(reqToken)
	// Get article id from the URL query parameter in string type and conver it to int type.
	atlID := "article_id"
	articleIDStr := r.URL.Query().Get(atlID)
	articleID, _ := strconv.Atoi(articleIDStr)
	// Get user's single post
	singlePost := model.SinglePost(articleID)
	// Get user's images
	singleImages := model.GetSiglePostImages(articleID)
	// Get an icon
	singleIcon := model.GetIcon(singlePost.UserID)

	post := dto.SinglePostDetail{
		ArticleID:    singlePost.ArticleID,
		LikedSum:     singlePost.LikedSum,
		ImageURL:     singleImages,
		IconURL:      singleIcon,
		UserName:     singlePost.UserName,
		Title:        singlePost.Title,
		Content:      singlePost.Content,
		CreatedTime:  singlePost.CreatedTime,
		ModifiedTime: singlePost.ModifiedTime,
	}

	if userID == 0 {

		resStruct := dto.SiglePost{
			Status:      true,
			ErrorCode:   utils.SuccessCode,
			UserID:      0,
			LikedFlg:    false,
			FavoriteFlg: false,
			Post:        post,
		}

		res, err := json.Marshal(resStruct)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Response JSON
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	}
	// Get a result if this post is liked by a user.
	// this method returns the result in boolean.
	// If the result is true, it means this post is liked.
	// Otherwise, it means this post is not liked.
	likedResult := model.LikedOrNot(userID, articleID)
	// Get a result if this post is favorited by a user.
	// this method returns the result in boolean.
	// If the result is true, it means this post is favorited.
	// Otherwise, it means this post is not favorited.
	favoriteResult := model.FavoriteOrNot(userID, articleID)

	resStruct := dto.SiglePost{
		Status:      true,
		ErrorCode:   utils.SuccessCode,
		UserID:      userID,
		LikedFlg:    likedResult,
		FavoriteFlg: favoriteResult,
		Post:        post,
	}

	res, err := json.Marshal(resStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Response JSON
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	return
}
