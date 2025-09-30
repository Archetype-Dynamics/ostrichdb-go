package sdk

import "ostrichdb-go/src/lib"
import "github.com/joho/godotenv"
import (
    "net/http"
    "time"
    "os"
)

//Assumes its stored in a .env file
func GetAPIKey() string{
	godotenv.Load()
	ostrichJWT:= os.Getenv("OSTRICHDB_JWT")
	return ostrichJWT
}


func NewConfigBuilder () *lib.Config{
	return &lib.Config{
		BaseURL: lib.OSTRICHDB_ADDRESS,
		Timeout: 30_000,
		ApiKey: GetAPIKey(),
	}
}

func NewClientBuilder(config *lib.Config) *lib.Client {
	return &lib.Client{
		BaseURL: config.BaseURL,
		ApiKey:  config.ApiKey,
		HTTPClient: &http.Client{
			Timeout: time.Duration(config.Timeout) * time.Millisecond,
		},
	}
}

