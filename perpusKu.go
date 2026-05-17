package main

import "fmt"

const NMAX = 100

type Buku struct {
	id string
	judul string
	kategori string
	penulis string
	tahunTerbit int
	status bool 
}

type tabBuku [NMAX] Buku


func main () {
	var koleksi tabBuku
	var jumlah int = 0 
	var pilihan int
	jumlah = 0 	

	for {
		fmt.Println("\n---------------------")
		fmt.Println("  SiPerpus - Katalog Digital")
		fmt.Println("---------------------")
		fmt.Println("1. Tambah Buku")
		fmt.Println("2. Tampilkan Semua Buku")
		fmt.Println("3. Ubah Data Buku")
		fmt.Println("4. Hapus Buku")
		fmt.Println("5. Cari Buku Berdasarkan Judul")
		fmt.Println("6. Cari Buku Berdasarkan ID")
		fmt.Println("7. Urutan Tahun Terbit - Ascending")
		fmt.Println("8. Urutan Tahun Terbit - Descending")
		fmt.Println("9. Tampilkan Statistik")
		fmt.Println("0. Keluar Aplikasi")
		fmt.Print("Pilih menu: ")	
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahBuku(&koleksi, &jumlah)
		case 2:
			tampilkanSemuaBuku(&koleksi, &jumlah)
		case 3:
			ubahBuku(&koleksi, &jumlah)
		case 4:
			hapusBuku(&koleksi, &jumlah)
		case 5:
			searchJudul()
		case 6:
			searchId()
		case 7:
			selectionSortTahunTerbit()
		case 8:
			insertionSortTahunTerbit()
		case 9:
			statistik()
		case 0:
			fmt.Println("Terima kasih sudah menggunakan perpusKu")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}

}


func tambahBuku(arr *tabBuku, n *int) {
	if *n >= NMAX {
		fmt.Println("Maaf, Kapasitas Perpustakaan Penuh")
		return 
	}
	var buku Buku
	fmt.Println("--- Tambah Buku Baru ---")
	fmt.Print("ID Buku: ")
	fmt.Scan(&buku.id)
	fmt.Print("Judul: ")
	fmt.Scan(&buku.judul)
	fmt.Print("Penulis: ")
	fmt.Scan(&buku.penulis)
	fmt.Print("Kategori: ")
	fmt.Scan(&buku.kategori)
	fmt.Print("Tahun Terbit: ")
	fmt.Scan(&buku.tahunTerbit)
	buku.status = true

	arr[*n] = buku
	*n++
	fmt.Println("Data buku berhasil ditambahkan")
}

func tampilkanSemuaBuku(arr *tabBuku, n *int) {
	var i int
	var buku Buku
	if *n == 0 {
		fmt.Println("Koleksi anda masih kosong")
		return
	}

	fmt.Println("\n-----DAFTAR KOLEKSI BUKU -----")
	for i=0; i < *n; i++ {
		buku = arr[i]
		fmt.Printf("[%d] ID: %s | Judul: %s | Penulis: %s | Kategori: %s | Tahun: %d | Status: %s\n",i + 1, buku.id, buku.judul, buku.penulis, buku.kategori, buku.tahunTerbit, buku.status)
	}
}

func ubahBuku(arr *tabBuku, n *int) {
	var id string
	var idx,i int
	var buku Buku
	idx = -1

	fmt.Println("Masukkan ID buku yang ingin anda edit: ")
	fmt.Scan(&id)

	for i = 0; i < *n; i++{
		if arr[i].id == id {
			idx = i 
			break	
		}
	}
	
	if idx == -1 {
		fmt.Println("Buku tidak ditemukan")
		return
	}
	fmt.Printf("Data Buku : Judul: %s | Penulis: %s | Kategori: %s | Tahun: %d | Status: %s\n", arr[idx].judul, arr[idx].penulis, arr[idx].kategori, arr[idx].tahunTerbit, arr[idx].status)

	buku.id = id
	fmt.Print("Judul         :")
	fmt.Scan(&buku.judul)
	fmt.Print("Penulis       :")
	fmt.Scan(&buku.penulis)
	fmt.Print("Kategori   	 :")
	fmt.Scan(&buku.kategori)
	fmt.Print("Tahun Terbit  :")
	fmt.Scan(&buku.tahunTerbit)
	fmt.Print("Status Baru   :")
	fmt.Scan(&buku.status)
}	

func hapusBuku(arr *tabBuku, n *int) {
	var id string
	var idx,i int
	idx= -1
	
	fmt.Print("Masukkan ID Buku yang ingin dihapus: ")
	fmt.Scan(&id)

	for i = 0; i < *n; i++ {
		if arr[i].id == id {
			idx = i 
			break
		}
	}


	if idx == -1 {
		fmt.Println("Buku tidak ditemukan")
		return
	}

	fmt.Printf("Menghapus -> ID: %s  | Judul:%s  | Penulis: %s.  | Tahun Terbit: %d\n", arr[idx].id, arr[idx].judul, arr[idx].penulis, arr[idx].tahunTerbit)

	for i = idx; i < *n - 1; i++ {
		arr[i] = arr[i+1]
	}
	*n--
	fmt.Println("Buku Berhasil Dihapus")
	
}

func searchJudul() {

}
func searchId() {

}

func selectionSortTahunTerbit() {

}

func insertionSortTahunTerbit() {

}


func sort() {

}

func statistik() {

}