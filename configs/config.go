// Install and manage third party packages

package main

import ( 
	"fmt"
	"os"
	"log"

	toml "github.com/pelletier/go-toml"
)

// Create a Config
type Config struct {
	Login struct {
		User string
		Password string
	}
}

func main() {
	fmt.Println("\n ======= MANAGE DEPENDENCIES ========= \n")

	file, err := os.Open("config.toml")
	if err != nil {
		log.Fatalf("error: cannot open config file - %s", err)
	}
	defer file.Close()

	cfg := &Config{}
	dec := toml.NewDecoder(file)
	if err := dec.Decode(cfg); err != nil {
		log.Fatalf("error: cannot decode config file - %s", err)
	}

	fmt.Printf("%v\n", cfg)
	fmt.Println("\n ======= MANAGE DEPENDENCIES ========= \n")
}