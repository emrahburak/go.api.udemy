package main

// Yeniden isimlendirme ve taşıma
import (
	"log"
	"os"
)

func main() {
	originalPath := "demo.txt"
	newPath := "./moved/test.txt"
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
