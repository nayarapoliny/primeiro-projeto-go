package handlers

import (
	"encoding/json"
	"net/http"
)

// RespostaAPI é a estrutura base para nossas respostas
type RespostaAPI struct {
	Mensagem string `json:"mensagem"`
	Status   string `json:"status"`
}

// HealthCheck verifica se a API está no ar
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resposta := RespostaAPI{
		Mensagem: "API desenvolvida em Go rodando com sucesso!",
		Status:   "ok",
	}

	json.NewEncoder(w).Encode(resposta)
}