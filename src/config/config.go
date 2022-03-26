package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL representa a URL para comunicação com a API
	APIURL    = ""
	// Porta onde a aplicação web está rodando
	Porta     = 0
	// Hashkey é utilizada para autenticar o cookie
	HashKey   []byte
	// Blockkey é utilizada para criptografar os dados do cookie
	Blockkey  []byte
)

// Carregar inicializa as variáveis de ambiente
func Carregar() {
  var erro error

	 if erro = godotenv.Load(); erro != nil {   // o Load lê o arquivo .env
		 log.Fatal(erro)
	 }

	 Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	 if erro != nil {
		 log.Fatal(erro)
	 }

	 APIURL =os.Getenv("API_URL")
	 HashKey = []byte(os.Getenv("HASH_KEY"))
	 Blockkey = []byte(os.Getenv("BLOCK_KEY"))

}