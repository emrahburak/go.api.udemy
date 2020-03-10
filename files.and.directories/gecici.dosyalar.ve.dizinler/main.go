package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	tempDirPath, err := ioutil.TempDir("", "geciciklasor")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Geçici klasör oluşturuldu: ", tempDirPath)

	//geçici dosya oluştur

	tempFile, err := ioutil.TempFile(tempDirPath, "gecicidosya.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Gecici dosya oluşturuldu: ", tempFile.Name)

	/// close file

	err = tempFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	// silme

	err = os.Remove(tempFile.Name())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s dosyası silindi.\n", tempFile.Name())
	}

	err = os.Remove(tempDirPath)
	if err != nil {
		log.Fatal(err)
	}
}
