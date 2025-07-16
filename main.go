package main

import (
    "encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

type Pokemon struct {
    Image  string `json:"image"`
    Number string `json:"number"`
}

var db map[string]Pokemon

func loadPokemonsFromFile(filename string) error {
    data, err := ioutil.ReadFile(filename)
	if err != nil {
        return err
    }	

	err = json.Unmarshal(data, &db)
    if err != nil {
        return err
    }
    
    return nil
}

func setupRouter() *gin.Engine {
	err := loadPokemonsFromFile("pokemons.json")
	if err != nil {
		log.Fatalf("Error ao carregar pokémons: $v", err)
	}	

	r := gin.Default()
	
	r.GET("/pokemons", func(c *gin.Context) { 
		c.JSON(http.StatusOK, db)
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
