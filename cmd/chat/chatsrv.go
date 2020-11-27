package main

import (
	"flag"
	"fmt"
	"github.com/SoleMer/proyectoGO/internal/config"
	"github.com/SoleMer/proyectoGO/internal/database"
	"github.com/SoleMer/proyectoGO/internal/service/chat"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"os"
)

func main() {
	// go run cmd/chat/chatsrv.go -config ./config/config.yaml
	cfg := readConfig()
	
	db, err := database.NewDatabase(cfg)
	defer db.Close()

	if err != nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err := createSchema(db); err != nil {
		panic(err)
	}

	service, _ := chat.New(db, cfg)
	httpService := chat.NewHTTPTransport(service)

	r := gin.Default()
	httpService.Register(r)
	r.Run()


}

func readConfig() *config.Config{
	configFile := flag.String("config", "./config.yaml", "this is de service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return cfg
}

func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS messages (
		id integer primary key autoincrement,
		text varchar(60));`

	//execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	//or, you can use MustExec, which panics on error
	insertMessage := `INSERT INTO messages (text) VALUES (?)`
	s := fmt.Sprintf("Hola")
	db.MustExec(insertMessage, s)
	return nil
}
