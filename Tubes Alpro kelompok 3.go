package main

import "fmt"

const NMAX int = 100

type rekening struct {
	nama         string
	norek        string
	tanggalBuat  string
	caBank       string
	saldo        float64
	tabTransaksi [NMAX]transaksi
	hitungT      int
}

type nasabah struct {
	nama        string
	norek       string
	tabrekening rekening
}

type transaksi struct {
	tanggal string
	jumlah  float64
	tipe    string // "Debit" or "Credit"
}

type tabRekening [NMAX]rekening
type tabNasabah [NMAX]nasabah

func addRekening(R *tabRekening, hitungR *int) {
	if *hitungR < NMAX {
		fmt.Printf("Masukkan nama untuk rekening %d: ", *hitungR+1)
		fmt.Scan(&R[*hitungR].nama)
		fmt.Printf("Masukkan nomor rekening untuk rekening %d: ", *hitungR+1)
		fmt.Scan(&R[*hitungR].norek)
		fmt.Printf("Masukkan tanggal lahir untuk rekening %d (YYYY-MM-DD): ", *hitungR+1)
		fmt.Scan(&R[*hitungR].tanggalBuat)
		fmt.Printf("Masukkan cabang bank untuk rekening %d: ", *hitungR+1)
		fmt.Scan(&R[*hitungR].caBank)
		fmt.Printf("Masukkan saldo awal untuk rekening %d: ", *hitungR+1)
		fmt.Scan(&R[*hitungR].saldo)
		*hitungR++
		fmt.Println("〚 Data rekening berhasil disimpan 〛")
	} else {
		fmt.Println("〚 Data rekening penuh 〛")
	}
}

func addNasabah(N *tabNasabah, hitungN *int, R tabRekening, hitungR int) {
	var i int

	if *hitungN < NMAX {

		fmt.Printf("Masukkan nama untuk nasabah ke-%d: ", *hitungN+1)
		fmt.Scan(&N[*hitungN].nama)
		fmt.Printf("Masukkan nomor rekening yang dimiliki nasabah ke-%d: ", *hitungN+1)
		fmt.Scan(&N[*hitungN].norek)

		for i = 0; i < hitungR; i++ {
			if N[*hitungN].norek == R[i].norek {
				N[*hitungN].tabrekening = R[i]
			}
		}
		*hitungN++

		fmt.Println("〚 Data nasabah berhasil disimpan 〛")
	} else {
		fmt.Println("〚 Data nasabah penuh 〛")
	}
}

func pilihanCariData(R tabRekening, hitungR int) {
	var pilih, index, j int
	var xStr string
	var xTrs transaksi

	index = -1

	fmt.Scan(&pilih)
	if pilih == 1 {
		fmt.Print("Masukkan nomor rekening yang akan dicari: ")
		fmt.Scan(&xStr)
		index = cariDatabyNoRek(R, hitungR, xStr)
		if index == -1 {
			fmt.Println("〚 Akun tidak ditemukkan 〛")
		}
	} else if pilih == 2 {
		fmt.Print("Masukkan cabang bank yang akan dicari: ")
		fmt.Scan(&xStr)
		cariDatabyCaBank(R, hitungR, xStr)
	} else if pilih == 3 {
		fmt.Print("Masukkan data transaksi yang akan dicari (tanggal  tipe  jumlah): ")
		fmt.Scan(&xTrs.tanggal, &xTrs.tipe, &xTrs.jumlah)
		index = cariDatabyTransaksi(R, hitungR, xTrs.tanggal, xTrs.tipe, xTrs.jumlah)
		if index == -1 {
			fmt.Println("〚 Akun tidak ditemukkan 〛")
		}
	} else if pilih == 4 {
		fmt.Print()
	} else {
		fmt.Println("〚 Pilihan tidak ditemukkan 〛")
	}

	if index != -1 {
		sortTransaksi(&R[index], R[index].hitungT)
		fmt.Println("——————————————————————————————————————————————————")
		fmt.Printf("Nama              : %s\n", R[index].nama)
		fmt.Printf("Nomor Rekening    : %s\n", R[index].norek)
		fmt.Printf("Tanggal Pembuatan : %s\n", R[index].tanggalBuat)
		fmt.Printf("Cabang Bank       : %s\n", R[index].caBank)
		fmt.Printf("Saldo             : %.2f\n", R[index].saldo)
		fmt.Println("〚 Riwayat Transaksi 〛")
		if R[index].hitungT <= 0 {
			fmt.Println("Tidak ada")
		}
		for j = 0; j < R[index].hitungT; j++ {
			fmt.Printf("%s        : %s, %.2f\n", R[index].tabTransaksi[j].tanggal, R[index].tabTransaksi[j].tipe, R[index].tabTransaksi[j].jumlah)
		}
		fmt.Println("——————————————————————————————————————————————————")
	}
}

func cariDatabyNoRek(R tabRekening, hitungR int, x string) int {
	var i, found int

	found = -1
	i = 0
	for i < hitungR && found == -1 {
		if R[i].norek == x {
			found = i
		}
		i++
	}
	return found
}

func cariDatabyCaBank(R tabRekening, hitungR int, x string) {
	var i, j, k int

	k = 1

	for i = 0; i < hitungR; i++ {

		sortTransaksi(&R[i], R[i].hitungT)

		if x == R[i].caBank {
			fmt.Println("——————————————————————————————————————————————————")
			fmt.Printf("Rekening ke-%d\n", k)
			fmt.Printf("Nama              : %s\n", R[i].nama)
			fmt.Printf("Nomor Rekening    : %s\n", R[i].norek)
			fmt.Printf("Tanggal Pembuatan : %s\n", R[i].tanggalBuat)
			fmt.Printf("Cabang Bank       : %s\n", R[i].caBank)
			fmt.Printf("Saldo             : %.2f\n", R[i].saldo)
			fmt.Println("〚 Riwayat Transaksi 〛")
			if R[i].hitungT <= 0 {
				fmt.Println("Tidak ada")
			}
			for j = 0; j < R[i].hitungT; j++ {
				fmt.Printf("%s        : %s, %.2f\n", R[i].tabTransaksi[j].tanggal, R[i].tabTransaksi[j].tipe, R[i].tabTransaksi[j].jumlah)
			}
			k++
		}
	}
	if k == 1 {
		fmt.Println("〚 Akun tidak ditemukkan 〛")
	}
}

func cariDatabyTransaksi(R tabRekening, hitungR int, xTgl, xTip string, xJum float64) int {
	var i, j, found int

	found = -1
	i = 0
	for i < hitungR && found == -1 {
		j = 0
		for j < R[i].hitungT && found == -1 {
			if R[i].tabTransaksi[j].tanggal == xTgl && R[i].tabTransaksi[j].tipe == xTip && R[i].tabTransaksi[j].jumlah == xJum {
				found = i
			}
			j++
		}
		i++
	}
	return found
}

func editRekening(R *tabRekening, hitungR int) {
	var index int
	var namaBaru, norekBaru, tanggalBaru, cabankBaru string
	var norek string

	fmt.Print("Masukkan nomor rekening yang akan di edit: ")
	fmt.Scan(&norek)
	index = cariDatabyNoRek(*R, hitungR, norek)

	if index != -1 {
		fmt.Print("Masukkan nama baru: ")
		fmt.Scan(&namaBaru)
		fmt.Print("Masukkan nomor rekening baru: ")
		fmt.Scan(&norekBaru)
		fmt.Print("Masukkan tanggal baru (YYYY-MM-DD): ")
		fmt.Scan(&tanggalBaru)
		fmt.Print("Masukkan cabang bank baru: ")
		fmt.Scan(&cabankBaru)

		R[index].nama = namaBaru
		R[index].norek = norekBaru
		R[index].tanggalBuat = tanggalBaru
		R[index].caBank = cabankBaru

		fmt.Println("〚 Akun telah selesai di edit 〛")
	} else {
		fmt.Println("〚 Akun tidak ditemukan 〛")
	}
}

func hapusRekening(R *tabRekening, hitungR *int) {
	var norek string
	var index, i int
	fmt.Print("Masukkan nomor rekening yang akan di hapus: ")
	fmt.Scan(&norek)
	index = cariDatabyNoRek(*R, *hitungR, norek)
	if index != -1 {
		for i = index; i < *hitungR-1; i++ {
			R[i] = R[i+1]
		}
		*hitungR--
		fmt.Println("〚 Akun telah dihapus 〛")
	} else {
		fmt.Println("〚 Akun tidak ditemukan 〛")
	}
}

func sortTanggal(R *tabRekening, hitungR int) {
	var i, pass int
	var temp rekening

	pass = 1
	for pass < hitungR {
		temp = R[pass]
		i = pass

		for i > 0 && temp.tanggalBuat < R[i-1].tanggalBuat {
			R[i] = R[i-1]
			i = i - 1
		}

		R[i] = temp
		pass = pass + 1
	}
}

func sortTransaksi(R *rekening, hitungT int) {
	var i, pass int
	var temp transaksi

	pass = 1
	for pass < hitungT {
		temp = R.tabTransaksi[pass]
		i = pass

		for i > 0 && temp.tanggal < R.tabTransaksi[i-1].tanggal {
			R.tabTransaksi[i] = R.tabTransaksi[i-1]
			i = i - 1
		}

		R.tabTransaksi[i] = temp
		pass = pass + 1
	}
}

func cetakRekening(N tabNasabah, R tabRekening, hitungR, hitungN int) {
	var i, j, k int

	sortTanggal(&R, hitungR)

	for i = 0; i < hitungR; i++ {
		sortTransaksi(&R[i], R[i].hitungT)
		fmt.Println("——————————————————————————————————————————————————")
		fmt.Printf("Rekening ke-%d\n", i+1)

		for j = 0; j < hitungN; j++ {
			if R[i].norek == N[j].norek {
				fmt.Printf("Nasabah           : %s\n", N[j].nama)
			}
		}

		fmt.Printf("Nama              : %s\n", R[i].nama)
		fmt.Printf("Nomor Rekening    : %s\n", R[i].norek)
		fmt.Printf("Tanggal Pembuatan : %s\n", R[i].tanggalBuat)
		fmt.Printf("Cabang Bank       : %s\n", R[i].caBank)
		fmt.Printf("Saldo             : %.2f\n", R[i].saldo)
		fmt.Println("〚 Riwayat Transaksi 〛")

		if R[i].hitungT <= 0 {
			fmt.Println("Tidak ada")
		}

		for k = 0; k < R[i].hitungT; k++ {
			fmt.Printf("%s        : %s, %.2f\n", R[i].tabTransaksi[k].tanggal, R[i].tabTransaksi[k].tipe, R[i].tabTransaksi[k].jumlah)
		}
	}
}

func addTransaksi(N *tabNasabah, R *rekening, jumlah float64, tanggal, tipe string, hitungN int) {
	var i int

	if R.hitungT < NMAX {

		R.tabTransaksi[R.hitungT].tanggal = tanggal
		R.tabTransaksi[R.hitungT].tipe = tipe
		R.tabTransaksi[R.hitungT].jumlah = jumlah
		R.hitungT++

		if tipe == "debit" {
			R.saldo = R.saldo - jumlah
		} else if tipe == "kredit" {
			R.saldo = R.saldo + jumlah
		}

		for i = 0; i < hitungN; i++ {
			if R.norek == N[i].norek {
				N[i].tabrekening = *R
			}
		}

	} else {
		fmt.Println("〚 Data transaksi penuh 〛")
	}
}

func transaksiSaldo(N *tabNasabah, R *tabRekening, hitungR, hitungN int) {
	var rAwal, rTuju string
	var jumlah float64
	var tanggal string
	var awal, tuju int
	fmt.Print("Masukkan nomor rekening sumber: ")
	fmt.Scan(&rAwal)
	fmt.Print("Masukkan nomor rekening tujuan: ")
	fmt.Scan(&rTuju)
	fmt.Print("Masukkan jumlah yang akan di transfer: ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan tanggal transaksi (YYYY-MM-DD): ")
	fmt.Scan(&tanggal)

	awal = cariDatabyNoRek(*R, hitungR, rAwal)
	tuju = cariDatabyNoRek(*R, hitungR, rTuju)

	if awal != -1 && tuju != -1 {
		if R[awal].saldo >= jumlah {
			addTransaksi(N, &R[awal], jumlah, tanggal, "debit", hitungN)
			addTransaksi(N, &R[tuju], jumlah, tanggal, "kredit", hitungN)
			fmt.Println("〚 Transaksi berhasil 〛")
		} else {
			fmt.Println("〚 Saldo tidak cukup 〛")
		}
	} else {
		fmt.Println("〚 Akun tidak ditemukkan 〛")
	}
}

func menu() {
	fmt.Println("「 ✦ MENU SISTEM APLIKASI PERBANKAN ✦ 」")
	fmt.Println("    ┌───────────── ❉  ─────────────┐")
	fmt.Println("    |	1. Tambah Rekening         |")
	fmt.Println("    |	2. Tambah Nasabah          |")
	fmt.Println("    | 	3. Cari Rekening           |")
	fmt.Println("    | 	4. Edit Rekening           |")
	fmt.Println("    | 	5. Hapus Rekening          |")
	fmt.Println("    | 	6. Transfer Dana           |")
	fmt.Println("    | 	7. Tampilan Rekening       |")
	fmt.Println("    | 	8. Keluar                  |")
	fmt.Println("    └───────────── ❉  ─────────────┘")
	fmt.Print("Masukkan pilihan Anda: ")
}

func menuC() {
	fmt.Println("「 ✦  MENU CARI DATA REKENING  ✦ 」")
	fmt.Println(" ┌───────────── ❉  ─────────────┐")
	fmt.Println(" | 1. Cari Nomor Rekening       |")
	fmt.Println(" | 2. Cari Cabang Bank          |")
	fmt.Println(" | 3. Cari Transaksi            |")
	fmt.Println(" | 4. Kembali                   |")
	fmt.Println(" └───────────── ❉  ─────────────┘")
	fmt.Print("Masukkan pilihan Anda: ")
}

func main() {
	var R tabRekening
	var N tabNasabah
	var pilih, nR, nN int
	for {
		menu()
		fmt.Scan(&pilih)
		if pilih == 1 {
			addRekening(&R, &nR)

		} else if pilih == 2 {
			if nR == 0 {
				fmt.Println("〚 Tambahkan data rekening terlebih dahulu 〛")
			} else {
				addNasabah(&N, &nN, R, nR)
			}

		} else if pilih == 3 {
			if nR == 0 {
				fmt.Println("〚 Tambahkan data rekening terlebih dahulu 〛")
			} else {
				menuC()
				pilihanCariData(R, nR)
			}

		} else if pilih == 4 {
			if nR == 0 {
				fmt.Println("〚 Tambahkan data rekening terlebih dahulu 〛")
			} else {
				editRekening(&R, nR)
			}

		} else if pilih == 5 {
			if nR == 0 {
				fmt.Println("〚 Tambahkan data rekening terlebih dahulu 〛")
			} else {
				hapusRekening(&R, &nR)
			}

		} else if pilih == 6 {
			if nR == 0 {
				fmt.Println("〚 Tambahkan data rekening terlebih dahulu 〛")
			} else if nR < 2 {
				fmt.Println("〚 Data Rekening kurang 〛")
			} else {
				transaksiSaldo(&N, &R, nR, nN)
			}

		} else if pilih == 7 {
			if nR == 0 {
				fmt.Println("〚 Tambahkan data rekening terlebih dahulu 〛")
			} else {
				cetakRekening(N, R, nR, nN)
			}

		} else if pilih == 8 {
			fmt.Println("Keluar...")
			return

		} else {
			fmt.Printf("〚 Tidak terdapat nomor %d 〛\n", pilih)
		}
	}
}
