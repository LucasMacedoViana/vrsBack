package api

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"vrs/models"
	"vrs/utils"
)

func CreateAluno(c *fiber.Ctx) error {
	body := c.Body()
	student := models.Aluno{}
	if err := json.Unmarshal(body, &student); err != nil {
		utils.CreateFileDay(utils.Message{File: "CreateAluno()", Error: err.Error()})
		return c.SendString("BadRequest")
	}
	if student.Nome == "" {
		utils.CreateFileDay(utils.Message{File: "CreateAluno()", Error: "Parametros obrigatorios nao informados"})
		return c.SendString("BadRequest")
	}

	if err := student.Create(); err != nil {
		utils.CreateFileDay(utils.Message{File: "CreateAluno()", Error: err.Error()})
		return c.SendString("Error")
	}
	return c.JSON(student)
}

func GetAluno(c *fiber.Ctx) error {
	body := c.Body()
	student := models.Aluno{}
	if err := json.Unmarshal(body, &student); err != nil {
		utils.CreateFileDay(utils.Message{File: "GetAluno()", Error: err.Error()})
		return c.SendString("BadRequest")
	}
	if student.Codigo == 0 {
		utils.CreateFileDay(utils.Message{File: "GetAluno()", Error: "Parametros obrigatorios nao informados"})
		return c.SendString("BadRequest")
	}

	if err := student.FindById(); err != nil {
		utils.CreateFileDay(utils.Message{File: "GetAluno()", Error: err.Error()})
		return c.SendString("Error")
	}

	return c.JSON(student)
}

func GetAlunos(c *fiber.Ctx) error {
	students := models.Aluno{}
	list, err := students.FindAll()
	if err != nil {
		utils.CreateFileDay(utils.Message{File: "GetAlunos()", Error: err.Error()})
		return c.SendString("Error")
	}
	return c.JSON(list)
}

func UpdateAluno(c *fiber.Ctx) error {
	body := c.Body()
	student := models.Aluno{}
	if err := json.Unmarshal(body, &student); err != nil {
		utils.CreateFileDay(utils.Message{File: "UpdateAluno()", Error: err.Error()})
		return c.SendString("BadRequest")
	}
	if student.Codigo == 0 {
		utils.CreateFileDay(utils.Message{File: "UpdateAluno()", Error: "Parametros obrigatorios nao informados"})
		return c.SendString("BadRequest")
	}

	if err := student.Update(); err != nil {
		utils.CreateFileDay(utils.Message{File: "UpdateAluno()", Error: err.Error()})
		return c.SendString("Error")
	}
	return c.JSON(student)
}

func DeleteAluno(c *fiber.Ctx) error {
	body := c.Body()
	student := models.Aluno{}
	if err := json.Unmarshal(body, &student); err != nil {
		utils.CreateFileDay(utils.Message{File: "DeleteAluno()", Error: err.Error()})
		return c.JSON(utils.Message{Error: "BadRequest"})
	}
	if student.Codigo == 0 {
		utils.CreateFileDay(utils.Message{File: "DeleteAluno()", Error: "Parametros obrigatorios nao informados"})
		return c.JSON(utils.Message{Error: "BadRequest"})
	}

	if err := student.Delete(); err != true {
		return c.JSON(utils.Message{Error: "Error"})
	}
	return c.JSON(utils.Message{Info: "Success"})
}
