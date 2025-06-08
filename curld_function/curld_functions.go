package curldfunction

import (
	"context"
	"time"
	"tutorial_one/common_components"
	connectdb "tutorial_one/connect_db"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAlbums(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := connectdb.Collection.Find(ctx, bson.D{})
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString("Error while fetching data")
	}
	defer cursor.Close(ctx)

	var albums []common_components.Album // or common_components.Album if needed

	if err := cursor.All(ctx, &albums); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error reading data")
	}

	return c.JSON(albums)
}
func PostAlbum(c *fiber.Ctx) error {
	var newAlbum common_components.Album

	if err := c.BodyParser(&newAlbum); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	newAlbum.Id = common_components.GetId()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := connectdb.Collection.InsertOne(ctx, newAlbum)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to insert album into DB")
	}

	return c.Status(fiber.StatusCreated).JSON(newAlbum)
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"id": id}
	result, err := connectdb.Collection.DeleteOne(ctx, filter)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete")
	}
	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).SendString("Not Found")
	}
	return c.SendString("Delete SuccesFully")
}

func UpdateAlbumById(c *fiber.Ctx) error {
	id := c.Params("id")
	var updatedAlbum common_components.Album
	err := c.BodyParser(&updatedAlbum)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid RequestBody")
	}
	updatedAlbum.Id = id
	// updateAlbum(id, updatedAlbum)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{
			"title":  updatedAlbum.Title,
			"artist": updatedAlbum.Artist,
			"price":  updatedAlbum.Price,
			"id":     updatedAlbum.Id,
		},
	}
	result, err := connectdb.Collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete album")
	}
	if result.MatchedCount == 0 {
		return c.Status(fiber.StatusNotFound).SendString("Not Found")
	}
	// if result.ModifiedCount == 0 {
	// 	return c.Status(fiber.StatusInternalServerError).SendString("Failed To update")

	// }
	return c.SendString("Update succesfully")
}

// func updateAlbum(id string, updatedAlbum common_components.Album) {
// 	var newAlbum []common_components.Album
// 	for _, a := range album1 {
// 		if a.Id == id {
// 			a = updatedAlbum
// 			newAlbum = append(newAlbum, a)
// 		} else {
// 			newAlbum = append(newAlbum, a)
// 		}
// 	}
// 	album1 = newAlbum
// 	fmt.Print(album1)
// }
