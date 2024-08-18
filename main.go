package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/izsal/belajar-golang-restfull-api/app"
	"github.com/izsal/belajar-golang-restfull-api/controller"
	"github.com/izsal/belajar-golang-restfull-api/repository"
	"github.com/izsal/belajar-golang-restfull-api/service"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
}
