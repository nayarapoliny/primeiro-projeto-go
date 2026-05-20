package main

import (
	"fmt"
	"net/http"

	// Importando nossos pacotes internos
	"github.com/nayarapoliny/primeiro-projeto-go/internal/db"
	"github.com/nayarapoliny/primeiro-projeto-go/internal/handlers"
)

func main() {
	// 1. Inicia a conexão com o banco de dados e cria a tabela se não existir
	db.Conectar()

	// 2. Registra a rota de verificação de saúde da API
	http.HandleFunc("GET /api/health", handlers.HealthCheck)
	
	// 3. Registra as rotas de Usuários separadas por Método (recurso do Go 1.22+)
	http.HandleFunc("POST /api/usuarios", handlers.CriarUsuario)   // Para Salvar no banco
	http.HandleFunc("GET /api/usuarios", handlers.ListarUsuarios)  // Para Listar do banco

	fmt.Println("🚀 Servidor rodando na porta 8080...")
	
	// 4. Inicializa o servidor HTTP
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Erro crítico no servidor: %v\n", err)
	}
}