package entities

type Users struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Sobrenome     string `json:"sobrenome"`
	Telefone      int    `json:"telefone"`
	Email         string `json:"email"`
	Sexo          string `json:"sexo"`
	Cpf           int    `json:"cpf"`
	Endereco      string `json:"endereco"`
	DataNascimento string `json:"dataNascimento"`
	Matricula     string `json:"matricula"`
	Tipo          string `json:"tipo"`
}

type Administrador struct {
	Users
	Permissao bool `json:"permissao"`
}

type Professor struct {
	Users
	Disciplina   string `json:"disciplina"`
	Especialidade string `json:"especialidade"`
	Turno        string `json:"turno"`
}

type Aluno struct {
	Users
	Serie int `json:"serie"`
}

type Responsavel struct {
	Users
	Aluno string `json:"aluno"`
}
