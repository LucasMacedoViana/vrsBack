package models

import "vrs/utils"

type Aluno struct {
	Codigo int    `json:"codigo"`
	Nome   string `json:"nome"`
}

func (a *Aluno) Create() error {
	sql := "INSERT INTO aluno (nome) VALUES ($1)"
	_, err := utils.DB.Exec(sql, a.Nome)
	if err != nil {
		utils.CreateFileDay(utils.Message{File: "Aluno.Create()", Error: err.Error()})
		return err
	}
	return err
}

func (a *Aluno) Update() error {
	sql := "UPDATE aluno SET nome = $1 WHERE codigo = $2"
	_, err := utils.DB.Exec(sql, a.Nome, a.Codigo)
	if err != nil {
		utils.CreateFileDay(utils.Message{File: "Aluno.Update()", Error: err.Error()})
		return err
	}
	return err
}

func (a *Aluno) Delete() bool {
	sql := "DELETE FROM aluno WHERE codigo = $1"
	_, err := utils.DB.Exec(sql, a.Codigo)
	if err != nil {
		utils.CreateFileDay(utils.Message{File: "Aluno.Delete()", Error: err.Error()})
		return false
	}
	return true
}

func (a *Aluno) FindAll() (interface{}, error) {
	var list []Aluno
	sql := "SELECT codigo, nome FROM aluno"
	rows, err := utils.DB.Query(sql)
	if err != nil {
		utils.CreateFileDay(utils.Message{File: "Aluno.FindAll()", Error: err.Error()})
		return nil, err
	}
	for rows.Next() {
		var aluno Aluno
		rows.Scan(&aluno.Codigo, &aluno.Nome)
		list = append(list, aluno)
	}
	return list, nil
}

func (a *Aluno) FindById() error {
	sql := "SELECT nome FROM aluno WHERE codigo = $1"
	row := utils.DB.QueryRow(sql, a.Codigo)
	err := row.Scan(&a.Nome)
	if err != nil {
		utils.CreateFileDay(utils.Message{File: "Aluno.FindById()", Error: err.Error()})
		return err
	}
	return nil
}
