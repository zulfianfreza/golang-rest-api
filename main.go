package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/zulfianfreza/golang-rest-api/app"
	"github.com/zulfianfreza/golang-rest-api/controller"
	"github.com/zulfianfreza/golang-rest-api/helper"
	"github.com/zulfianfreza/golang-rest-api/middleware"
	"github.com/zulfianfreza/golang-rest-api/repository"
	"github.com/zulfianfreza/golang-rest-api/service"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
