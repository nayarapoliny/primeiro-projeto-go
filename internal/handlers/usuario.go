package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	
	// Importando o nosso novo pacote repository no lugar do db
	"github.com/nayarapoliny/primeiro-projeto-go/internal/repository"
)

type CriarUsuarioRequest struct {
	Nome  string `json:"nome" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

var validate = validator.New()

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	var req CriarUsuarioRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Erro ao processar o JSON.", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Dados inválidos: verifique o nome e e-mail."})
		return
	}

	// -------------------------------------------------------------------
	// A MÁGICA AQUI: O Handler não sabe mais o que é SQL. 
	// Ele apenas pede para o Repository salvar e aguarda a resposta.
	// -------------------------------------------------------------------
	novoID, err := repository.SalvarUsuario(req.Nome, req.Email)
	
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Erro ao salvar. Este e-mail já está cadastrado?"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"mensagem": "Usuário salvo com sucesso no banco de dados!",
		"id":       novoID,
		"nome":     req.Nome,
		"email":    req.Email,
	})
}

func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	// Pede a lista de usuários para o Repository
	usuarios, err := repository.BuscarTodosUsuarios()
	
	if err != nil {
		http.Error(w, "Erro interno ao buscar usuários", http.StatusInternalServerError)
		return
	}

	// Se a lista vier vazia, garante que retornamos um array vazio [] e não "null"
	if usuarios == nil {
		usuarios = []repository.Usuario{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuarios)
}