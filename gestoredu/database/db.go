package database

import (
	"fmt"

	"git.intelbras.com.br/isec/linha-future/jovens/gestoredu.git/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Db() {
	var dsn = "intelbras:Admin@tcp(10.100.61.174:3306)/db_api"
	var v = "Não conseguiu conectar ao banco de dados"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.User{})
	if err != nil {
		panic(v)
	}
	fmt.Println("conexão OK!")
	fmt.Println(db)
}
