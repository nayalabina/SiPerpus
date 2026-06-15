package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Buku struct {
	id          int
	judul       string
	penulis     string
	kategori    string
	tahunTerbit int
	status      string
}

var dataBuku [100]Buku
var jumlahData int

var reader = bufio.NewReader(os.Stdin)

func isiDataDummy() {

	dataBuku[0] = Buku{1, "Harry Potter", "J.K. Rowling", "Fantasi", 1997, "tersedia"}
	dataBuku[1] = Buku{2, "Laskar Pelangi", "Andrea Hirata", "Novel", 2005, "dipinjam"}
	dataBuku[2] = Buku{3, "Bumi", "Tere Liye", "Fantasi", 2014, "tersedia"}
	dataBuku[3] = Buku{4, "Atomic Habits", "James Clear", "Pengembangan Diri", 2018, "tersedia"}
	dataBuku[4] = Buku{5, "Filosofi Teras", "Henry Manampiring", "Psikologi", 2018, "dipinjam"}
	dataBuku[5] = Buku{6, "Negeri 5 Menara", "Ahmad Fuadi", "Novel", 2009, "tersedia"}
	dataBuku[6] = Buku{7, "Dilan 1990", "Pidi Baiq", "Romantis", 2014, "dipinjam"}
	dataBuku[7] = Buku{8, "Cantik Itu Luka", "Eka Kurniawan", "Sastra", 2002, "tersedia"}
	dataBuku[8] = Buku{9, "Ayat Ayat Cinta", "Habiburrahman", "Religi", 2004, "tersedia"}
	dataBuku[9] = Buku{10, "Bulan", "Tere Liye", "Fantasi", 2015, "dipinjam"}
	dataBuku[10] = Buku{11, "Pergi", "Tere Liye", "Aksi", 2018, "tersedia"}
	dataBuku[11] = Buku{12, "Pulang", "Tere Liye", "Aksi", 2015, "tersedia"}
	dataBuku[12] = Buku{13, "Rindu", "Tere Liye", "Religi", 2014, "dipinjam"}
	dataBuku[13] = Buku{14, "Mariposa", "Luluk HF", "Romantis", 2018, "tersedia"}
	dataBuku[14] = Buku{15, "Koala Kumal", "Raditya Dika", "Komedi", 2015, "tersedia"}
	dataBuku[15] = Buku{16, "Rich Dad Poor Dad", "Robert Kiyosaki", "Bisnis", 1997, "dipinjam"}
	dataBuku[16] = Buku{17, "The Psychology of Money", "Morgan Housel", "Keuangan", 2020, "tersedia"}
	dataBuku[17] = Buku{18, "Think and Grow Rich", "Napoleon Hill", "Motivasi", 1937, "tersedia"}
	dataBuku[18] = Buku{19, "Deep Work", "Cal Newport", "Produktivitas", 2016, "dipinjam"}
	dataBuku[19] = Buku{20, "Clean Code", "Robert Martin", "Teknologi", 2008, "tersedia"}

	jumlahData = 20
}

func cekIDDuplikat(idBaru int) bool {
	//  cek id biar gada yg duplikat karna idbuku harus beda

	var i int

	for i = 0; i < jumlahData; i++ {

		if dataBuku[i].id == idBaru {
			return true
		}
	}

	return false
}

func inputTahun() int { // validasi input jadi harus angka gaboleh string
	var tahunString string
	var tahun int
	var valid bool
	var i int

	for valid == false {

		valid = true

		fmt.Print("Tahun Terbit: ")
		tahunString, _ = reader.ReadString('\n')
		tahunString = strings.TrimSpace(tahunString)

		if len(tahunString) == 0 {
			valid = false
		}

		for i = 0; i < len(tahunString); i++ {

			if tahunString[i] < '0' || tahunString[i] > '9' {
				valid = false
			}
		}

		if valid == false {
			fmt.Println("Input tahun harus angka!")
		}
	}

	tahun = 0

	for i = 0; i < len(tahunString); i++ {
		tahun = tahun*10 + int(tahunString[i]-'0')
	}

	return tahun
}

func tambahBuku() {
	// tambah data buku baru (eel)

	if jumlahData >= 100 { // kondisi jika buku sdh diinput 100, perlu ditanyakan dl
		fmt.Println("Kapasitas buku penuh")
		return
	}

	var input string
	var idBaru int

	for {

		fmt.Print("ID Buku: ")
		fmt.Scan(&idBaru)

		if idBaru <= 0 {
			fmt.Println("ID harus lebih dari 0")
		} else if cekIDDuplikat(idBaru) {
			fmt.Println("ID sudah digunakan")
		} else {
			break
		}
	}

	dataBuku[jumlahData].id = idBaru

	reader.ReadString('\n')

	fmt.Print("Judul Buku: ")
	input, _ = reader.ReadString('\n')
	dataBuku[jumlahData].judul = strings.TrimSpace(input)

	fmt.Print("Penulis: ")
	input, _ = reader.ReadString('\n')
	dataBuku[jumlahData].penulis = strings.TrimSpace(input)

	fmt.Print("Kategori: ")
	input, _ = reader.ReadString('\n')
	dataBuku[jumlahData].kategori = strings.TrimSpace(input)

	dataBuku[jumlahData].tahunTerbit = inputTahun()

	for {

		fmt.Print("Status Buku (tersedia/dipinjam): ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "tersedia" ||
			strings.ToLower(input) == "dipinjam" {

			dataBuku[jumlahData].status = input
			break
		}

		fmt.Println("Status hanya boleh tersedia atau dipinjam")
	}

	jumlahData++

	fmt.Println("Data buku berhasil ditambahkan")
}

func tampilBuku() {
	// nampilin seluruh data buku (@jebb_24)

	var i int

	if jumlahData == 0 {
		fmt.Println("Belum ada data buku")
	} else {

		for i = 0; i < jumlahData; i++ {

			fmt.Println("========================")
			fmt.Println("ID Buku      :", dataBuku[i].id)
			fmt.Println("Judul        :", dataBuku[i].judul)
			fmt.Println("Penulis      :", dataBuku[i].penulis)
			fmt.Println("Kategori     :", dataBuku[i].kategori)
			fmt.Println("Tahun Terbit :", dataBuku[i].tahunTerbit)

			if strings.ToLower(dataBuku[i].status) == "dipinjam" {
				fmt.Println("Status       : Sedang Dipinjam")
			} else {
				fmt.Println("Status       : Tersedia")
			}
		}
	}
}

func sequentialSearch() {
	// sequential search cari judul buku (eel)

	var cari string
	var i int
	var ketemu bool

	fmt.Print("Masukkan judul buku: ")
	cari, _ = reader.ReadString('\n')
	cari = strings.TrimSpace(cari)

	for i = 0; i < jumlahData; i++ {

		if strings.ToLower(dataBuku[i].judul) == strings.ToLower(cari) { // untuk cari tipe string jd bs capslock atau ga

			fmt.Println("Data ditemukan")
			fmt.Println("========================")
			fmt.Println("ID Buku      :", dataBuku[i].id)
			fmt.Println("Judul        :", dataBuku[i].judul)
			fmt.Println("Penulis      :", dataBuku[i].penulis)
			fmt.Println("Kategori     :", dataBuku[i].kategori)
			fmt.Println("Tahun Terbit :", dataBuku[i].tahunTerbit)
			fmt.Println("Status       :", dataBuku[i].status)

			ketemu = true
		}
	}

	if ketemu == false {
		fmt.Println("Data tidak ditemukan")
	}
}

func selectionSort() {
	// selection sort ascending tahun terbit (@jebb_24)

	var i, j, min int
	var temp Buku

	for i = 0; i < jumlahData-1; i++ {

		min = i

		for j = i + 1; j < jumlahData; j++ {

			if dataBuku[j].tahunTerbit < dataBuku[min].tahunTerbit {
				min = j
			}
		}

		temp = dataBuku[i]
		dataBuku[i] = dataBuku[min]
		dataBuku[min] = temp
	}

	fmt.Println("Data berhasil diurutkan ascending")
}

func insertionSort() {
	// insertion sort descending tahun terbit (eel)

	var i, j int
	var temp Buku

	for i = 1; i < jumlahData; i++ {

		temp = dataBuku[i]
		j = i - 1

		for j >= 0 && dataBuku[j].tahunTerbit < temp.tahunTerbit {

			dataBuku[j+1] = dataBuku[j]
			j--
		}

		dataBuku[j+1] = temp
	}

	fmt.Println("Data berhasil diurutkan descending")
}

func urutkanID() { // urutkan data berdasarkan id terlebih dahulu
	var i, j int
	var temp Buku

	for i = 0; i < jumlahData-1; i++ {
		for j = i + 1; j < jumlahData; j++ {

			if dataBuku[i].id > dataBuku[j].id {

				temp = dataBuku[i]
				dataBuku[i] = dataBuku[j]
				dataBuku[j] = temp
			}
		}
	}
}

func binarySearch() {
	// binary search berdasarkan ID buku (bisa digunakan pada data yg sdh urut)

	urutkanID()

	var cari int
	var kiri, kanan, tengah int
	var found bool

	fmt.Print("Masukkan ID Buku: ")
	fmt.Scan(&cari)

	kiri = 0
	kanan = jumlahData - 1

	for kiri <= kanan && found == false {

		tengah = (kiri + kanan) / 2

		if dataBuku[tengah].id == cari {

			found = true

			fmt.Println("Data ditemukan")
			fmt.Println("========================")
			fmt.Println("ID Buku      :", dataBuku[tengah].id)
			fmt.Println("Judul        :", dataBuku[tengah].judul)
			fmt.Println("Penulis      :", dataBuku[tengah].penulis)
			fmt.Println("Kategori     :", dataBuku[tengah].kategori)
			fmt.Println("Tahun Terbit :", dataBuku[tengah].tahunTerbit)
			fmt.Println("Status       :", dataBuku[tengah].status)

		} else if cari > dataBuku[tengah].id {

			kiri = tengah + 1

		} else {

			kanan = tengah - 1
		}
	}

	if found == false {
		fmt.Println("Data tidak ditemukan")
	}
}

func editBuku() {
	// mengubah data buku berdasarkan ID (eel)

	var idEdit int
	var i int
	var input string
	var ketemu bool

	fmt.Print("Masukkan ID Buku yang ingin diedit: ")
	fmt.Scan(&idEdit)

	reader.ReadString('\n')

	for i = 0; i < jumlahData; i++ {

		if dataBuku[i].id == idEdit {

			fmt.Print("Judul Baru: ")
			input, _ = reader.ReadString('\n')
			dataBuku[i].judul = strings.TrimSpace(input)

			fmt.Print("Penulis Baru: ")
			input, _ = reader.ReadString('\n')
			dataBuku[i].penulis = strings.TrimSpace(input)

			fmt.Print("Kategori Baru: ")
			input, _ = reader.ReadString('\n')
			dataBuku[i].kategori = strings.TrimSpace(input)

			dataBuku[i].tahunTerbit = inputTahun()

			for {

				fmt.Print("Status Baru (tersedia/dipinjam): ")
				input, _ = reader.ReadString('\n')
				input = strings.TrimSpace(input)

				if strings.ToLower(input) == "tersedia" ||
					strings.ToLower(input) == "dipinjam" {

					dataBuku[i].status = input
					break
				}

				fmt.Println("Status hanya boleh tersedia atau dipinjam")
			}

			ketemu = true

			fmt.Println("Data berhasil diubah")
			break
		}
	}

	if ketemu == false {
		fmt.Println("Data tidak ditemukan")
	}
}

func hapusBuku() {
	// menghapus data buku berdasarkan ID (@jebb_24)

	var idHapus int
	var i, j int
	var ketemu bool

	fmt.Print("Masukkan ID Buku yang ingin dihapus: ")
	fmt.Scan(&idHapus)

	for i = 0; i < jumlahData; i++ {

		if dataBuku[i].id == idHapus {

			ketemu = true

			for j = i; j < jumlahData-1; j++ {

				dataBuku[j] = dataBuku[j+1]
			}

			jumlahData--

			fmt.Println("Data berhasil dihapus")
			break
		}
	}

	if ketemu == false {
		fmt.Println("Data tidak ditemukan")
	}
}

func statistikBuku() {
	// menghitung statistik buku dan kategori (eel)

	var tersedia int
	var dipinjam int
	var i, j int

	var kategori [100]string
	var jumlahKategori [100]int
	var banyakKategori int
	var ditemukan bool

	for i = 0; i < jumlahData; i++ {

		if strings.ToLower(dataBuku[i].status) == "dipinjam" {
			dipinjam++
		} else {
			tersedia++
		}

		ditemukan = false

		for j = 0; j < banyakKategori; j++ {

			if kategori[j] == dataBuku[i].kategori {
				jumlahKategori[j]++
				ditemukan = true
			}
		}

		if ditemukan == false {
			kategori[banyakKategori] = dataBuku[i].kategori
			jumlahKategori[banyakKategori] = 1
			banyakKategori++
		}
	}

	fmt.Println("===== STATISTIK =====")
	fmt.Println("Total Buku     :", jumlahData)
	fmt.Println("Buku Tersedia  :", tersedia)
	fmt.Println("Buku Dipinjam  :", dipinjam)

	fmt.Println()
	fmt.Println("Jumlah Buku per Kategori")

	for i = 0; i < banyakKategori; i++ {
		fmt.Println(kategori[i], ":", jumlahKategori[i])
	}
}

func main() {

	var pilihan int = -1
	var eel int
	var jebb_24 string

	eel = 24
	jebb_24 = "SiPerpus"

	fmt.Println("Selamat datang di", jebb_24)
	fmt.Println("Kode bonus:", eel)

	isiDataDummy()

	for pilihan != 0 {

		fmt.Println()
		fmt.Println("===== MENU SIPERPUS =====")
		fmt.Println("1. Tambah Buku")
		fmt.Println("2. Tampilkan Buku")
		fmt.Println("3. Sequential Search")
		fmt.Println("4. Binary Search")
		fmt.Println("5. Selection Sort")
		fmt.Println("6. Insertion Sort")
		fmt.Println("7. Statistik Buku")
		fmt.Println("8. Edit Buku")
		fmt.Println("9. Hapus Buku")
		fmt.Println("0. Keluar")

		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		reader.ReadString('\n')

		switch pilihan {

		case 1:
			tambahBuku()

		case 2:
			tampilBuku()

		case 3:
			sequentialSearch()

		case 4:
			binarySearch()

		case 5:
			selectionSort()

		case 6:
			insertionSort()

		case 7:
			statistikBuku()

		case 8:
			editBuku()

		case 9:
			hapusBuku()

		case 0:
			fmt.Println("Program selesai")

		default:
			fmt.Println("Menu tidak tersedia")
		}
	}
}
