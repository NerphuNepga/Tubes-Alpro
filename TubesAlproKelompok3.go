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
	nama  string
	norek string
}

type transaksi struct {
	tanggal string
	jumlah  float64
	tipe    string // "debit" atau "kredit"
}

type tabRekening [NMAX]rekening
type tabNasabah [NMAX]nasabah

func addRekening(R *tabRekening, hitungR *int) { // menambahkan data
	// hitungR adalah jumlah rekening yang ada
	if *hitungR < NMAX {
		// jika jumlah rekening kurang dari NMAX, masukkan data rekening yang akan dibuat
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
		fmt.Println("〚 Data rekening berhasil disimpan 〛")

		//jika sudah selesai maka jumlah rekening (hitungR) bertambah dan menampilkan "data rekening berhasil disimpan"
		*hitungR++
	} else {
		// jika jumlah rekening melebihi NMAX maka akan stop dan menampilkan "data rekening penuh"
		fmt.Println("〚 Data rekening penuh 〛")
	}
}

func addNasabah(N *tabNasabah, hitungN *int, R tabRekening, hitungR int) { //menambahkan data nasabah

	// hitungN adalah jumlah nasabah yang ada
	if *hitungN < NMAX {
		//jika jumlah nasabah kurang dari NMAX, masukkan data nasabah
		fmt.Printf("Masukkan nama untuk nasabah ke-%d: ", *hitungN+1)
		fmt.Scan(&N[*hitungN].nama)
		fmt.Printf("Masukkan nomor rekening yang dimiliki nasabah ke-%d: ", *hitungN+1)
		fmt.Scan(&N[*hitungN].norek)

		//jumlah nasabah bertambah
		*hitungN++

		fmt.Println("〚 Data nasabah berhasil disimpan 〛")
	} else {
		// jika jumlah nasabah melebihi NMAX, maka akan berhenti dan menampilkan "data nasabah penuh"
		fmt.Println("〚 Data nasabah penuh 〛")
	}
}

func pilihanCariData(R tabRekening, hitungR int) {
	var pilih, index, j int
	var xStr string
	var xTrs transaksi

	//kondisi sebelum pencarian
	index = -1

	// masukan pilihan yang terdapat di menuC
	fmt.Scan(&pilih)
	if pilih == 1 {
		fmt.Print("Masukkan nomor rekening yang akan dicari: ")
		fmt.Scan(&xStr)
		index = cariDatabyNoRek(R, hitungR, xStr) //mengupdate nilai index dengan yang ditemukan

		// jika setelah dicari index masih -1, maka akun tidak ketemu
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
		index = cariDatabyTransaksi(R, hitungR, xTrs.tanggal, xTrs.tipe, xTrs.jumlah) //mengupdate nilai index dengan yang ditemukan

		// jika setelah dicari index masih -1, maka akun tidak ketemu
		if index == -1 {
			fmt.Println("〚 Akun tidak ditemukkan 〛")
		}
	} else if pilih == 4 {
		//pilihan kembali ke menu awal
		fmt.Print()
	} else {
		//pilihan tidak sama dengan yang di menuC
		fmt.Println("〚 Pilihan tidak ditemukkan 〛")
	}

	//menampilkan hasil cari tersebut
	if index != -1 {
		//melakukan sorting agar tanggal transaksi terurut
		sortTransaksi(&R[index], R[index].hitungT)
		fmt.Println("——————————————————————————————————————————————————")
		fmt.Printf("Nama              : %s\n", R[index].nama)
		fmt.Printf("Nomor Rekening    : %s\n", R[index].norek)
		fmt.Printf("Tanggal Pembuatan : %s\n", R[index].tanggalBuat)
		fmt.Printf("Cabang Bank       : %s\n", R[index].caBank)
		fmt.Printf("Saldo             : %.2f\n", R[index].saldo)
		fmt.Println("〚 Riwayat Transaksi 〛")

		//jika tidak terdapat riwayat transaksi
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

	//kondisi sebelum pencarian
	found = -1
	i = 0

	//melakukan pencarian hingga ketemu atau sampai semua rekening dicek
	for i < hitungR && found == -1 {
		if R[i].norek == x { // kondisi nomor rekening ditemukan
			found = i
		}
		i++
	}
	return found
}

func cariDatabyCaBank(R tabRekening, hitungR int, x string) {
	var i, j, k int

	k = 1 //k adalah urutan data rekening

	for i = 0; i < hitungR; i++ {
		//mengurutkan tanggal transaksi
		sortTransaksi(&R[i], R[i].hitungT)

		if x == R[i].caBank { // kondisi cabang bank ditemukan
			fmt.Println("——————————————————————————————————————————————————")
			fmt.Printf("Rekening ke-%d\n", k)
			fmt.Printf("Nama              : %s\n", R[i].nama)
			fmt.Printf("Nomor Rekening    : %s\n", R[i].norek)
			fmt.Printf("Tanggal Pembuatan : %s\n", R[i].tanggalBuat)
			fmt.Printf("Cabang Bank       : %s\n", R[i].caBank)
			fmt.Printf("Saldo             : %.2f\n", R[i].saldo)
			fmt.Println("〚 Riwayat Transaksi 〛")

			//jika tidak terdapat riwayat transaksi
			if R[i].hitungT <= 0 {
				fmt.Println("Tidak ada")
			} else {
				//mengeluarkan riwayat transaksi
				for j = 0; j < R[i].hitungT; j++ {
					fmt.Printf("%s        : %s, %.2f\n", R[i].tabTransaksi[j].tanggal, R[i].tabTransaksi[j].tipe, R[i].tabTransaksi[j].jumlah)
				}
			}
			k++
		}
	}
	if k == 1 { //kondisi akun tidak ketemu
		fmt.Println("〚 Akun tidak ditemukkan 〛")
	}
}

func cariDatabyTransaksi(R tabRekening, hitungR int, xTgl, xTip string, xJum float64) int {
	var i, j, found int

	//kondisi sebelum pencarian
	found = -1
	i = 0

	//melakukan pencarian hingga ketemu atau sampai semua rekening dicek
	for i < hitungR && found == -1 {
		j = 0

		//melakukan pencarian hingga ketemu atau sampai semua transaksi dicek
		for j < R[i].hitungT && found == -1 {
			// kondisi transaksi ditemukan
			if R[i].tabTransaksi[j].tanggal == xTgl && R[i].tabTransaksi[j].tipe == xTip && R[i].tabTransaksi[j].jumlah == xJum {
				found = i
			}
			j++
		}
		i++
	}
	return found
}

func editRekening(N *tabNasabah, R *tabRekening, hitungR, hitungN int) {
	//mengedit data rekening
	var index, i, j int
	var namaBaru, norekBaru, tanggalBaru, cabankBaru string
	var norek string

	//menampilkan nama dan nomor rekening yang tersedia
	for i = 0; i < hitungR; i++ {
		fmt.Println("——————————————————————————————————————————————————")
		fmt.Printf("Rekening ke-%d\n", i+1)
		fmt.Printf("Nama              : %s\n", R[i].nama)
		fmt.Printf("Nomor Rekening    : %s\n", R[i].norek)
	}

	fmt.Print("Masukkan nomor rekening yang akan di edit: ")
	fmt.Scan(&norek)

	//mencari nomor rekening yang dimasukkan sebelumnya
	index = cariDatabyNoRek(*R, hitungR, norek)

	//kondisi nomor rekening ditemukan
	if index != -1 {
		//memasukkan data baru
		fmt.Print("Masukkan nama baru: ")
		fmt.Scan(&namaBaru)
		fmt.Print("Masukkan nomor rekening baru: ")
		fmt.Scan(&norekBaru)
		fmt.Print("Masukkan tanggal baru (YYYY-MM-DD): ")
		fmt.Scan(&tanggalBaru)
		fmt.Print("Masukkan cabang bank baru: ")
		fmt.Scan(&cabankBaru)

		//mengganti data Nasabah lama dengan yang baru
		for j = 0; j < hitungN; j++ {
			if N[j].norek == R[index].norek {
				N[j].norek = norekBaru
			}
		}

		//mengganti data rekening lama dengan yang baru
		R[index].nama = namaBaru
		R[index].norek = norekBaru
		R[index].tanggalBuat = tanggalBaru
		R[index].caBank = cabankBaru

		fmt.Println("〚 Akun telah selesai di edit 〛")
	} else { //jika nomor rekening tidak ditemukan
		fmt.Println("〚 Akun tidak ditemukan 〛")
	}
}

func hapusRekening(R *tabRekening, hitungR *int) {
	//menghapus data rekening
	var norek string
	var index, i, j int

	//menampilkan nama dan nomor rekening yang tersedia
	for j = 0; j < *hitungR; j++ {
		fmt.Println("——————————————————————————————————————————————————")
		fmt.Printf("Rekening ke-%d\n", j+1)
		fmt.Printf("Nama              : %s\n", R[j].nama)
		fmt.Printf("Nomor Rekening    : %s\n", R[j].norek)
	}

	fmt.Print("Masukkan nomor rekening yang akan di hapus: ")
	fmt.Scan(&norek)

	//mencari nomor rekening yang dimasukkan
	index = cariDatabyNoRek(*R, *hitungR, norek)

	//kondisi nomor rekening ditemukan
	if index != -1 {
		for i = index; i < *hitungR-1; i++ {
			//mengupdate nilai index i dengan index berikutnya
			R[i] = R[i+1]
		}
		//mengurangi jumlah rekening
		*hitungR--
		fmt.Println("〚 Akun telah dihapus 〛")
	} else {
		fmt.Println("〚 Akun tidak ditemukan 〛")
	}
}

func hapusTransaksi(R *tabRekening, hitungR int) {	// Menghapus data transaksi pada sebuah akun rekening
	var norek, tgl string
	var index, i, j, k, l int
	var found bool

	// Menampilkan semua rekening yang ada
	for i = 0; i < hitungR; i++ {
		fmt.Println("——————————————————————————————————————————————————")
		fmt.Printf("Rekening ke-%d\n", i+1)
		fmt.Printf("Nama              : %s\n", R[i].nama)
		fmt.Printf("Nomor Rekening    : %s\n", R[i].norek)
	}

	fmt.Print("Masukkan nomor rekening yang akan di hapus transaksinya: ")
	fmt.Scan(&norek)

	// Mencari indeks rekening berdasarkan nomor rekening
	index = cariDatabyNoRek(*R, hitungR, norek)

	// Jika rekening ditemukan
	if index != -1 {
		fmt.Println("——————————————————————————————————————————————————")
		// Menampilkan semua transaksi dari rekening yang dipilih
		for j = 0; j < R[index].hitungT; j++ {
			fmt.Printf("%s        : %s, %.2f\n", R[index].tabTransaksi[j].tanggal, R[index].tabTransaksi[j].tipe, R[index].tabTransaksi[j].jumlah)
		}

		fmt.Print("Masukkan tanggal transaksi yang akan di hapus (YYYY-MM-DD): ")
		fmt.Scan(&tgl)

		found = false
		// Loop untuk mencari transaksi berdasarkan tanggal
		for k = 0; k < R[index].hitungT; k++ {
			if R[index].tabTransaksi[k].tanggal == tgl {
				// Menggeser transaksi untuk menghapus transaksi yang ditemukan
				for l = k; l < R[index].hitungT-1; l++ {
					R[index].tabTransaksi[l] = R[index].tabTransaksi[l+1]
				}
				// Jumlah transaksi berkurang
				R[index].hitungT--
				found = true
			}
		}
		if found {
			fmt.Println("〚 Transaksi telah dihapus 〛")
		} else {
			fmt.Println("〚 Transaksi tidak ditemukan 〛")
		}
	} else {
		fmt.Println("〚 Akun tidak ditemukan 〛")
	}
}

func sortTanggal(R *tabRekening, hitungR int) {
	//mengurutkan tanggal dari data rekening
	var i, pass int
	var temp rekening

	pass = 1
	for pass < hitungR {
		temp = R[pass]
		i = pass

		//kondisi tanggal yang dapat diurutkan
		for i > 0 && temp.tanggalBuat < R[i-1].tanggalBuat {
			//mengupdate nilai index i menjadi index sebelumnya
			R[i] = R[i-1]
			i = i - 1
		}

		//melakukan penukaran index
		R[i] = temp
		pass = pass + 1
	}
}

func sortTransaksi(R *rekening, hitungT int) {
	//mengurutkan tanggal dari data transaksi
	var i, pass int
	var temp transaksi

	pass = 1
	for pass < hitungT {
		temp = R.tabTransaksi[pass]
		i = pass

		//kondisi tanggal yang dapat diurutkan
		for i > 0 && temp.tanggal < R.tabTransaksi[i-1].tanggal {
			//mengupdate nilai index i menjadi index sebelumnya
			R.tabTransaksi[i] = R.tabTransaksi[i-1]
			i = i - 1
		}

		//melakukan penukaran index
		R.tabTransaksi[i] = temp
		pass = pass + 1
	}
}

func cetakRekening(N tabNasabah, R tabRekening, hitungR, hitungN int) {
	//mengeluarkan data nasabah dan data rekening yang sudah terdaftar
	var i, j, k int

	//mengurutkan tanggal data rekening
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

		//jika tidak terdapat riwayat transaksi
		if R[i].hitungT <= 0 {
			fmt.Println("Tidak ada")
		} else {
			//mengeluarkan riwayat transaksi
			for k = 0; k < R[i].hitungT; k++ {
				fmt.Printf("%s        : %s, %.2f\n", R[i].tabTransaksi[k].tanggal, R[i].tabTransaksi[k].tipe, R[i].tabTransaksi[k].jumlah)
			}
		}
	}
}

func addTransaksi(R *rekening, jumlah float64, tanggal, tipe string) { //melakukan proses transfer

	//hitungT adalah jumlah transaksi yang telah dilakukan
	if R.hitungT < NMAX {

		//mengisi data transaksi; menambahkan jumlah transaksi
		R.tabTransaksi[R.hitungT].tanggal = tanggal
		R.tabTransaksi[R.hitungT].tipe = tipe
		R.tabTransaksi[R.hitungT].jumlah = jumlah
		R.hitungT++

		if tipe == "debit" { //jika debit maka saldo berkurang
			R.saldo = R.saldo - jumlah
		} else if tipe == "kredit" { //jika kredit maka saldo bertambah
			R.saldo = R.saldo + jumlah
		}

	} else { //jika jumlah transaksi melebihi NMAX
		fmt.Println("〚 Data transaksi penuh 〛")
	}
}

func transaksiSaldo(R *tabRekening, hitungR int) { //melakukan transfer antar akun
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

	//mencari nomor rekening yang akan melakukan proses transfer
	awal = cariDatabyNoRek(*R, hitungR, rAwal)
	tuju = cariDatabyNoRek(*R, hitungR, rTuju)

	//jika akun ditemukan, melanjutkan ke proses transfer
	if awal != -1 && tuju != -1 {
		if R[awal].saldo >= jumlah {
			//proses transfer
			addTransaksi(&R[awal], jumlah, tanggal, "debit")
			addTransaksi(&R[tuju], jumlah, tanggal, "kredit")
			fmt.Println("〚 Transaksi berhasil 〛")
		} else {
			fmt.Println("〚 Saldo tidak cukup 〛")
		}
	} else {
		fmt.Println("〚 Akun tidak ditemukkan 〛")
	}
}

func menu() { //menu awal
	fmt.Println("「 ✦ MENU SISTEM APLIKASI PERBANKAN ✦ 」")
	fmt.Println("    ┌───────────── ❉  ─────────────┐")
	fmt.Println("    |	1. Tambah Rekening         |")
	fmt.Println("    |	2. Tambah Nasabah          |")
	fmt.Println("    | 	3. Cari Rekening           |")
	fmt.Println("    | 	4. Edit Rekening           |")
	fmt.Println("    | 	5. Hapus Rekening          |")
	fmt.Println("    | 	6. Hapus Transaksi         |")
	fmt.Println("    | 	7. Transfer Dana           |")
	fmt.Println("    | 	8. Tampilan Rekening       |")
	fmt.Println("    | 	9. Keluar                  |")
	fmt.Println("    └───────────── ❉  ─────────────┘")
	fmt.Print("Masukkan pilihan Anda: ")
}

func menuC() { //menu cari data
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

	//perulangan sampai pilihan ke-8 terpilih
	for pilih != 9 {
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
				editRekening(&N, &R, nR, nN)
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
			} else {
				hapusTransaksi(&R, nR)
			}

		} else if pilih == 7 {
			if nR == 0 {
				fmt.Println("〚 Tambahkan data rekening terlebih dahulu 〛")
			} else if nR < 2 {
				fmt.Println("〚 Data Rekening kurang 〛")
			} else {
				transaksiSaldo(&R, nR)
			}

		} else if pilih == 8 {
			if nR == 0 {
				fmt.Println("〚 Tambahkan data rekening terlebih dahulu 〛")
			} else {
				cetakRekening(N, R, nR, nN)
			}

		} else if pilih == 9 {
			fmt.Println("Keluar...")

		} else {
			fmt.Printf("〚 Tidak terdapat pilihan ke-%d 〛\n", pilih)
		}
	}
}
