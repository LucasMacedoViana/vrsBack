package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// CreateDatabase cria um banco de dados e as tabelas necessárias
func CreateDatabase(username, password, host, port, dbname string) error {
	// Conecta ao PostgreSQL sem especificar um banco de dados específico
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s sslmode=disable", username, password, host, port)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("falha ao conectar ao PostgreSQL: %w", err)
	}
	defer db.Close()

	// Verifica se o banco de dados já existe
	exists, err := databaseExists(db, dbname)
	if err != nil {
		return fmt.Errorf("falha ao verificar a existência do banco de dados %s: %w", dbname, err)
	}

	// Se o banco de dados não existir, cria o banco de dados
	if !exists {
		if err := createDB(db, dbname); err != nil {
			return err
		}
		log.Printf("Banco de dados %s criado com sucesso!\n", dbname)
	} else {
		log.Printf("Banco de dados %s já existe!\n", dbname)
	}

	// Conecta ao banco de dados recém-criado
	dsn = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", username, password, host, port, dbname)
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("falha ao conectar ao banco de dados %s: %w", dbname, err)
	}
	defer db.Close()

	// Cria as tabelas necessárias
	if err := createTables(db); err != nil {
		return err
	}

	log.Printf("Tabelas criadas com sucesso no banco de dados %s!\n", dbname)
	return nil
}

// databaseExists verifica se um banco de dados já existe
func databaseExists(db *sql.DB, dbname string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE lower(datname) = lower($1))"
	err := db.QueryRow(query, dbname).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// createDB cria um banco de dados
func createDB(db *sql.DB, dbname string) error {
	query := fmt.Sprintf("CREATE DATABASE %s", dbname)
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("falha ao criar banco de dados %s: %w", dbname, err)
	}
	return nil
}

// createTables cria as tabelas necessárias
func createTables(db *sql.DB) error {
	tables := `
	CREATE TABLE IF NOT EXISTS curso (
		codigo SERIAL PRIMARY KEY,
		descricao VARCHAR(50),
		ementa TEXT,
		created_at TIMESTAMPTZ DEFAULT NOW(),
		updated_at TIMESTAMPTZ DEFAULT NOW(),
		deleted_at TIMESTAMPTZ
	);

	CREATE TABLE IF NOT EXISTS aluno (
		codigo SERIAL PRIMARY KEY,
		nome VARCHAR(50),
		created_at TIMESTAMPTZ DEFAULT NOW(),
		updated_at TIMESTAMPTZ DEFAULT NOW(),
		deleted_at TIMESTAMPTZ
	);

	CREATE TABLE IF NOT EXISTS curso_aluno (
		codigo SERIAL PRIMARY KEY,
		codigo_aluno INT REFERENCES aluno(codigo) ON DELETE CASCADE,
		codigo_curso INT REFERENCES curso(codigo) ON DELETE CASCADE,
		created_at TIMESTAMPTZ DEFAULT NOW(),
		updated_at TIMESTAMPTZ DEFAULT NOW(),
		deleted_at TIMESTAMPTZ
	);
	`

	_, err := db.Exec(tables)
	if err != nil {
		return fmt.Errorf("falha ao criar tabelas: %w", err)
	}
	return nil
}
