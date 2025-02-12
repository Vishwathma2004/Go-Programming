package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define a command-line flag "name"
	name := flag.String("name", "", "The name to say hello to")
	flag.Parse() // Parse the command-line flags

	// Check if the name flag is empty
	if *name == "" {
		// If the name is not provided, print usage information and exit
		fmt.Println("Must add the name to add this tool!")
		flag.Usage() // Print the usage information for the flags
		os.Exit(1)   // Exit with a non-zero status code
	}

	// If a name is provided, print a hello message
	fmt.Printf("Hello %s\n", *name)
}
