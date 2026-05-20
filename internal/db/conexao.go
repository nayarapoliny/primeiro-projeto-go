package db

import (
	"database/sql"
	"fmt"
	"log"
	"os" // Pacote nativo do Go para ler variáveis do sistema operacinal

	// Importando a nova biblioteca do .env
	"github.com/joho/godotenv"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var Conexao *sql.DB

func Conectar() {
	// 1. Carrega as variáveis do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Arquivo .env não encontrado. Garantindo leitura das variáveis de sistema.")
	}

	// 2. Busca a URL de conexão direto do ambiente de forma segura
	dsn := os.Getenv("DATABASE_URL")

	// Previne que a aplicação tente conectar sem a string de conexão
	if dsn == "" {
		log.Fatal("Erro: A variável DATABASE_URL não foi encontrada no .env")
	}

	// 3. Abre a conexão usando a variável segura
	Conexao, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Erro ao abrir o banco de dados: %v", err)
	}

	err = Conexao.Ping()
	if err != nil {
		log.Fatalf("Erro ao conectar no banco de dados (Ping falhou): %v", err)
	}

	fmt.Println("📦 Conectado ao PostgreSQL com sucesso e de forma segura!")

	criarTabelas()
}

func criarTabelas() {
	query := `
	CREATE TABLE IF NOT EXISTS usuarios (
		id SERIAL PRIMARY KEY,
		nome VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := Conexao.Exec(query)
	if err != nil {
		log.Fatalf("Erro ao criar a tabela de usuários: %v", err)
	}
}