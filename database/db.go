package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"git.intelbras.com.br/isec/linha-future/jovens/gestoredu.git/models"
)

// DbConfig configura a conexão com o banco de dados e retorna a instância do GORM
func DbConfig() *gorm.DB {
	var dsn = "intelbras:Admin@tcp(10.100.61.174:3306)/db_api"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Não conseguiu conectar ao banco de dados: " + err.Error())
	}
	fmt.Println("Conexão OK!")

	// Migra as estruturas de banco de dados para o banco
	db.AutoMigrate(&models.User{})
	return db
}
