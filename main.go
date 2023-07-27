package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tckthecreator/clean_arch_go/controller"
	"github.com/tckthecreator/clean_arch_go/repository"
	"github.com/tckthecreator/clean_arch_go/service"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_POSTGRES_CONNSTR"))
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	repo := repository.NewRepository(db)
	ser := service.NewService(repo)
	controller.NewController(gin.Default(), ser)
}
