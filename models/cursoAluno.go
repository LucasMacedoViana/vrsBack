package models

import "vrs/utils"

type CursoAluno struct {
	Codigo      int `json:"codigo"`
	CodigoAluno int `json:"codigo_aluno"`
	CodigoCurso int `json:"codigo_curso"`
}

func (ca *CursoAluno) Create() error {
	sql := "INSERT INTO curso_aluno (codigo_aluno, codigo_curso) VALUES ($1, $2)"
	_, err := utils.DB.Exec(sql, ca.CodigoAluno, ca.CodigoCurso)
	if err != nil {
		utils.CreateFileDay(utils.Message{File: "CursoAluno.Create()", Error: err.Error()})
		return err
	}
	return err
}

func (ca *CursoAluno) Update() error {
	sql := "UPDATE curso_aluno SET codigo_aluno = $1, codigo_curso = $2 WHERE codigo = $3"
	_, err := utils.DB.Exec(sql, ca.CodigoAluno, ca.CodigoCurso, ca.Codigo)
	if err != nil {
		utils.CreateFileDay(utils.Message{File: "CursoAluno.Update()", Error: err.Error()})
		return err
	}
	return err
}

func (ca *CursoAluno) Delete() bool {
	sql := "DELETE FROM curso_aluno WHERE codigo = $1"
	_, err := utils.DB.Exec(sql, ca.Codigo)
	if err != nil {
		utils.CreateFileDay(utils.Message{File: "CursoAluno.Delete()", Error: err.Error()})
		return false
	}
	return true
}

func (ca *CursoAluno) FindAll() (interface{}, error) {
	var list []CursoAluno
	sql := "SELECT codigo, codigo_aluno, codigo_curso FROM curso_aluno"
	rows, err := utils.DB.Query(sql)
	if err != nil {
		utils.CreateFileDay(utils.Message{File: "CursoAluno.FindAll()", Error: err.Error()})
		return nil, err
	}
	for rows.Next() {
		var cursoAluno CursoAluno
		rows.Scan(&cursoAluno.Codigo, &cursoAluno.CodigoAluno, &cursoAluno.CodigoCurso)
		list = append(list, cursoAluno)
	}
	return list, nil
}

func (ca *CursoAluno) FindById() error {
	sql := "SELECT codigo_aluno, codigo_curso FROM curso_aluno WHERE codigo = $1"
	row := utils.DB.QueryRow(sql, ca.Codigo)
	err := row.Scan(&ca.CodigoAluno, &ca.CodigoCurso)
	if err != nil {
		utils.CreateFileDay(utils.Message{File: "CursoAluno.FindById()", Error: err.Error()})
		return err
	}
	return err
}
