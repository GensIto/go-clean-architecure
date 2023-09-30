package main

import (
	"database/sql"
	"fmt"
	"go-crean-aarchitecture/controller"
	"go-crean-aarchitecture/repository"
	"go-crean-aarchitecture/usecase"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	return db, err
}

func main() {

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	tr := repository.NewTaskRepository(db)
	tu := usecase.NewTaskUsecase(tr)
	tc := controller.NewTaskController(tu)

	e.GET("/tasks/:id", tc.Get)
	e.POST("/tasks", tc.Create)
	e.PUT("/tasks/:id", tc.Update)
	e.DELETE("/tasks/:id", tc.Delete)

	e.Logger.Fatal(e.Start(":8080"))
}
