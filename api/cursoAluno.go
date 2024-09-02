package api

import "github.com/gofiber/fiber/v2"

func CreateCursoAluno(c *fiber.Ctx) error {
	return c.SendString("CreateCursoAluno")
}

func GetCursoAluno(c *fiber.Ctx) error {
	return c.SendString("GetCursoAluno")
}

func GetCursosAluno(c *fiber.Ctx) error {
	return c.SendString("GetCursosAluno")
}

func UpdateCursoAluno(c *fiber.Ctx) error {
	return c.SendString("UpdateCursoAluno")
}

func DeleteCursoAluno(c *fiber.Ctx) error {
	return c.SendString("DeleteCursoAluno")
}
