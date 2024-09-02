package api

import (
	"github.com/gofiber/fiber/v2"
	"vrs/configs"
)

const (
	api               = "/api"
	course            = "/course"
	student           = "/student"
	cursoAluno        = "/curso_aluno"
	errParams         = "Error in parameters"
	errParamsDataBase = "Error when querying database"
	success           = "Success"
	errNotFound       = "Not found"
)

func Routes() *fiber.App {
	app := configs.ConfigsAndRandomRoutes()

	//Aluno
	app.Post(api+student, CreateAluno)
	app.Get(api+student+"/:id", GetAluno)
	app.Get(api+student, GetAlunos)
	app.Put(api+student+"/:id", UpdateAluno)
	app.Delete(api+student+"/:id", DeleteAluno)

	//Curso
	app.Post(api+course, CreateCurso)
	app.Get(api+course+"/:id", GetCurso)
	app.Get(api+course, GetCourses)
	app.Put(api+course+"/:id", UpdateCurso)
	app.Delete(api+course+"/:id", DeleteCurso)

	//CursoAluno
	app.Post(api+cursoAluno, CreateCursoAluno)
	app.Get(api+cursoAluno+"/:id", GetCursoAluno)
	app.Get(api+cursoAluno, GetCursosAluno)
	app.Put(api+cursoAluno+"/:id", UpdateCursoAluno)
	app.Delete(api+cursoAluno+"/:id", DeleteCursoAluno)

	return app
}
