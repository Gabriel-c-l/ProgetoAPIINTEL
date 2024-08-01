package models

type User struct {
	ID              string `gorm:"primaryKey"`
	Nome            string
	Sobrenome       string
	Telefone        string
	Email           string
	CPF             string
	Endereco        string
	Idade           int
	DataNascimento  string
	Matricula       string
	Tipo            string
}
