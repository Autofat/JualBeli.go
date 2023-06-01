package main

import "fmt"

const NMAX = 1000 // Max number of arrBarang

type barang struct {
	// SubType untuk memberi nama, katagori, harga modal, dan harga jual
	nama, kategori        string
	hargaModal, hargaJual float64
	transaksi             int
}

type arrBarang [NMAX - 1]barang // Array untuk menyimpan data barang yang akan di jual

func main() {
	var b arrBarang
	var n, perintah_main int
	var target, A string

	for {
		clearScreen()
		UI_Main()
		fmt.Print("Silahkan pilih perintah: ")
		fmt.Scan(&perintah_main)

		if perintah_main == 1 {
			clearScreen()
			isiBarang(&b, &n)
			fmt.Println("==========================================================")
			fmt.Print("          TEKAN [X] UNTUK KEMBALI KE HALAMAN UTAMA          ")
			fmt.Scan(&A)
		} else if perintah_main == 2 {
			clearScreen()
			fmt.Println("=====================| EDIT BARANG |======================")
			fmt.Print("NAMA BARANG YANG INGIN DI EDIT: ")
			fmt.Scan(&target)
			cekBarang(b, n, target)
			if cekBarang(b, n, target) >= 0 {
				clearScreen()
				fmt.Println("=====================| EDIT BARANG |======================")
				fmt.Println("|                                                        |")
				fmt.Println("|                   [ DATA DITEMUKAN ]                   |")
				fmt.Println("|         SILAHKAN MASUKAN DATA YANG MAU DI EDIT         |")
				fmt.Println("|                                                        |")
				fmt.Println("=====================|  TOKO KEREN  |=====================")
				fmt.Println(b[cekBarang(b, n, target)])
				fmt.Println("==========================================================")
				editData(&b, n, cekBarang(b, n, target))

			} else {
				clearScreen()
				fmt.Println("=====================| EDIT BARANG |======================")
				fmt.Println("|                                                        |")
				fmt.Println("|                [ DATA TIDAK DITEMUKAN ]                |")
				fmt.Println("|        SILAHKAN CEK KEMBALI DATA YANG INGIN DI EDIT    |")
				fmt.Println("|                                                        |")
				fmt.Println("=====================|  TOKO KEREN  |=====================")
				fmt.Print("          TEKAN [X] UNTUK KEMBALI KE HALAMAN UTAMA          ")
				fmt.Scan(&A)
			}
		}

	}
}

func UI_Main() {
	// tampilan layar utama (home page awal)
	fmt.Println("======================| ADMIN MENU |=======================")
	fmt.Println("|                        [ MENU ]                         |")
	fmt.Println("|                                                         |")
	fmt.Println("|                    1. Input Barang                      |")
	fmt.Println("|                    2. Edit Barang                       |")
	fmt.Println("|                    3. Cek Barang                        |")
	fmt.Println("|                    4. Exit                              |")
	fmt.Println("|                                                         |")
	fmt.Println("======================|  TOKO KEREN  |====================")
}

func isiBarang(b *arrBarang, n *int) {
	var namaBarang, kategori string
	var hargaModal, hargaJual float64
	var transaksi int
	fmt.Println("=====================| INPUT BARANG |=====================")
	fmt.Println("|                                                        |")
	fmt.Println("|                MASUKAN DATA DENGAN FORMAT              |")
	fmt.Println("| NAMA-BARANG KATAGORI HARGA-MODAL HARGA-JUAL TRANSAKSI  |")
	fmt.Println("|                                                        |")
	fmt.Println("=====================|  TOKO KEREN  |=====================")
	fmt.Scan(&namaBarang, &kategori, &hargaModal, &hargaJual, &transaksi)
	for i := 0; i < NMAX && namaBarang != "XXX"; i++ {
		b[i].nama = namaBarang
		b[i].kategori = kategori
		b[i].hargaModal = hargaModal
		b[i].hargaJual = hargaJual
		b[i].transaksi = transaksi
		*n++
		fmt.Scan(&namaBarang, &kategori, &hargaModal, &hargaJual, &transaksi)
	}
}

func editData(b *arrBarang, n, i int) {
	fmt.Println("DENGAN FORMAT: NAMA-BARANG KATAGORI HARGA-MODAL HARGA-JUAL TRANSAKSI")
	fmt.Scan(&b[i].nama, &b[i].kategori, &b[i].hargaModal, &b[i].hargaJual, &b[i].transaksi)
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
