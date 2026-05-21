package main

import "fmt"

const NMAX = 100

type Buku struct {
	id          string
	judul       string
	kategori    string
	penulis     string
	tahunTerbit int
	status      string
}

type tabBuku [NMAX]Buku

func main() {
	var koleksi tabBuku
	var jumlahBuku int = 0
	var pilih int
	jumlahBuku = 0

	for {
		fmt.Println("--------------------------")
		fmt.Println("SiPerpus - Katalog Digital")
		fmt.Println("--------------------------")
		fmt.Println("1. Tambah Buku")
		fmt.Println("2. Tampilkan Semua Buku")
		fmt.Println("3. Ubah Data Buku")
		fmt.Println("4. Hapus Buku")
		fmt.Println("5. Cari Buku Berdasarkan Judul")
		fmt.Println("6. Cari Buku Berdasarkan ID")
		fmt.Println("7. Urutan Tahun Terbit - Ascending - Seletion Sort")
		fmt.Println("8. Urutan Tahun Terbit - Descending - Insertion Sort")
		fmt.Println("9. Tampilkan Statistik")
		fmt.Println("0. Keluar Aplikasi")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahBuku(&koleksi, &jumlahBuku)
		case 2:
			tampilkanSemuaBuku(&koleksi, &jumlahBuku)
		case 3:
			ubahBuku(&koleksi, &jumlahBuku)
		case 4:
			hapusBuku(&koleksi, &jumlahBuku)
		case 5:
			searchJudul(&koleksi, &jumlahBuku)
		case 6:
			searchId(&koleksi, &jumlahBuku)
		case 7:
			selectionSortTahunTerbit(&koleksi, &jumlahBuku)
		case 8:
			insertionSortTahunTerbit(&koleksi, &jumlahBuku)
		case 9:
			statistik(&koleksi, &jumlahBuku)
		case 0:
			fmt.Println("Terima kasih sudah menggunakan perpusKu")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}

}

func tambahBuku(koleksi *tabBuku, jumlahBuku *int) {
	if *jumlahBuku >= NMAX {
		fmt.Println("Maaf, Kapasitas Perpustakaan Penuh")
		return
	}
	var buku Buku
	fmt.Println("--- Tambah Buku Baru ---")
	fmt.Print("ID Buku      : ")
	fmt.Scan(&buku.id)
	fmt.Print("Judul        : ")
	fmt.Scan(&buku.judul)
	fmt.Print("Penulis      : ")
	fmt.Scan(&buku.penulis)
	fmt.Print("Kategori     : ")
	fmt.Scan(&buku.kategori)
	fmt.Print("Tahun Terbit : ")
	fmt.Scan(&buku.tahunTerbit)
	buku.status = "Tersedia"

	koleksi[*jumlahBuku] = buku
	*jumlahBuku++
	fmt.Println("Data buku berhasil ditambahkan")
}

func tampilkanSemuaBuku(koleksi *tabBuku, jumlahBuku *int) {
	var i int
	var buku Buku
	if *jumlahBuku == 0 {
		fmt.Println("Koleksi anda masih kosong")
		return
	}

	fmt.Println("----DAFTAR KOLEKSI BUKU -----")
	for i = 0; i < *jumlahBuku; i++ {
		buku = koleksi[i]
		fmt.Printf("[%d] ID: %s | Judul: %s | Penulis: %s | Kategori: %s | Tahun: %d | Status: %s\n", i+1, buku.id, buku.judul, buku.penulis, buku.kategori, buku.tahunTerbit, buku.status)
	}
}

func ubahBuku(koleksi *tabBuku, jumlahBuku *int) {
	var id string
	var idx, i int
	var buku Buku
	idx = -1

	fmt.Println("Masukkan ID buku yang ingin anda edit: ")
	fmt.Scan(&id)

	for i = 0; i < *jumlahBuku; i++ {
		if koleksi[i].id == id {
			idx = i
			break
		}
	}

	if idx == -1 {
		fmt.Println("Buku tidak ditemukan")
		return
	}
	fmt.Printf("Data Buku : Judul: %s | Penulis: %s | Kategori: %s | Tahun: %d | Status: %s\n", koleksi[idx].judul, koleksi[idx].penulis, koleksi[idx].kategori, koleksi[idx].tahunTerbit, koleksi[idx].status)

	buku.id = id
	fmt.Print("Judul         :")
	fmt.Scan(&buku.judul)
	fmt.Print("Penulis       :")
	fmt.Scan(&buku.penulis)
	fmt.Print("Kategori   	 :")
	fmt.Scan(&buku.kategori)
	fmt.Print("Tahun Terbit  :")
	fmt.Scan(&buku.tahunTerbit)
	fmt.Print("Status (Tersedia/Tidak_Tersedia): ")
	fmt.Scan(&buku.status)
	koleksi[idx] = buku
	fmt.Println("Buku Berhasil diubah")
}

func hapusBuku(koleksi *tabBuku, jumlahBuku *int) {
	var id string
	var idx, i int
	idx = -1

	fmt.Print("Masukkan ID Buku yang ingin dihapus: ")
	fmt.Scan(&id)

	for i = 0; i < *jumlahBuku; i++ {
		if koleksi[i].id == id {
			idx = i
			break
		}
	}

	if idx == -1 {
		fmt.Println("Buku tidak ditemukan")
		return
	}

	fmt.Printf("Menghapus -> ID: %s  | Judul:%s  | Penulis: %s.  | Tahun Terbit: %d\n", koleksi[idx].id, koleksi[idx].judul, koleksi[idx].penulis, koleksi[idx].tahunTerbit)

	for i = idx; i < *jumlahBuku-1; i++ {
		koleksi[i] = koleksi[i+1]
	}
	*jumlahBuku--
	fmt.Println("Buku Berhasil Dihapus")

}

func searchJudul(koleksi *tabBuku, jumlahBuku *int) {
	var i int
	var ketemu bool
	var find string
	fmt.Print("Masukkan Judul Buku: ")
	fmt.Scan(&find)
	for i = 0; i < *jumlahBuku; i++ {
		if koleksi[i].judul == find {
			fmt.Println("----------------------")
			fmt.Println("    Buku Ditemukan    ")
			fmt.Println("----------------------")
			fmt.Printf("ID           : %s\n", koleksi[i].id)
			fmt.Printf("Judul        : %s\n", koleksi[i].judul)
			fmt.Printf("Kategori     : %s\n", koleksi[i].kategori)
			fmt.Printf("Penulis      : %s\n", koleksi[i].penulis)
			fmt.Printf("Tahun Terbit : %d\n", koleksi[i].tahunTerbit)
			fmt.Printf("Status       : %s\n", koleksi[i].status)
			ketemu = true
		}
	}
	if ketemu == false {
		fmt.Println("Buku Tidak Ditemukan")
	}
}

func searchId(koleksi *tabBuku, jumlahBuku *int) {
	var find string
	var kiri, kanan, mid int
	fmt.Print("Masukkan ID Buku: ")
	fmt.Scan(&find)

	kiri= 0
	kanan = *jumlahBuku - 1

	for kiri <= kanan {
		mid = (kiri + kanan) / 2
		if koleksi[mid].id == find {
			fmt.Println("----------------------")
			fmt.Println("    Buku Ditemukan    ")
			fmt.Println("----------------------")
			fmt.Printf("ID           : %s\n", koleksi[mid].id)
			fmt.Printf("Judul        : %s\n", koleksi[mid].judul)
			fmt.Printf("Kategori     : %s\n", koleksi[mid].kategori)
			fmt.Printf("Penulis      : %s\n", koleksi[mid].penulis)
			fmt.Printf("Tahun Terbit : %d\n", koleksi[mid].tahunTerbit)
			fmt.Printf("Status       : %s\n", koleksi[mid].status)
			return
		}else if find < koleksi[mid].id {
			kanan = mid - 1
		}else{
			kiri = mid + 1
		}
	}

	fmt.Println("Buku tidak ditemukan")

}

func selectionSortTahunTerbit(koleksi *tabBuku, jumlahBuku *int) {
	var i, minIdx, N int
	var temp Buku 

	for N = 0; N < *jumlahBuku - 1; N++{ 
		minIdx = N
		for i = N + 1; i < * jumlahBuku;i++{
			if koleksi[i].tahunTerbit < koleksi[minIdx].tahunTerbit {
				minIdx = i
			}
		}
		if minIdx != N {
			temp  = koleksi[minIdx]
			koleksi[minIdx] = koleksi[N]
			koleksi[N] = temp
		}
	}
}

func insertionSortTahunTerbit(koleksi *tabBuku, jumlahBuku *int) {
	var pass, i int
    var temp Buku
    pass = 1
    for pass < *jumlahBuku {
        i = pass
        temp = koleksi[pass]
        for i > 0 && temp.tahunTerbit > koleksi[i-1].tahunTerbit { // ← > bukan 
            koleksi[i] = koleksi[i-1]
            i = i - 1
        }
        koleksi[i] = temp
        pass = pass + 1
    }
    fmt.Println("Data Berhasil Diurutkan (Descending)")
    tampilkanSemuaBuku(koleksi, jumlahBuku)
}

func statistik(koleksi *tabBuku, jumlahBuku *int) {
    var i, j int
    var totalTersedia int
    var kategoriList [NMAX]string
    var kategoriCount [NMAX]int
    var jumlahKategori int
    var ketemu bool

    if *jumlahBuku == 0 {
        fmt.Println("Belum ada buku dalam koleksi.")
        return
    }

    for i = 0; i < *jumlahBuku; i++ {
        if koleksi[i].status == "Tersedia" {
            totalTersedia++
        }

        ketemu = false
        for j = 0; j < jumlahKategori; j++ {
            if kategoriList[j] == koleksi[i].kategori {
                kategoriCount[j]++
                ketemu = true
                break
            }
        }
        if !ketemu {
            kategoriList[jumlahKategori] = koleksi[i].kategori
            kategoriCount[jumlahKategori] = 1
            jumlahKategori++
        }
    }

    fmt.Println("---------- Statistik PerpusKu ----------")
    fmt.Printf("Total Koleksi Buku : %d\n", *jumlahBuku)
    fmt.Printf("Total Tersedia     : %d\n", totalTersedia)
    fmt.Printf("Total Tidak Tersedia: %d\n", *jumlahBuku - totalTersedia)
    fmt.Println("-------------------------------")
    fmt.Println("Jumlah Buku per Kategori:")
    for i = 0; i < jumlahKategori; i++ {
        fmt.Printf("  %-20s : %d buku\n", kategoriList[i], kategoriCount[i])
    }
    fmt.Println("===============================")
}
