package entities

import (
	"github.com/google/uuid"
)

type Users struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Sobrenome      string `json:"sobrenome"`
	Telefone       int    `json:"telefone"`
	Email          string `json:"email"`
	Sexo           string `json:"sexo"`
	CPF            int    `json:"cpf"`
	Endereco       string `json:"endereco"`
	DataNascimento string `json:"dataNascimento"`
	Matricula      string `json:"matricula"`
	Tipo           string `json:"tipo"`
}

//o user receberia um IDtipe esse id faria a divisao dos users 

//ou fazer com que a funçao fassa essa separação

type Administrador struct {
	Users
	Permissao bool `json:"permissao"`
}

type Professor struct {
	Users
	Disciplina    string `json:"disciplina"`
	Especialidade string `json:"especialidade"`
	Turno         string `json:"turno"`
}

type Aluno struct {
	Users
	Serie int `json:"serie"`
}

type Responsavel struct {
	Users
	Aluno string `json:"aluno"`
}

func NewUser(tipo string) interface{} {
	baseUser := Users{
		ID:   uuid.New().String(),
		Tipo: tipo,
	}

	switch tipo {
	case "administrador":
		return &Administrador{
			Users: baseUser,
		}
	case "professor":
		return &Professor{
			Users: baseUser,
		}
	case "aluno":
		return &Aluno{
			Users: baseUser,
		}
	case "responsavel":
		return &Responsavel{
			Users: baseUser,
		}
	default:
		return nil
	}
}