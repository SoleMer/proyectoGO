package chat

import (
	"flag"
	"fmt"
	"o s"

	"github.com/SoleMer/proyectoGO/internal/config"
	"github.com/SoleMer/proyectoGO/internal/service/chat"
)

func main() {

	cfg := readConfig();
	
	db, err := database.Newdatabase(cfg)
	if err != nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
	service, _ := chat.New(db, cfg)

	for _, m := range service.findAll() {
		fmt.Println(m)
	}
}

func readConfig(){
	configFile := flag.String("config", "./config.yaml", "this is de service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}
}

func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS messages (
		id integer primaty key autoincrement,
		text varchar);`

	//execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	//or, you can use MustExec, which panics on error
	insertMessage := `INSER INTO messages (text) VALUES (?)`
	s := fmt.Sprintf("Message number %v", time.Now().Nanosecond())
	db.MustExec(insertMessage, s)
	return nil
}
