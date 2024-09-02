package models

import (
	"database/sql"
	"vrs/utils"
)

type Curso struct {
	Codigo    int     `json:"codigo"`
	Descricao string  `json:"descricao"`
	Ementa    string  `json:"ementa"`
	Alunos    []Aluno `json:"alunos"`
}

func (c *Curso) Create() error {
	sql := "INSERT INTO curso (descricao, ementa) VALUES ($1, $2)"
	_, err := utils.DB.Exec(sql, c.Descricao, c.Ementa)
	if err != nil {
		utils.CreateFileDay(utils.Message{File: "Curso.Create()", Error: err.Error()})
		return err
	}
	return err
}

func (c *Curso) Update() error {
	sql := "UPDATE curso SET descricao = $1, ementa = $2 WHERE codigo = $3"
	_, err := utils.DB.Exec(sql, c.Descricao, c.Ementa, c.Codigo)
	if err != nil {
		utils.CreateFileDay(utils.Message{File: "Curso.Update()", Error: err.Error()})
		return err
	}
	return err
}

func (c *Curso) Delete() bool {
	sql := "DELETE FROM curso WHERE codigo = $1"
	_, err := utils.DB.Exec(sql, c.Codigo)
	if err != nil {
		utils.CreateFileDay(utils.Message{File: "Curso.Delete()", Error: err.Error()})
		return false
	}
	return true
}

func (c *Curso) FindAll() ([]Curso, error) {
	var cursos []Curso
	var cursoMap = make(map[int]*Curso)

	// Consulta SQL para obter cursos e alunos associados
	query := `
		SELECT c.codigo, c.descricao, c.ementa, a.codigo AS aluno_codigo, a.nome AS aluno_nome
		FROM curso c
		LEFT JOIN curso_aluno ca ON c.codigo = ca.codigo_curso
		LEFT JOIN aluno a ON ca.codigo_aluno = a.codigo
	`
	rows, err := utils.DB.Query(query)
	if err != nil {
		utils.CreateFileDay(utils.Message{File: "Curso.FindAll()", Error: err.Error()})
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var curso Curso
		var alunoCodigo sql.NullInt32
		var alunoNome sql.NullString

		if err := rows.Scan(&curso.Codigo, &curso.Descricao, &curso.Ementa, &alunoCodigo, &alunoNome); err != nil {
			utils.CreateFileDay(utils.Message{File: "Curso.FindAll()", Error: err.Error()})
			return nil, err
		}

		// Adiciona o curso ao mapa se ainda não estiver lá
		if _, exists := cursoMap[curso.Codigo]; !exists {
			cursoMap[curso.Codigo] = &curso
		}

		// Adiciona o aluno ao curso
		if alunoCodigo.Valid {
			aluno := Aluno{
				Codigo: int(alunoCodigo.Int32),
				Nome:   alunoNome.String,
			}
			cursoMap[curso.Codigo].Alunos = append(cursoMap[curso.Codigo].Alunos, aluno)
		}
	}

	// Converte o mapa para uma lista
	for _, curso := range cursoMap {
		cursos = append(cursos, *curso)
	}

	return cursos, nil
}

func (c *Curso) FindById() error {
	sql := "SELECT descricao, ementa FROM curso WHERE codigo = $1"
	row := utils.DB.QueryRow(sql, c.Codigo)
	err := row.Scan(&c.Descricao, &c.Ementa)
	if err != nil {
		utils.CreateFileDay(utils.Message{File: "Curso.FindById()", Error: err.Error()})
		return err
	}
	return nil
}
