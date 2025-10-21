package db

import (
	"fmt"
	"log"
	"os"

	"github.com/felipecveiga/task_manager/model"
	
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORTA"),
		os.Getenv("DB_NOME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha com a conexao com o database", err)
	}

	db.AutoMigrate(&model.User{}, &model.Task{})
	fmt.Println("Conexao com o banco de dados realizada com sucesso")
	return db
}
