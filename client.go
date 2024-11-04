package main

import (
	"log"
	"os"
)

func getAuthorFeed(author string) ([]byte, error) {
	log.Printf("getting feed for %s", author)
	return os.ReadFile("getAuthorFeed.json")
}
