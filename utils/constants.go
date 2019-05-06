package utils

// Application form type
const (
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json"
)

// HTTP request method
const (
	Post = "POST"
	Get  = "GET"
)

// Each path
const (
	LoginPath      = "/login"
	SignInPath     = "/signIn"
	PostList       = "/"
	ShowLikedPosts = "/likedPostsList/:user_id"
)

// The directory of configuration file
const ConfigFile = "./config/development.toml"
