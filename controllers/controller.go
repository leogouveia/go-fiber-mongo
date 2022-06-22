package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leogouveia/go-fiber-mongo/config"
	"github.com/leogouveia/go-fiber-mongo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Index(c *fiber.Ctx) error {
	return c.SendString("Olá mundo")
}
func GetAllRealize(c *fiber.Ctx) error {
	query := bson.D{{}}
	realizeCol := config.Mg.Db.Collection("realize")

	cursor, err := realizeCol.Find(c.Context(), query)
	if err != nil {
		return err
	}

	var resultado []models.Realize = make([]models.Realize, 0)

	if err := cursor.All(c.Context(), &resultado); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(resultado)
}

func GetRealize(c *fiber.Ctx) error {
	query := bson.D{{}}
	realizeCol := config.Mg.Db.Collection("realize")

	var realize models.Realize

	err := realizeCol.FindOne(c.Context(), query).Decode(&realize)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(realize)

}
