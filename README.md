# 🚀 API em Go com PostgreSQL e Repository Pattern

Uma API RESTful robusta e de alta performance desenvolvida em Go (Golang). Este projeto foi construído do zero com foco em boas práticas de engenharia de software, utilizando a biblioteca padrão da linguagem para roteamento e implementando uma arquitetura limpa e escalável.

## ✨ Funcionalidades e Arquitetura

* **Roteamento Nativo:** Utilização do padrão de rotas moderno do Go (1.22+), sem dependência de frameworks web pesados.
* **Repository Pattern:** Separação clara de responsabilidades entre os controladores (Handlers) e a camada de dados (Repository).
* **Validação Estrutural:** Proteção de rotas com `go-playground/validator` para garantir a integridade dos dados na entrada.
* **Segurança de Credenciais:** Gerenciamento de variáveis de ambiente com `.env` utilizando `godotenv`.
* **Banco de Dados Isolado:** Infraestrutura de dados rodando em contêineres utilizando Docker Compose.

## 🛠️ Tecnologias Utilizadas

* **Linguagem:** Go
* **Banco de Dados:** PostgreSQL
* **Driver de Banco:** `pgx` (alta performance)
* **Infraestrutura:** Docker & Docker Compose
* **Testes de API:** Insomnia

## 📂 Estrutura de Pastas

A arquitetura do projeto segue o *Standard Go Project Layout*:

```text
/
├── cmd/
│   └── api/
│       └── main.go              # Ponto de entrada da aplicação
├── internal/
│   ├── db/
│   │   └── conexao.go           # Gerenciamento da conexão com o banco e inicialização
│   ├── handlers/
│   │   ├── health.go            # Rota de verificação de saúde da API
│   │   └── usuario.go           # Controladores de requisições web (JSON/HTTP)
│   └── repository/
│       └── usuario_repository.go # Lógica de comunicação direta com o PostgreSQL (SQL)
├── .env                         # Variáveis de ambiente (não versionado)
├── docker-compose.yml           # Configuração do banco de dados
├── go.mod                       # Gerenciador de dependências
└── README.md


## 🚀 Como Rodar o Projeto

### Pré-requisitos
* Go (versão 1.22 ou superior)
* Docker e Docker Desktop com integração WSL (para Windows)

### 1. Clonar o repositório
```bash
git clone [https://github.com/nayarapoliny/primeiro-projeto-go.git](https://github.com/nayarapoliny/primeiro-projeto-go.git)
cd primeiro-projeto-go

### 2. Configurar Variáveis de Ambiente
Crie um arquivo chamado .env na raiz do projeto e adicione a string de conexão do banco de dados:
```bash
DATABASE_URL=postgres://root:secretpassword@localhost:5432/primeiro_projeto_go?sslmode=disable

### 3. Subir o Banco de Dados
Inicie o contêiner do PostgreSQL em segundo plano:
```bash
docker-compose up -d

### 4. Baixar Dependências
Sincronize as bibliotecas do projeto:
```bash
go mod tidy

### 5. Iniciar o Servidor
Execute a aplicação. O banco de dados e as tabelas serão sincronizados automaticamente:
```bash
go run cmd/api/main.go

O servidor estará rodando em http://localhost:8080.

## 🧪 Testando as Rotas (via Insomnia)
###1. Health Check
Método: GET

URL: http://localhost:8080/api/health

###2. Criar Usuário
Método: POST

URL: http://localhost:8080/api/usuarios

Body (JSON):
```bash
{
	"nome": "Teste",
	"email": "teste@email.com"
}

###3. Listar Usuários
Método: GET

URL: http://localhost:8080/api/usuarios

Desenvolvido com 💜 por Nay .