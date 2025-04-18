package main

import (
	"fmt"
	"log"
	"quoter_back/utils"
)

func main_salt_generator() {
	for i := 0; i < 10; i++ {
		salt, err := utils.GenerateSalt()
		if err != nil {
			log.Fatalf("Error generating salt: %v", err)
		}
		fmt.Printf("Salt %d: %s\n", i+1, salt)
	}
}