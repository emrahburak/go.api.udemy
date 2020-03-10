package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Arşivlenmiş Dosyaları Dışarı Cıkarmak

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatal(err)
		}

	}

}

func main() {

	zr, err := zip.OpenReader("zip4447.zip")
	if err != nil {
		log.Fatal(err)
	}

	defer zr.Close()

	for _, file := range zr.Reader.File {
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}

		defer zippedFile.Close()

		targetDir := "./"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)

		dirName := strings.Split(file.Name, "/")
		CreateDirIfNotExist(dirName[0])

		if file.FileInfo().IsDir() {
			log.Println("Klasör oluşturuluyor .", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())

		} else {
			log.Println("Dosya Cıkarılıyor ", file.Name)

			outFile, err := os.OpenFile(

				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)

			if err != nil {
				log.Fatal(err)
			}

			defer outFile.Close()

			_, err = io.Copy(outFile, zippedFile)

			if err != nil {
				log.Fatal(err)

			}

		}

	}

}
