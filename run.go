package main

import "fmt"

const NMAX = 1024

type barang struct {
	nama, kategori        string
	hargaModal, hargaJual float64
	stok, penjualan       int
}

type transaksi struct {
	tanggal string
	jumlah  int
	barang  barang
}

type arrBarang [NMAX]barang
type arrTransaksi [NMAX]transaksi

// ===================================== FUNC MAIN ============================================ //
func main() {
	var perintah_main, perintah_barang, n, i int
	var targetBarang, A, perintah_hapus_barang string
	var b arrBarang
	// var t arrTransaksi
	for perintah_main != 4 {
		clearScreen()
		interfaceMain()
		fmt.Print("Silahkan pilih perintah: ")
		fmt.Scan(&perintah_main)
		switch perintah_main {
		case 1:
			clearScreen()
			interfaceBarang()
			fmt.Print("Silahkan pilih perintah: ")
			fmt.Scan(&perintah_barang)
			if perintah_barang == 1 {
				clearScreen()
				inputBarang(&b, &n)
			} else if perintah_barang == 2 {
				clearScreen()
				fmt.Println("=====================| EDIT BARANG |======================")
				fmt.Print("NAMA BARANG YANG INGIN DI EDIT: ")
				fmt.Scan(&targetBarang)
				cekBarang(b, n, targetBarang)
				if cekBarang(b, n, targetBarang) >= 0 {
					clearScreen()
					fmt.Println("=====================| EDIT BARANG |======================")
					fmt.Println("|                                                        |")
					fmt.Println("|                   [ DATA DITEMUKAN ]                   |")
					fmt.Println("|         SILAHKAN MASUKAN DATA YANG MAU DI EDIT         |")
					fmt.Println("|                                                        |")
					fmt.Println("=====================|  TOKO KEREN  |=====================")
					fmt.Println(b[cekBarang(b, n, targetBarang)])
					fmt.Println("==========================================================")
					editBarang(&b, n, i)

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
			} else if perintah_barang == 3 {
				clearScreen()
				fmt.Println("====================| HAPUS BARANG |======================")
				fmt.Print("NAMA BARANG YANG INGIN DI HAPUS: ")
				fmt.Scan(&targetBarang)
				cekBarang(b, n, targetBarang)
				if cekBarang(b, n, targetBarang) >= 0 {
					clearScreen()
					fmt.Println("====================| HAPUS BARANG |======================")
					fmt.Println("|                                                        |")
					fmt.Println("|                   [ DATA DITEMUKAN ]                   |")
					fmt.Println("|               APAKAH DATA INGIN DI HAPUS               |")
					fmt.Println("|                       - Y / N -                        |")
					fmt.Println("|                                                        |")
					fmt.Println("=====================|  TOKO KEREN  |=====================")
					fmt.Println(b[cekBarang(b, n, targetBarang)])
					fmt.Println("==========================================================")
					fmt.Scan(&perintah_hapus_barang)
					if perintah_hapus_barang == "Y" {
						hapusBarang(&b, &n, cekBarang(b, n, targetBarang))
					}
				}else{				\
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
			} else if perintah_barang == 4 {
				clearScreen()
				fmt.Println("=====================| EDIT BARANG |======================")
				displayBarang(b, n)
				fmt.Println("=====================|  TOKO KEREN  |=====================")
				fmt.Print("          TEKAN [X] UNTUK KEMBALI KE HALAMAN UTAMA          ")
				fmt.Scan(&A)
			}

		}

	}

}

// ===================================== END FUNC MAIN ============================================ //

// ===================================== FUNC INTERFACE ============================================ //
func interfaceMain() {
	// tampilan layar utama (home page awal)
	fmt.Println("======================| ADMIN MENU |=======================")
	fmt.Println("|                        [ MENU ]                         |")
	fmt.Println("|                                                         |")
	fmt.Println("|                    1. Barang                            |")
	fmt.Println("|                    2. Transaksi                         |")
	fmt.Println("|                    3. Statistik                         |")
	fmt.Println("|                    4. Exit                              |")
	fmt.Println("|                                                         |")
	fmt.Println("======================|  TOKO KEREN  |=====================")
}

func interfaceBarang() {
	// tampilan layar utama (home page awal)
	fmt.Println("========================| BARANG |=========================")
	fmt.Println("|                        [ MENU ]                         |")
	fmt.Println("|                                                         |")
	fmt.Println("|                    1. Input Barang                      |")
	fmt.Println("|                    2. Edit Barang                       |")
	fmt.Println("|                    3. Hapus Barang                      |")
	fmt.Println("|                    4. Cek Barang                        |")
	fmt.Println("|                    5. Exit                              |")
	fmt.Println("|                                                         |")
	fmt.Println("======================|  TOKO KEREN  |=====================")
}
func interfaceTransaksi() {
	// tampilan layar utama (home page awal)
	fmt.Println("=======================| TRANSAKSI |=======================")
	fmt.Println("|                        [ MENU ]                         |")
	fmt.Println("|                                                         |")
	fmt.Println("|                    1. Input Transaksi                   |")
	fmt.Println("|                    2. Edit Transaksi                    |")
	fmt.Println("|                    3. Hapus Transaksi                   |")
	fmt.Println("|                    4. Cek Transaksi                     |")
	fmt.Println("|                    5. Exit                              |")
	fmt.Println("|                                                         |")
	fmt.Println("======================|  TOKO KEREN  |=====================")
}

// ===================================== END FUNC INTERFACE ============================================ //

// ===================================== FUNC BARANG ============================================ //
func inputBarang(b *arrBarang, n *int) {
	var namaBarang, kategori string
	var hargaModal, hargaJual float64
	var stok int
	fmt.Println("=====================| INPUT BARANG |=====================")
	fmt.Println("|                                                        |")
	fmt.Println("|                MASUKAN DATA DENGAN FORMAT              |")
	fmt.Println("|    NAMA-BARANG KATAGORI HARGA-MODAL HARGA-JUAL STOK    |")
	fmt.Println("|                                                        |")
	fmt.Println("=====================|  TOKO KEREN  |=====================")
	fmt.Scan(&namaBarang, &kategori, &hargaModal, &hargaJual, &stok)
	for i := 0; namaBarang != "XXX"; i++ {
		b[i].nama = namaBarang
		b[i].kategori = kategori
		b[i].hargaModal = hargaModal
		b[i].hargaJual = hargaJual
		b[i].stok = stok
		*n++
		fmt.Scan(&namaBarang, &kategori, &hargaModal, &hargaJual, &stok)
	}
}
func editBarang(b *arrBarang, n, i int) {
	fmt.Println("DENGAN FORMAT: NAMA-BARANG KATAGORI HARGA-MODAL HARGA-JUAL STOK")
	fmt.Scan(&b[i].nama, &b[i].kategori, &b[i].hargaModal, &b[i].hargaJual, &b[i].stok)
}

func cekBarang(b arrBarang, n int, target string) int {
	for i := 0; i <= n; i++ {
		if target == b[i].nama {
			return i
		}
	}
	return -1

}
func shortingBarang(b *arrBarang, n int) {
	// shorting via penjualan, akan menampilkan barang yang paling banayk di beli/paling banyak terjual
	for i := 0; i < n; i++ {
		max := i
		for j := i + 1; j < n; j++ {
			if b[max].penjualan < b[j].penjualan {
				max = j
			}
		}
		b[i], b[max] = b[max], b[i]
	}
}

func displayBarang(b arrBarang, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(b[i])
	}
}

func hapusBarang(b *arrBarang, n *int, target int) {
	for i := target; i < *n; i++ {
		b[i] = b[i+1]
	}
	*n--

}

// ===================================== END FUNC BARANG ============================================ //

// ===================================== FUNC TRANSAKSI ============================================ //

// ===================================== END FUNC TRANSAKSI ============================================ //
func clearScreen() {
	// membersihkan UI agar enak dilihat
	fmt.Print("\033[H\033[2J")
}
