package main

import "fmt"

const NMAX = 1024

type barang struct {
	nama, kategori        string
	hargaModal, hargaJual float64
	stok, penjualan       int
}

type transaksi struct {
	tanggal   string
	idx       int
	pembelian arrPembelian
}

type pembelian struct {
	barang string
	jumlah int
}

type arrBarang [NMAX]barang
type arrTransaksi [NMAX]transaksi
type arrPembelian [NMAX]pembelian

// ===================================== FUNC MAIN ============================================ //
func main() {
	var perintah_main, perintah_barang, perintah_transaksi, n, i, m, tgl int
	var totalModal, totalPendapatan float64
	var targetBarang, A, perintah_hapus_barang, targetTransaksi, targetTanggal string
	var b arrBarang
	var t arrTransaksi
	// var t arrTransaksi
	// var t arrTransaksi
	for perintah_main != 4 {
		clearScreen()
		interfaceMain()
		fmt.Print("Silahkan pilih perintah: ")
		fmt.Scan(&perintah_main)
		switch perintah_main {
		// ================================================================= BARANG COMMANDS =================================================================
		case 1:
			clearScreen()
			interfaceBarang()
			fmt.Print("Silahkan pilih perintah: ")
			fmt.Scan(&perintah_barang)
			if perintah_barang == 1 {
				clearScreen()
				inputBarang(&b, &n, &totalModal)
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
				} else {
					clearScreen()
					fmt.Println("====================| HAPUS BARANG |======================")
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
		// ================================================================= END BARANG COMMANDS ================================================================

		// ================================================================= TRANSAKSI COMMANDS =================================================================
		case 2:
			clearScreen()
			interfaceTransaksi()
			fmt.Print("Silahkan pilih perintah: ")
			fmt.Scan(&perintah_transaksi)
			if perintah_transaksi == 1 {
				clearScreen()
				inputTransaksi(&t, &b, &tgl, n)
			} else if perintah_transaksi == 2 {
				clearScreen()
				fmt.Println("====================| EDIT TRANSAKSI |====================")
				fmt.Println("MASUKAN TANGGAL TRANSAKSI YANG MAU DI EDIT: ")
				fmt.Scan(&targetTanggal)
				CTT := cekTransaksiTanggal(t, m, targetTanggal)
				if CTT >= 0 {
					clearScreen()
					fmt.Println("====================| EDIT TRANSAKSI |====================")
					fmt.Println("|                                                        |")
					fmt.Println("|                   [ DATA DITEMUKAN ]                   |")
					fmt.Println("|       SILAHKAN PILIH NAMA ITEM YANG INGIN DI RUBAH     |")
					fmt.Println("|                                                        |")
					fmt.Println("==========================================================")
					fmt.Println("TANGGAL: ", targetTanggal)
					fmt.Println("=====================|    BARANG    |=====================")
					// for i := 0; i < brg; i++ {
					// 	if t[CTT].pembelian[i].barang != "" {
					// 		fmt.Println(t[CTT].pembelian[i])
					// 	}
					// }
					fmt.Println(t[CTT].pembelian)
					fmt.Println("=====================| EDIT BARANG |======================")
					fmt.Print("MASUKAN NAMA BARANG YG INGIN DI EDIT: ")
					fmt.Scan(&targetTransaksi)
					fmt.Println("=====================|  TOKO KEREN  |=====================")
					CTB := cekTransaksiBarang(t, CTT, targetTransaksi)
					if CTB >= 0 {
						clearScreen()
						fmt.Println("====================| EDIT TRANSAKSI |====================")
						fmt.Println("|                                                        |")
						fmt.Println("|                   [ DATA DITEMUKAN ]                   |")
						fmt.Println("|       SILAHKAN PILIH NAMA ITEM YANG INGIN DI RUBAH     |")
						fmt.Println("|                                                        |")
						fmt.Println("==========================================================")
						fmt.Println("ITEM: ", t[CTT].pembelian[CTB])
						fmt.Println("=====================|    BARANG    |=====================")
						editTransaksi(&t, CTT, CTB)
					}
				} else {
					clearScreen()
					fmt.Println("====================| EDIT TRANSAKSI |====================")
					fmt.Println("|                                                        |")
					fmt.Println("|                [ DATA TIDAK DITEMUKAN ]                |")
					fmt.Println("|          SILAHKAN CEK KEMBALI INPUTAN TANGGAL          |")
					fmt.Println("|                                                        |")
					fmt.Println("=====================|  TOKO KEREN  |=====================")
					fmt.Print("          TEKAN [X] UNTUK KEMBALI KE HALAMAN UTAMA          ")
					fmt.Scan(&A)
				}
			} else if perintah_transaksi == 3 {
				clearScreen()
				fmt.Println("====================| HAPUS TRANSAKSI |======================")
				fmt.Print("PADA TANGGA BERAPA YANG INGIN DI HAPUS: ")
				fmt.Scan(&targetTanggal)
				CTT := cekTransaksiTanggal(t, n, targetTanggal)
				if CTT >= 0 {
					clearScreen()
					fmt.Println("===================| HAPUS TRANSAKSI |====================")
					fmt.Println("|                                                        |")
					fmt.Println("|                   [ DATA DITEMUKAN ]                   |")
					fmt.Println("|               APAKAH DATA INGIN DI HAPUS               |")
					fmt.Println("|                       - Y / N -                        |")
					fmt.Println("|                                                        |")
					fmt.Println("=====================|  TOKO KEREN  |=====================")
					for i := 0; i < t[CTT].idx; i++ {
						fmt.Println(t[CTT].pembelian[i])
					}
					fmt.Println("==========================================================")
					fmt.Scan(&perintah_hapus_barang)
					if perintah_hapus_barang == "Y" {
						hapusBarangTransaksi(&t, &b, CTT)
						hapusTransaksi(&t, &tgl, CTT)
						fmt.Println(t[CTT].pembelian)
					}
				} else {
					clearScreen()
					fmt.Println("===================| HAPUS TRANSAKSI |====================")
					fmt.Println("|                                                        |")
					fmt.Println("|                [ DATA TIDAK DITEMUKAN ]                |")
					fmt.Println("|        SILAHKAN CEK KEMBALI TANGGAL TRANSAKSI          |")
					fmt.Println("|                  YANG INGIN DI HAPUS                   |")
					fmt.Println("|                                                        |")
					fmt.Println("=====================|  TOKO KEREN  |=====================")
					fmt.Print("          TEKAN [X] UNTUK KEMBALI KE HALAMAN UTAMA          ")
					fmt.Scan(&A)

				}
			} else if perintah_transaksi == 4 {
				clearScreen()
				fmt.Println("====================| CEK TRANSAKSI |=====================")
				displayTransaksi(t, tgl)
				fmt.Println("=====================|  TOKO KEREN  |=====================")
				fmt.Print("          TEKAN [X] UNTUK KEMBALI KE HALAMAN UTAMA          ")
				fmt.Scan(&A)

			}
		case 3:
			clearScreen()
			statistics(&b, &t, n, m, totalModal, totalPendapatan)
		}
	}
}

// ================================================================= END TRANSAKSI COMMANDS =================================================================

// ===================================== END FUNC MAIN ============================================ //

// ===================================== FUNC INTERFACE ============================================ //
func interfaceMain() {
	// tampilan layar utama (home page awal)
	fmt.Println("======================| ADMIN MENU |=======================")
	fmt.Println("|                        [ MENU ]                         |")
	fmt.Println("|                                                         |")
	fmt.Println("|                    1. 📦 Barang                         |")
	fmt.Println("|                    2. 💵 Transaksi                      |")
	fmt.Println("|                    3. 📊 Statistik                      |")
	fmt.Println("|                    4. ❌ Exit                           |")
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
func inputBarang(b *arrBarang, n *int, totalModal *float64) {
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
	for i := 0; i < *n; i++ {
		*totalModal += b[i].hargaModal * float64(b[i].stok)
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
func inputTransaksi(t *arrTransaksi, b *arrBarang, tgl *int, n int) {
	var namaBarang, tanggal string
	var jmlBeli int
	fmt.Println("===================| INPUT TRANSAKSI |====================")
	fmt.Println("|                                                        |")
	fmt.Println("|                MASUKAN DATA DENGAN FORMAT              |")
	fmt.Println("|      TANGGAL-JUAL NAMA-BARANG JUMLAH-PEMBELIAN         |")
	fmt.Println("|                                                        |")
	fmt.Println("=====================|  TOKO KEREN  |=====================")
	fmt.Scan(&tanggal)
	fmt.Scan(&namaBarang, &jmlBeli)

	for namaBarang != "XXX" && t[*tgl].idx < NMAX {
		t[*tgl].tanggal = tanggal
		t[*tgl].pembelian[t[*tgl].idx].barang = namaBarang
		t[*tgl].pembelian[t[*tgl].idx].jumlah = jmlBeli
		for i := 0; i < n; i++ {
			if b[i].nama == namaBarang {
				b[i].penjualan = b[i].penjualan + jmlBeli
				b[i].stok = b[i].stok - jmlBeli
			}
		}
		t[*tgl].idx++
		fmt.Scan(&namaBarang, &jmlBeli)
	}
	*tgl++
}

func editTransaksi(t *arrTransaksi, CTT, CTB int) {
	fmt.Println("DENGAN FORMAT: NAMA-BARANG  JUMLAH-BELI")
	fmt.Scan(&t[CTT].pembelian[CTB].barang, &t[CTT].pembelian[CTB].jumlah)
}

func cekTransaksiTanggal(t arrTransaksi, m int, target string) int {
	for i := 0; i <= m; i++ {
		if target == t[i].tanggal {
			return i
		}
	}
	return -1
}

func cekTransaksiBarang(t arrTransaksi, CTT int, target string) int {
	for i := 0; i <= t[CTT].idx; i++ {
		if target == t[CTT].pembelian[i].barang {
			return i
		}
	}
	return -1
}

func hapusTransaksi(t *arrTransaksi, tgl *int, target int) {
	for i := target; i < *tgl; i++ {
		t[i] = t[i+1]
		*tgl--
	}
}
func hapusBarangTransaksi(t *arrTransaksi, b *arrBarang, CTT int) {
	for i := 0; i < NMAX-1; i++ {
		for j := 0; j < NMAX-1; j++ {
			if t[CTT].pembelian[i].barang == b[j].nama {
				b[j].penjualan = b[j].penjualan - t[CTT].pembelian[i].jumlah
				b[j].stok = b[j].stok + t[CTT].pembelian[i].jumlah
			}
		}
		t[CTT].pembelian[i] = t[CTT].pembelian[i+1]
		t[CTT].idx--
	}

}

func displayTransaksi(t arrTransaksi, tgl int) {
	for j := 0; j < tgl; j++ {
		fmt.Println(t[j].tanggal)
		for i := 0; i < t[j].idx; i++ {
			fmt.Println(t[j].pembelian[i])
		}
		fmt.Println("--------------------------------")
	}
}

// ===================================== END FUNC TRANSAKSI ============================================ //
// ===================================== FUNC STATISTIK ============================================ //
func statistics(b *arrBarang, t *arrTransaksi, n, m int, totalModal float64, totalPendapatan float64) {
	fmt.Println("=======================| STATISTIK |========================")
	fmt.Println("|                                                          |")
	fmt.Println("|                 DATA MODAL DAN PENDAPATAN                |")
	fmt.Println("|                                                          |")
	fmt.Println("============================================================")
	for i := 0; i < n; i++ {
		totalPendapatan += b[i].hargaJual * float64(b[i].penjualan)
	}
	fmt.Println("Total Modal: Rp.", totalModal)
	fmt.Println("Total Pendapatan: Rp.", totalPendapatan)
	fmt.Println("============================================================")
	fmt.Println("|                                                          |")
	fmt.Println("|            [ 5 BARANG PALING BANYAK TERJUAL ]            |")
	fmt.Println("|                                                          |")
	fmt.Println("============================================================")

	shortingBarang(b, n)

	// Display top 5
	for i := 0; i < 5 && i < n; i++ {
		fmt.Println(b[i].nama, "-", b[i].penjualan, "unit")
	}
	fmt.Println("============================================================")
	fmt.Println("|                      TOKO KEREN                          |")
	fmt.Println("============================================================")
	fmt.Println("          TEKAN [X] UNTUK KEMBALI KE HALAMAN UTAMA          ")
	var A string
	fmt.Scan(&A)
}

// ===================================== END FUNC STATISTIK ======================================== //
func clearScreen() {
	// membersihkan UI agar enak dilihat
	fmt.Print("\033[H\033[2J")
}
