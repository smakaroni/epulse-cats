package app

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/smakaroni/epulse-cats/pkg/cats"
	"github.com/smakaroni/epulse-cats/pkg/config"
	"log"
	"net/http"
)

type App struct {
	Router *gin.Engine
	DB     *sql.DB
}

// Init initialize the app, routes and db
func (a *App) Init(c *config.Config) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.Url, c.Port, c.Username, c.Password, c.DbName)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {

		log.Fatal(err)
	}

	a.Router = gin.Default()
	a.initRoutes()
}

func (a *App) initRoutes() {
	routes := a.Router.Group("/cats")
	routes.POST("/create", a.createCat)
	routes.GET("/list", a.getAllCats)
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

// createCat needs a json with name and description
func (a *App) createCat(ctx *gin.Context) {
	var c cats.Cat

	if err := ctx.BindJSON(&c); err != nil {
		log.Printf("invalid JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	if err := c.AddCat(a.DB); err != nil {
		log.Printf("error writing cat to db: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not save cat"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": "cat saved"})
}

// getAllCats takes no input
func (a *App) getAllCats(ctx *gin.Context) {
	catList, err := cats.GetAllCats(a.DB)
	if err != nil {
		log.Printf("error getting all cats: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get cats"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": catList})

}
