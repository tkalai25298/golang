package config

import (
	"os"
	"log"
	"errors"
)

//Config for vault server configuration
type Config struct {
	VaultURL string `json:"vaultURL"`
	VaultToken string `json:"VaultToken"`
}

//SetConfig to fetch the vault config from env variable
func SetConfig() (*Config,error) {

	config := Config{}

	config.VaultURL = os.Getenv("VAULT_URL")
	config.VaultToken =  os.Getenv("VAULT_ROOT_TOKEN")

	if len(config.VaultURL) == 0{
		log.Println("[ERROR]Set the VAULT_URL env variable ")
		return nil, errors.New("Set the VAULT_URL env variable")
	}

	if len(config.VaultToken) == 0{
		log.Println("[ERROR] Set the VAULT_ROOT_TOKEN env variable ")
		return nil, errors.New("Set the VAULT_ROOT_TOKEN env variable")
	}

	return &config,nil

}