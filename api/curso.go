package api

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"vrs/models"
	"vrs/utils"
)

func CreateCurso(c *fiber.Ctx) error {
	body := c.Body()
	course := models.Curso{}
	if err := json.Unmarshal(body, &course); err != nil {
		utils.CreateFileDay(utils.Message{File: "CreateCurso()", Error: err.Error()})
		return c.SendString("BadRequest")
	}
	if course.Descricao == "" || course.Ementa == "" {
		utils.CreateFileDay(utils.Message{File: "CreateCurso()", Error: "Parametros obrigatorios nao informados"})
		return c.SendString("BadRequest")
	}

	if err := course.Create(); err != nil {
		utils.CreateFileDay(utils.Message{File: "CreateCurso()", Error: err.Error()})
		return c.SendString("Error")
	}
	return c.JSON(course)
}

func GetCurso(c *fiber.Ctx) error {
	body := c.Body()
	course := models.Curso{}
	if err := json.Unmarshal(body, &course); err != nil {
		utils.CreateFileDay(utils.Message{File: "GetCurso()", Error: err.Error()})
		return c.SendString("BadRequest")
	}
	if course.Codigo == 0 {
		utils.CreateFileDay(utils.Message{File: "GetCurso()", Error: "Parametros obrigatorios nao informados"})
		return c.SendString("BadRequest")
	}

	if err := course.FindById(); err != nil {
		utils.CreateFileDay(utils.Message{File: "GetCurso()", Error: err.Error()})
		return c.SendString("Error")
	}

	return c.JSON(course)
}

func GetCourses(c *fiber.Ctx) error {
	courses := models.Curso{}
	list, err := courses.FindAll()
	if err != nil {
		utils.CreateFileDay(utils.Message{File: "GetCursos()", Error: err.Error()})
		return c.SendString("Error")
	}
	return c.JSON(list)
}

func UpdateCurso(c *fiber.Ctx) error {
	body := c.Body()
	course := models.Curso{}
	if err := json.Unmarshal(body, &course); err != nil {
		utils.CreateFileDay(utils.Message{File: "UpdateCurso()", Error: err.Error()})
		return c.SendString("BadRequest")
	}
	if course.Codigo == 0 {
		utils.CreateFileDay(utils.Message{File: "UpdateCurso()", Error: "Parametros obrigatorios nao informados"})
		return c.SendString("BadRequest")
	}

	if err := course.Update(); err != nil {
		utils.CreateFileDay(utils.Message{File: "UpdateCurso()", Error: err.Error()})
		return c.SendString("Error")
	}
	return c.JSON(course)
}

func DeleteCurso(c *fiber.Ctx) error {
	body := c.Body()
	course := models.Curso{}
	if err := json.Unmarshal(body, &course); err != nil {
		utils.CreateFileDay(utils.Message{File: "DeleteCurso()", Error: err.Error()})
		return c.SendString("BadRequest")
	}
	if course.Codigo == 0 {
		utils.CreateFileDay(utils.Message{File: "DeleteCurso()", Error: "Parametros obrigatorios nao informados"})
		return c.SendString("BadRequest")
	}

	if err := course.Delete(); err != true {
		return c.SendString("Error")
	}

	return c.JSON(utils.Message{Info: "Curso deletado com sucesso"})
}
