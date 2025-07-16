# 🧬 Pokémon API - Golang + Gin + Swagger

Uma API simples feita em Go que lista pokémons com base em um arquivo JSON. A API utiliza o framework [Gin](https://github.com/gin-gonic/gin) e gera documentação automática com [Swagger (Swaggo)](https://github.com/swaggo/swag).

---

## 🚀 Funcionalidades

- 📄 Lista todos os pokémons disponíveis no arquivo `pokemons.json`
- 🔍 Interface Swagger para explorar e testar a API
- 📦 Estrutura simples e modular com uso de `go mod`

---

## 📦 Instalação

### Pré-requisitos

- Go 1.23 ou superior instalado
- Git

### Clonando o projeto

```bash
git clone https://github.com/LmarDark/pokemon-api-golang.git
cd pokemon-api-golang
```

### Instalando dependências

```bash
go mod tidy
```

### Gerando a documentação Swagger

```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```

> Isso cria a pasta `/docs` automaticamente, usada para exibir a documentação.

---

## 🧪 Executando o projeto

```bash
go run main.go
```

A aplicação estará disponível em:

- API: [http://localhost/pokemons](http://localhost/pokemons)
- Swagger: [http://localhost/swagger/index.html](http://localhost/swagger/index.html)

---

## 🧾 Exemplo de `pokemons.json`

```json
{
  "Bulbasaur": {
    "image": "example_1",
    "number": "001"
  },
  "Charmander": {
    "image": "example_2",
    "number": "004"
  },
  "Squirtle": {
    "image": "example_3",
    "number": "007"
  }
}
```

---

## 🧰 Tecnologias usadas

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)
- [Swaggo/Swagger](https://github.com/swaggo/swag)
- JSON como banco de dados simples

---

## 📂 Estrutura do projeto

```
pokemon-api-golang/
├── docs/              # Documentação gerada pelo swag
├── pokemons.json      # "Banco de dados" dos pokémons
├── go.mod
├── go.sum
└── main.go
```

---

## 📝 Licença

MIT License © [Lucas Matheus (LmarDark)](https://github.com/LmarDark)
