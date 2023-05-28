package main

import "fmt"

const NMAX = 1000 // Max number of arrBarang

type barang struct {
	// SubType untuk memberi nama, katagori, harga modal, dan harga jual
	nama, kategori        string
	hargaModal, hargaJual float64
}

type arrBarang [NMAX - 1]barang // Array untuk menyimpan data barang yang akan di jual

func main() {
	var b arrBarang
	var n, perintah_main int
	var target string

	UI_Main()
	fmt.Print("Silahkan pilih perintah: ")
	fmt.Scan(&perintah_main)
	// if perintah_main == 1 {
	clearScreen()
	isiBarang(&b, &n)
	displayArr(b, n)
	// } else if perintah_main == 2 {
	fmt.Scan(&target)
	cekBarang(b, n, target)
	if cekBarang(b, n, target) > 0 {
		fmt.Println("=====================| EDIT BARANG |=====================")
		fmt.Println("DATA DITEMUKAN")
		fmt.Println("=====================|  TOKO KEREN  |=====================")
	} else {
		fmt.Println("=====================| EDIT BARANG |=====================")
		fmt.Println("DATA TIDAK DITEMUKAN")
		fmt.Println("=====================|  TOKO KEREN  |=====================")
	}
	displayArr(b, n)
	// }

}

func UI_Main() {
	// tampilan layar utama (home page awal)
	fmt.Println("=====================|ADMIN MENU V.1|=====================")
	fmt.Println("%s\n", "INPUT BARANG  [1]")
	fmt.Println("%s\n", "EDIT BARANG   [2]")
	fmt.Println("%s\n", "CEK BARANG    [3]")
	fmt.Println("=====================|  TOKO KEREN  |=====================")
}

func isiBarang(b *arrBarang, n *int) {
	var namaBarang, kategori string
	var hargaModal, hargaJual float64
	fmt.Println("=====================| INPUT BARANG |=====================")
	fmt.Println("Masukan dengan format")
	fmt.Println("NAMA-BARANG KATAGORI HARGA-MODAL HARGA-JUAL")
	fmt.Println("=====================|  TOKO KEREN  |=====================")
	fmt.Scan(&namaBarang, &kategori, &hargaModal, &hargaJual)
	for i := 0; i < NMAX && namaBarang != "XXX"; i++ {
		b[i].nama = namaBarang
		b[i].kategori = kategori
		b[i].hargaModal = hargaModal
		b[i].hargaJual = hargaJual
		*n++
		fmt.Scan(&namaBarang, &kategori, &hargaModal, &hargaJual)
	}
}

func cekBarang(b arrBarang, n int, target string) int {
	for i := 0; i <= n; i++ {
		if target == b[i].nama {
			return i
		}
	}
	return -1
}

func displayArr(b arrBarang, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(b[i])
	}
}

func clearScreen() {
	// membersihkan UI agar enak dilihat
	fmt.Print("\033[H\033[2J")
}
