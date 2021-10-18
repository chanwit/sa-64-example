package main

import (
	"github.com/chanwit/sa-64-example/controller"
	"github.com/chanwit/sa-64-example/entity"
	"github.com/chanwit/sa-64-example/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// User Routes
			protected.GET("/users", controller.ListUsers)
			protected.GET("/user/:id", controller.GetUser)
			protected.PATCH("/users", controller.UpdateUser)
			protected.DELETE("/users/:id", controller.DeleteUser)

			// Video Routes
			protected.GET("/videos", controller.ListVideos)
			protected.GET("/video/:id", controller.GetVideo)
			protected.POST("/videos", controller.CreateVideo)
			protected.PATCH("/videos", controller.UpdateVideo)
			protected.DELETE("/videos/:id", controller.DeleteVideo)

			// Playlist Routes
			protected.GET("/playlists", controller.ListPlaylists)
			protected.GET("/playlist/:id", controller.GetPlaylist)
			protected.GET("/playlist/watched/user/:id", controller.GetPlaylistWatchedByUser)
			protected.POST("/playlists", controller.CreatePlaylist)
			protected.PATCH("/playlists", controller.UpdatePlaylist)
			protected.DELETE("/playlists/:id", controller.DeletePlaylist)

			// Resolution Routes
			protected.GET("/resolutions", controller.ListResolutions)
			protected.GET("/resolution/:id", controller.GetResolution)
			protected.POST("/resolutions", controller.CreateResolution)
			protected.PATCH("/resolutions", controller.UpdateResolution)
			protected.DELETE("/resolutions/:id", controller.DeleteResolution)

			// WatchVideo Routes
			protected.GET("/watch_videos", controller.ListWatchVideos)
			protected.GET("/watchvideo/:id", controller.GetWatchVideo)
			protected.POST("/watch_videos", controller.CreateWatchVideo)
			protected.PATCH("/watch_videos", controller.UpdateWatchVideo)
			protected.DELETE("/watchvideors/:id", controller.DeleteWatchVideo)

		}
	}

	// User Routes
	r.POST("/users", controller.CreateUser)

	// Authentication Routes
	r.POST("/login", controller.Login)

	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
