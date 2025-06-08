package main

import (
	"fmt"
	"tutorial_one/common_components"
	connectdb "tutorial_one/connect_db"
	curldfunction "tutorial_one/curld_function"

	"github.com/gofiber/fiber/v2"
)

var album1 = []common_components.Album{}

func init() {
	fmt.Print("you are in init ")
	connectdb.Startdb()
	connectdb.PrintCollectionData()
}

func main() {
	// router := gin.Default()
	// router.GET("/albums", getAlbums)
	// router.Run("localhost:8080")
	app := fiber.New()
	app.Delete("/albums/:id", curldfunction.Delete)
	app.Get("/albums", curldfunction.GetAlbums)
	app.Patch("/update/:id", curldfunction.UpdateAlbumById)
	app.Post("/create", curldfunction.PostAlbum)
	app.Listen(":8080")
}
