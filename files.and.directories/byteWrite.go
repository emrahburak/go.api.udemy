package main

// byte ları dosyaya Yazma

import (
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile(
		"demo.txt",
		os.O_WRONLY,
		0666)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	byteSlice := []byte("Dosyaya yazdık\n")
	byteWritten, err := file.Write(byteSlice)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Wrotes %d bytes\n", byteWritten)

}
