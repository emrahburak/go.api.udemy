package main

// Dosya bilgisini almak
import (
	"fmt"
	"log"
	"os"
)

var (
	fileInfo *os.FileInfo
	err      error
)

func main() {
	// Dosya bilgisini döndürür
	// Hata! dosya yoksa..
	fileInfo, err := os.Stat("demo.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File name: ", fileInfo.Name())
	fmt.Println("Size in bytes: ", fileInfo.Size())
	fmt.Println("Permissions: ", fileInfo.Mode())
	fmt.Println("Last modified: ", fileInfo.ModTime())
	fmt.Println("Is Dir", fileInfo.IsDir())
	fmt.Println("System Interface Type: ", fileInfo.Sys())
	fmt.Println("System Info: %+v\n\n ", fileInfo.Sys())
}
