package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes

	// session
	rt.router.POST("/session/", rt.wrap(rt.doLogin))

	// set
	rt.router.PUT("/settings/", rt.wrap(rt.setMyUserName))

	// follow
	rt.router.PUT("/following/:id/", rt.wrap(rt.followUser))
	rt.router.DELETE("/following/:id/", rt.wrap(rt.unfollowUser))

	// banned
	rt.router.PUT("/ban/:id/", rt.wrap(rt.banUser))
	rt.router.DELETE("/ban/:id/", rt.wrap(rt.unbanUser))

	// photo
	rt.router.POST("/photos/", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/photos/:pid/", rt.wrap(rt.deletePhoto))

	// like
	rt.router.PUT("/likes/:pid/", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/likes/:pid/", rt.wrap(rt.unlikePhoto))

	// comment
	rt.router.POST("/photos/:pid/comment/", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/photos/:pid/comment/:comment_id/", rt.wrap(rt.uncommentPhoto))

	// get
	rt.router.GET("/users/:id/", rt.wrap(rt.getUserProfile))
	rt.router.GET("/users/:id/following/", rt.wrap(rt.getFollowing))
	rt.router.GET("/users/:id/followers/", rt.wrap(rt.getFollowers))
	rt.router.GET("/users/:id/photo/", rt.wrap(rt.getPhotosList))
	rt.router.GET("/photos/:pid/", rt.wrap(rt.getPhoto))
	rt.router.GET("/photos/:pid/comment/", rt.wrap(rt.getCommentsList))
	rt.router.GET("/photos/:pid/likes/", rt.wrap(rt.getLikesList))
	rt.router.GET("/stream/", rt.wrap(rt.getMyStream))

	// search
	rt.router.GET("/search/", rt.wrap(rt.searchUserByUsername))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
