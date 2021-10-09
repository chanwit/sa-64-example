package main

import (
  "github.com/chanwit/sa-64-example/controller"
  "github.com/chanwit/sa-64-example/entity"
  "github.com/gin-gonic/gin"
)
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

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	// User Routes
	r.GET("/users", controller.ListUsers)
	r.GET("/user/:id", controller.GetUser)
	r.POST("/users", controller.CreateUser)
	r.PATCH("/users", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)

	// Video Routes
	r.GET("/videos", controller.ListVideos)
	r.GET("/video/:id", controller.GetVideo)
	r.POST("/videos", controller.CreateVideo)
	r.PATCH("/videos", controller.UpdateVideo)
	r.DELETE("/videos/:id", controller.DeleteVideo)

	// Playlist Routes
	r.GET("/playlists", controller.ListPlaylists)
	r.GET("/playlist/:id", controller.GetPlaylist)
	r.POST("/playlists", controller.CreatePlaylist)
	r.PATCH("/playlists", controller.UpdatePlaylist)
	r.DELETE("/playlists/:id", controller.DeletePlaylist)

  	// Resolution Routes
	r.GET("/resolutions", controller.ListResolutions)
	r.GET("/resolution/:id", controller.GetResolution)
	r.POST("/resolutions", controller.CreateResolution)
	r.PATCH("/resolutions", controller.UpdateResolution)
	r.DELETE("/resolutions/:id", controller.DeleteResolution)

  	// WatchVideo Routes
	r.GET("/watchvideos", controller.ListWatchVideos)
	r.GET("/watchvideo/:id", controller.GetWatchVideo)
	r.POST("/watchvideos", controller.CreateWatchVideo)
	r.PATCH("/watchvideos", controller.UpdateWatchVideo)
	r.DELETE("/watchvideors/:id", controller.DeleteWatchVideo)

	// Run the server
	r.Run()
}
