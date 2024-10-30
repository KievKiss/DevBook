package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioUsuarios Cria em um repositorio de usuarios
func NovoRepositorioUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuario no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO usuarios(nome, nick, email, senha)values(?,?,?,?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nome, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}
