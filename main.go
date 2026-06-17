package main;

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/microsoft/go-mssqldb"
	
)

type Pessoa struct {
    ID          int
    RazaoSocial string
    CNPJ        string
}

func main(){

	err := godotenv.Load()
    if err != nil {
        log.Fatal("Erro ao carregar o arquivo .env")
    }

	fmt.Println("Tentando conexão com SQL Server...")

	query := url.Values{}
	query.Add("database", "WiserSeDb-MAXI")
	query.Add("encrypt", "disable")

	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(os.Getenv("DATABASEUSER"), os.Getenv("DATABASEPASSWORD")),
		Host:     os.Getenv("DATABASEHOST") + ":1433",
		RawQuery: query.Encode(),
	}

	db, err := sql.Open("sqlserver", u.String())
	if err != nil {
		log.Fatal("Erro ao abrir conexão: ", err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao conectar: ", err.Error())
	}
	fmt.Println("Conexão com SQL Server estabelecida com sucesso!")

	rows, err := db.Query("SELECT TOP 5 Id, RazaoSocial, CNPJ FROM Pessoa WHERE CNPJ IS NOT NULL")
	if err != nil {
		log.Fatal("Erro ao executar consulta: ", err.Error())
	}

	defer rows.Close()

	var pessoas []Pessoa
	for rows.Next() {
		var p Pessoa
		err := rows.Scan(&p.ID, &p.RazaoSocial, &p.CNPJ)
		if err != nil {
			log.Fatal("Erro ao escanear linha: ", err.Error())
		}
		pessoas = append(pessoas, p)
	}
	fmt.Println("Pessoas encontradas:")
	for _, p := range pessoas {
		fmt.Printf("ID: %d, Razão Social: %s, CNPJ: %s\n", p.ID, p.RazaoSocial, p.CNPJ)
	}
}