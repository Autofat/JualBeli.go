package main

import "fmt"

const NMAX = 1000 // Max number of arrBarang

type barang struct {
	// SubType untuk memberi nama, katagori, harga modal, dan harga jual
	nama, katagori        string
	hargaModal, hargaJual float64
}

type arrBarang [NMAX - 1]barang // Array untuk menyimpan data barang yang akan di jual

func main() {
	UI_Main()
	clearScreen()
}

func UI_Main() {
	// tampilan layar utama (home page awal)
	fmt.Println("=====================|ADMIN MENU V.1|=====================")
	fmt.Printf("%s\n", "INPUT BARANG  [1]")
	fmt.Printf("%s\n", "EDIT BARANG   [2]")
	fmt.Printf("%s\n", "CEK BARANG    [3]")
	fmt.Println("=====================|  TOKO KEREN  |=====================")
}

func isiBarang(b *arrBarang, n *int) {}

func clearScreen() {
	// membersihkan UI agar enak dilihat
	fmt.Print("\033[H\033[2J")
}
