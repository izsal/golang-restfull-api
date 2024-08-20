package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/izsal/belajar-golang-restfull-api/app"
	"github.com/izsal/belajar-golang-restfull-api/controller"
	"github.com/izsal/belajar-golang-restfull-api/helper"
	"github.com/izsal/belajar-golang-restfull-api/middleware"
	"github.com/izsal/belajar-golang-restfull-api/repository"
	"github.com/izsal/belajar-golang-restfull-api/service"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
