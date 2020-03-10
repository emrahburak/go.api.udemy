package main

// copy  a file

import (
	"io"
	"log"
	"os"
)

func main() {
	originalFile, err := os.Open("demo.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer originalFile.Close()

	// Yeni bir dosya oluştur.

	newFile, err := os.Create("./folder/demo_copy.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer newFile.Close()
	// Byte' ları kaynaktan hedef kopyala.

	byteWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Copied %d bytes: ", byteWritten)

	// Dosya içeriğini işle

	// Beleği diske boşalt

	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}

}
