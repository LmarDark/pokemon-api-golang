// @title API de Pokémons
// @version 1.0
// @description Esta API serve para listar Pokémons
// @termsOfService http://swagger.io/terms/

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:80
// @BasePath /
package main

import (
    "github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
	"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"

	_ "github.com/LmarDark/pokemon-api-golang/docs" 
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

// @Summary Lista todos os pokémons
// @Produce json
// @Success 200 {object} map[string]Pokemon
// @Router /pokemons [get]
func getPokemons(c *gin.Context) {
	c.JSON(http.StatusOK, db)
}

func setupRouter() *gin.Engine {
	err := loadPokemonsFromFile("pokemons.json")
	if err != nil {
		log.Fatalf("Error ao carregar pokémons: $v", err)
	}	

	r := gin.Default()
	
	r.GET("/pokemons", getPokemons)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func main() {
	r := setupRouter()

	r.Run(":80")
}
