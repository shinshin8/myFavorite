package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/shinshin8/myFavorite_backend/controller"
	"github.com/shinshin8/myFavorite_backend/utils"
)

func main() {
	// initialize router
	r := mux.NewRouter()
	// Login
	r.HandleFunc(utils.LoginPath, controller.Login).Methods(utils.Post)
	// Sign-Up
	r.HandleFunc(utils.SignUpPath, controller.SignUp).Methods(utils.Post)
	// Trending
	r.HandleFunc(utils.Trending, controller.Trending).Methods(utils.Get)
	// Timeline
	r.HandleFunc(utils.Timeline, controller.Timeline).Methods(utils.Get)
	// Like post
	r.HandleFunc(utils.LikePost, controller.LikePost).Methods(utils.Put)
	// Delete liked post
	r.HandleFunc(utils.DeleteLikedPost, controller.DeleteLikedPost).Methods(utils.Delete)
	// Show favorite posts
	r.HandleFunc(utils.ShowFavoritePosts, controller.ShowFavoritePosts).Methods(utils.Get)
	// Create a favorite post
	r.HandleFunc(utils.FavoritePost, controller.FavoritePost).Methods(utils.Put)
	// Delete a favorite post
	r.HandleFunc(utils.DeleteFavoritePost, controller.DeleteFavoritePost).Methods(utils.Delete)
	// User's posts list
	r.HandleFunc(utils.UserPostsList, controller.UserPostsList).Methods(utils.Get)
	// Single post
	r.HandleFunc(utils.SinglePost, controller.SinglePost).Methods(utils.Get)
	// Create a post
	r.HandleFunc(utils.NewPost, controller.CreatePost).Methods(utils.Post)
	// Edit a post
	r.HandleFunc(utils.EditPost, controller.EditPost).Methods(utils.Put)
	// Delete a post
	r.HandleFunc(utils.DeletePost, controller.DeletePost).Methods(utils.Delete)
	// Show user's profile.
	r.HandleFunc(utils.ShowProfile, controller.ShowProfile).Methods(utils.Get)
	// Edit user's profile.
	r.HandleFunc(utils.EditProfile, controller.EditProfile).Methods(utils.Put)
	// Delete Account
	r.HandleFunc(utils.DeleteAccount, controller.DeleteAccount).Methods(utils.Delete)
	// Uploading icon
	r.HandleFunc(utils.UploadingIcon, controller.UploadingIcon).Methods(utils.Post)
	// Change icon
	r.HandleFunc(utils.ChangeIcon, controller.ChangeIcon).Methods(utils.Post)
	// Uploading photos
	r.HandleFunc(utils.UploadingImages, controller.UploadingImages).Methods(utils.Post)
	// Delete images
	r.HandleFunc(utils.DeleteImages, controller.DeleteImages).Methods(utils.Post)
	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedMethods:   []string{utils.Get, utils.Post, utils.Put, utils.Delete, utils.Options},
		AllowedHeaders:   []string{utils.ContentType, utils.Authorization},
		AllowedOrigins:   []string{"http://localhost:3000"},
	})

	handler := c.Handler(r)

	// listening port
	port := os.Getenv("PORT")
	// listener
	// serverError := http.ListenAndServeTLS(port, "/etc/letsencrypt/live/www.findmyfavorite.com/fullchain.pem", "/etc/letsencrypt/live/www.findmyfavorite.com/privkey.pem", handler)
	serverError := http.ListenAndServe(port, handler)
	if serverError != nil {
		log.Fatal("ListenServer: ", serverError)
	}
}
