package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CriarUsuario insere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario modelos.Usuario

	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	usuarioId, erro := repositorio.Criar(usuario)
	if erro != nil {
		log.Fatal(erro)
	}
	w.Write([]byte(fmt.Sprintf("Id inserido %d", usuarioId)))
}

// BuscarUsuarios busca todos os usuario no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os Usuarios!"))
}

// BuscarUsuario busca um usuario no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um Usuario!"))
}

// AtualizandoUsuario altera as informações de um usuario do banco
func AtualizandoUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuario!"))
}

// DeletandoUsuario deleta um usuario do banco
func DeletandoUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuario!"))
}
