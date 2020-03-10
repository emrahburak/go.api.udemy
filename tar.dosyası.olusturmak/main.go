package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
)

/// Dosyalari tar dosyasi olarak sikistirmak
//// Tar dosyasi üretmek

var fileFolderPath = "files/"

var files = []string{fileFolderPath + "demo.go", fileFolderPath + "note1.txt"}

func addFile(fileName string, tw *tar.Writer) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Dosya açılırken bir hata meydana geldi (%s): %s", fileName, err)

	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Dosya bilgileri alınırken bir hata meydana geldi (%s): %s", fileName, err)

	}

	hdr := &tar.Header{
		ModTime: stat.ModTime(),
		Name:    fileName,
		Size:    stat.Size(),
		Mode:    int64(stat.Mode().Perm()),
	}

	if err := tw.WriteHeader(hdr); err != nil {
		msg := "TAR header yazılırken bir hata meydana geldi %s: %s"
		return fmt.Errorf(msg, fileName, err)
	}

	copied, err := io.Copy(tw, file)
	if err != nil {
		return fmt.Errorf("%s dosyası TAR'a yazılarıken hata meydana geldi: %s", fileName, err)
	}

	if copied < stat.Size() {
		msg := "%s dosyasına %d kadar veri yazıldı . Ancak hedef beklenen %d kadardı"
		return fmt.Errorf(msg, copied, stat.Size())
	}

	return nil

}

func createArchiveTARFile(archiveFileName string) int {
	if len(archiveFileName) == 0 {
		return -1
	}

	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(archiveFileName+".tar", flags, 0644)
	if err != nil {
		log.Fatalf("TAR dosyasına veri yazmak için açılırken Hata meydana geldi : %s", err)
		return -1
	}
	defer file.Close()

	tw := tar.NewWriter(file)
	defer tw.Close()

	for _, fileName := range files {
		if err := addFile(fileName, tw); err != nil {
			log.Fatalf("%s dosyası  TAR dosyasına eklenirken hata meydana geldi : %s", fileName, err)
		}

	}
	return 1
}

func main() {
	result := createArchiveTARFile("dosyaX")
	if result > 0 {
		fmt.Println("işlem başarılı :", result)
	} else {
		fmt.Println("İşlem başarısız : ", result)
	}
}
