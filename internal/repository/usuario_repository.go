package repository

import (
	"github.com/nayarapoliny/primeiro-projeto-go/internal/db"
)

// Mudamos a struct Usuario para cá, pois ela representa uma tabela do banco
type Usuario struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	CriadoEm string `json:"criado_em"`
}

// SalvarUsuario foca puramente em inserir os dados no PostgreSQL
func SalvarUsuario(nome, email string) (int, error) {
	query := `INSERT INTO usuarios (nome, email) VALUES ($1, $2) RETURNING id`
	var novoID int
	
	err := db.Conexao.QueryRow(query, nome, email).Scan(&novoID)
	return novoID, err // Retorna o ID gerado ou o erro para o Handler tratar
}

// BuscarTodosUsuarios foca puramente em fazer o SELECT no banco
func BuscarTodosUsuarios() ([]Usuario, error) {
	rows, err := db.Conexao.Query("SELECT id, nome, email, criado_em FROM usuarios ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var u Usuario
		if err := rows.Scan(&u.ID, &u.Nome, &u.Email, &u.CriadoEm); err != nil {
			continue 
		}
		usuarios = append(usuarios, u)
	}

	return usuarios, nil
}