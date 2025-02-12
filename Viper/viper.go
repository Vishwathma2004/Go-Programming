package main

import (
	"flag"
	"fmt"
	"log"
	"github.com/spf13/viper"
)

func useCredential(username, password string) {
	fmt.Printf("Connecting to %s at %s\n",viper.GetString("webserver.name"),viper.GetString(("webserver.endpoint")))
	// Print the username and password correctly using format specifiers.
	fmt.Printf("Your username is %s and your password is %s\n", username, password)
}

func main() {
	// Declare username and password as string variables
	var username, password string

	// Use flag.StringVar to associate the command line flags with the variables
	// Change flag name to -username for username
	flag.StringVar(&username, "username", "", "username to use")
	flag.StringVar(&password, "password", "", "password to use")
	flag.Parse()
	viper.SetDefault("webserver.endpoint", "localhost:8080")
	viper.SetDefault("webserver.name","Testing Endpoint")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
	// Check if username or password is empty
	if username == "" || password == "" {
		log.Fatalln("Must pass credential")
	}
	viper.Set("Credential username",username)
	viper.Set("Credential password",password)
	viper.WriteConfig()

	// Call the useCredential function with the provided username and password
	useCredential(username, password)
}
