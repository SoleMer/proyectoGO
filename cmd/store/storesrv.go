package main

import (
	"flag"
	"fmt"
	"github.com/SoleMer/dulceCaliGo/internal/config"
	"github.com/SoleMer/dulceCaliGo/internal/database"
	"github.com/SoleMer/dulceCaliGo/internal/service/store"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"os"
)

func main() {
	// go run cmd/store/storesrv.go -config ./config/config.yaml
	cfg := readConfig()

	db, err := database.NewDatabase(cfg)
	fmt.Println(db)
	defer db.Close()

	if err != nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err := createSchema(db); err != nil {
		panic(err)
	}

	service, _ := store.New(db, cfg)
	httpService := store.NewHTTPTransport(service)

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
	schema := `CREATE TABLE IF NOT EXISTS clothes (
		id integer primary key autoincrement,
		name varchar(60),
		price integer,
		stock integer);`
	fmt.Println(schema)
	//execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	//or, you can use MustExec, which panics on error
	insertItem := `INSERT INTO clothes (name, price, stock) VALUES (?, ?, ?)`
	var cItem = store.ClothingItem{
		Name: "jean",
		Price: 700,
		Stock: 4}

	db.MustExec(insertItem, cItem.Name, cItem.Price, cItem.Stock)
	return nil
}