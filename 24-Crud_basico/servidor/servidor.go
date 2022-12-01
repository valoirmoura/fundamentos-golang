package servidor

import (
	"encoding/json"
	"fmt"
	"fundamentos-golang/24-Crud_basico/banco"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao Ler o Corpo da Requisição!"))
		return
	}

	var usuario usuario

	if err = json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		w.Write([]byte("Erro ao Converter o usuário para Struct"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao Conectar com o Banco de Dados!"))
		return
	}
	defer db.Close()

	stmt, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte(fmt.Sprintf("Erro ao Criar Statement: %v", err)))
		return
	}
	defer stmt.Close()

	insercao, err := stmt.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		w.Write([]byte("Erro ao executar o statement!"))
		return
	}

	idInserido, err := insercao.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ao Obter o ID inserido!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))

}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao Conectar com o Banco de Dados"))
		return
	}
	defer db.Close()

	//SELECT * FROM USUARIOS
	linhas, err := db.Query("select * from usuarios")
	if err != nil {
		w.Write([]byte("Erro ao buscar Usuarios"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario

	for linhas.Next() {
		var usuario usuario

		if err := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao escanear o Usuário"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		w.Write([]byte("Erro ao Converter Usuários para JSON"))
		return
	}
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	id, err := strconv.ParseUint(parametros["id"], 10, 64)
	if err != nil {
		w.Write([]byte("Erro ao Converter o Parâmetro para Inteiro"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao Conectar Banco de Dados"))
		return
	}
	defer db.Close()

	//Select
	linha, err := db.Query("select * from usuarios where id = ?", id)
	if err != nil {
		w.Write([]byte("Erro ao Buscar Usuario no Banco de Dados"))
		return
	}
	defer linha.Close()

	var usuario usuario

	if linha.Next() {
		if err := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao Scanear Usuario"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usuario); err != nil {
		w.Write([]byte("Erro ao Converter o usuario para JSON"))
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, err := strconv.ParseUint(parametros["id"], 10, 64)
	if err != nil {
		w.Write([]byte("Erro ao Converter o Parâmetro para Inteiro"))
		return
	}

	corpoRequisicao, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Erro ao Ler o Corpo da Requisicao"))
		return
	}

	var usuario usuario
	if err := json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		w.Write([]byte("Erro ao Converter para Struct"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao Conectar Banco de Dados!"))
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if err != nil {
		w.Write([]byte("Erro ao criar o stetemente!"))
		return
	}
	defer stmt.Close()

	if _, err := stmt.Exec(usuario.Nome, usuario.Email, ID); err != nil {
		w.Write([]byte("Erro ao Atualizar o usuario!"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Usuario Atualizado: %v", usuario)))

}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, err := strconv.ParseUint(parametros["id"], 10, 64)
	if err != nil {
		w.Write([]byte("Erro ao Converter para Uint"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao Conectar Banco de Dados"))
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		w.Write([]byte("Erro ao Criar Statement"))
		return
	}
	defer stmt.Close()

	if _, err := stmt.Exec(ID); err != nil {
		w.Write([]byte("Erro ao Deletar usuario no Banco de Dados"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
