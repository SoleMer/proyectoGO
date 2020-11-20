package chat

import (
	"flag"
	"fmt"
	"os"

	"github.com/SoleMer/proyectoGO/internal/config/config.go"
)

func main() {

	configFile := flag.String("config", "./config.yaml", "this is de service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}

	fmt.Println(config.DB.Driver)
	fmt.Println(config.Version)
}
