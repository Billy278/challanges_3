package main

import "fmt"

func main() {
	input := "selamat malam"
	jumlah := make(map[string]int)
	for _, v := range input {
		fmt.Printf("%c \n", v)
		jumlah[string(v)]++
	}
	fmt.Println(jumlah)
}
