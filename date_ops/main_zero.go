package main

import "fmt"
import "time"

// Tarih ve zaman operarsyonları

func main() {
	t := time.Date(2018, time.November, 10, 20, 0, 0, 0, time.UTC)
	fmt.Printf("Calışıyor : %s\n", t)
	fmt.Println("+++++++++++++++++")
	now := time.Now()
	fmt.Printf("Mevcut zaman  : %s\n", now)
	fmt.Println("+++++++++++++++++")
	fmt.Println("Ay :", now.Month())
	fmt.Println("Gün : ", now.Day())
	fmt.Println("Haftanın günü:", now.Weekday())
}
