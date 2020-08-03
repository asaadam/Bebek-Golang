package main

import (
	"bebek/config"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello GOLANG!"))
}

var configuration config.Configurations

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s ", err)
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

}

func main() {
	port := fmt.Sprintf(":%d", configuration.Server.Port)
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods("GET")
	fmt.Printf("Server Run %s", port)
	log.Fatal(http.ListenAndServe(port, router))

}
