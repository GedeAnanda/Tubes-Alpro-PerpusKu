package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

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

func bacaInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func bacaInt() int {
	text := bacaInput()
	num, _ := strconv.Atoi(text)
	return num
}

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func initDataDummy(koleksi *tabBuku, jumlahBuku *int) {
	koleksi[0] = Buku{"B001", "Laskar Pelangi", "Fiksi", "Andrea Hirata", 2005, "Tersedia"}
	koleksi[1] = Buku{"B002", "Clean Code", "Teknologi", "Robert C. Martin", 2008, "Tersedia"}
	koleksi[2] = Buku{"B003", "Bumi Manusia", "Fiksi", "Pramoedya A.T.", 1980, "Dipinjam"}
	koleksi[3] = Buku{"B004", "Golang for Beginners", "Teknologi", "John Doe", 2022, "Tersedia"}
	koleksi[4] = Buku{"B005", "Sapiens", "Sejarah", "Yuval Noah Harari", 2011, "Tersedia"}
	koleksi[5] = Buku{"B006", "Atomic Habits", "Bisnis", "James Clear", 2018, "Tersedia"}
	koleksi[6] = Buku{"B007", "Kalkulus Edisi 9", "Pendidikan", "Purcell", 2003, "Dipinjam"}
	koleksi[7] = Buku{"B008", "Filosofi Teras", "Filsafat", "Henry Manampiring", 2018, "Tersedia"}
	koleksi[8] = Buku{"B009", "Struktur Data", "Teknologi", "Rinaldi Munir", 2010, "Tersedia"}
	koleksi[9] = Buku{"B010", "Kosmos", "Sains", "Carl Sagan", 1980, "Tersedia"}

	*jumlahBuku = 10
}

func main() {
	var koleksi tabBuku
	var jumlahBuku int = 0
	var pilih int

	initDataDummy(&koleksi, &jumlahBuku)
	
	for {
		clearScreen()
		fmt.Println("╔══════════════════════════════════╗")
		fmt.Println("║  SiPerpus - Katalog Digital      ║")
		fmt.Println("╠══════════════════════════════════╣")
		fmt.Println("║  1. Tambah Buku                  ║")
		fmt.Println("║  2. Tampilkan Semua Buku         ║")
		fmt.Println("║  3. Ubah Data Buku               ║")
		fmt.Println("║  4. Hapus Buku                   ║")
		fmt.Println("║  5. Cari Buku - Judul (Seq)      ║")
		fmt.Println("║  6. Cari Buku - ID    (Bin)      ║")
		fmt.Println("║  7. Sort Tahun Terbit - Asc      ║")
		fmt.Println("║  8. Sort Tahun Terbit - Desc     ║")
		fmt.Println("║  9. Sort Judul Buku   - Asc      ║")
		fmt.Println("║ 10. Sort Judul Buku   - Desc     ║")
		fmt.Println("║ 11. Statistik                    ║")
		fmt.Println("║  0. Keluar                       ║")
		fmt.Println("╚══════════════════════════════════╝")
		fmt.Print("  Pilih menu: ")
		pilih = bacaInt()

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
			selectionSortTahunAsc(&koleksi, &jumlahBuku)
		case 8:
			selectionSortTahunDesc(&koleksi, &jumlahBuku)
		case 9:
			insertionSortJudulAsc(&koleksi, &jumlahBuku)
		case 10:
			insertionSortJudulDesc(&koleksi, &jumlahBuku)
		case 11:
			statistik(&koleksi, &jumlahBuku)
		case 0:
			clearScreen()
			fmt.Println("╔══════════════════════════════════╗")
			fmt.Println("║    Terima kasih - SiPerpus!      ║")
			fmt.Println("╚══════════════════════════════════╝")
			fmt.Println()
			return
		default:
			fmt.Println("\n  Pilihan tidak valid!\n")
			fmt.Print("  Tekan Enter untuk kembali...")
			bacaInput()
		}
	}
}

func tambahBuku(koleksi *tabBuku, jumlahBuku *int) {
	if *jumlahBuku >= NMAX {
		fmt.Println("\n  Kapasitas Perpustakaan Penuh!")
		fmt.Print("  Tekan Enter untuk kembali...")
		bacaInput()
		return
	}
	
	var buku Buku
	var tempID string
	var isDuplicate bool
	var i int
	var pilKategori int

	fmt.Println("\n╔══════════════════════════════════╗")
	fmt.Println("║        Tambah Buku Baru          ║")
	fmt.Println("╚══════════════════════════════════╝")
	
	fmt.Print("  ID Buku      : ")
	tempID = bacaInput()

	isDuplicate = false
	for i = 0; i < *jumlahBuku && !isDuplicate; i++ {
		if koleksi[i].id == tempID {
			isDuplicate = true
		}
	}

	if isDuplicate {
		fmt.Println("\n  GAGAL: ID Buku sudah terdaftar!")
		fmt.Print("  Tekan Enter untuk kembali...")
		bacaInput()
		return
	}

	buku.id = tempID
	fmt.Print("  Judul        : ")
	buku.judul = bacaInput()
	fmt.Print("  Penulis      : ")
	buku.penulis = bacaInput()
	
	fmt.Println("  Pilihan Kategori:")
	fmt.Println("  1. Fiksi")
	fmt.Println("  2. Teknologi")
	fmt.Println("  3. Sejarah")
	fmt.Println("  4. Pendidikan")
	fmt.Println("  5. Bisnis")
	fmt.Println("  6. Sains")
	fmt.Print("  Pilih Kategori (1-6): ")
	pilKategori = bacaInt()
	
	if pilKategori == 1 {
		buku.kategori = "Fiksi"
	} else if pilKategori == 2 {
		buku.kategori = "Teknologi"
	} else if pilKategori == 3 {
		buku.kategori = "Sejarah"
	} else if pilKategori == 4 {
		buku.kategori = "Pendidikan"
	} else if pilKategori == 5 {
		buku.kategori = "Bisnis"
	} else if pilKategori == 6 {
		buku.kategori = "Sains"
	} else {
		buku.kategori = "Lainnya"
	}
	
	fmt.Print("  Tahun Terbit : ")
	buku.tahunTerbit = bacaInt()
	buku.status = "Tersedia"

	koleksi[*jumlahBuku] = buku
	*jumlahBuku++
	fmt.Println("\n  Buku berhasil ditambahkan!")
	fmt.Print("  Tekan Enter untuk kembali...")
	bacaInput()
}

func tampilkanSemuaBuku(koleksi *tabBuku, jumlahBuku *int) {
	var i int
	if *jumlahBuku == 0 {
		fmt.Println("\n  Koleksi Masih Kosong!")
		fmt.Print("  Tekan Enter untuk kembali...")
		bacaInput()
		return
	}

	fmt.Printf("\n  Daftar Koleksi Buku (Total: %d)\n", *jumlahBuku)
	for i = 0; i < *jumlahBuku; i++ {
		fmt.Println("  +----------------------------------+")
		fmt.Printf("  | ID      : %-22s |\n", koleksi[i].id)
		fmt.Printf("  | Judul   : %-22s |\n", koleksi[i].judul)
		fmt.Printf("  | Penulis : %-22s |\n", koleksi[i].penulis)
		fmt.Printf("  | Kategori: %-22s |\n", koleksi[i].kategori)
		fmt.Printf("  | Tahun   : %-22d |\n", koleksi[i].tahunTerbit)
		fmt.Printf("  | Status  : %-22s |\n", koleksi[i].status)
	}
	fmt.Println("  +----------------------------------+")
	fmt.Print("  Tekan Enter untuk kembali...")
	bacaInput()
}

func ubahBuku(koleksi *tabBuku, jumlahBuku *int) {
	var id string
	var idx, i int
	var buku Buku
	var pilKategori int
	idx = -1

	fmt.Print("\n  Masukkan ID Buku yang diubah: ")
	id = bacaInput()

	for i = 0; i < *jumlahBuku && idx == -1; i++ {
		if koleksi[i].id == id {
			idx = i
		}
	}

	if idx == -1 {
		fmt.Println("\n  Buku tidak ditemukan!")
		fmt.Print("  Tekan Enter untuk kembali...")
		bacaInput()
		return
	}

	fmt.Println("\n  Masukkan Data Baru:")
	buku.id = id
	fmt.Print("  Judul        : ")
	buku.judul = bacaInput()
	fmt.Print("  Penulis      : ")
	buku.penulis = bacaInput()
	
	fmt.Println("  Pilihan Kategori:")
	fmt.Println("  1. Fiksi")
	fmt.Println("  2. Teknologi")
	fmt.Println("  3. Sejarah")
	fmt.Println("  4. Pendidikan")
	fmt.Println("  5. Bisnis")
	fmt.Println("  6. Sains")
	fmt.Print("  Pilih Kategori (1-6): ")
	pilKategori = bacaInt()
	
	if pilKategori == 1 {
		buku.kategori = "Fiksi"
	} else if pilKategori == 2 {
		buku.kategori = "Teknologi"
	} else if pilKategori == 3 {
		buku.kategori = "Sejarah"
	} else if pilKategori == 4 {
		buku.kategori = "Pendidikan"
	} else if pilKategori == 5 {
		buku.kategori = "Bisnis"
	} else if pilKategori == 6 {
		buku.kategori = "Sains"
	} else {
		buku.kategori = "Lainnya"
	}
	
	fmt.Print("  Tahun Terbit : ")
	buku.tahunTerbit = bacaInt()
	fmt.Print("  Status (Tersedia/Dipinjam): ")
	buku.status = bacaInput()

	koleksi[idx] = buku
	fmt.Println("\n  Buku berhasil diubah!")
	fmt.Print("  Tekan Enter untuk kembali...")
	bacaInput()
}

func hapusBuku(koleksi *tabBuku, jumlahBuku *int) {
	var id string
	var idx, i int
	idx = -1

	fmt.Print("\n  Masukkan ID Buku yang dihapus: ")
	id = bacaInput()

	for i = 0; i < *jumlahBuku && idx == -1; i++ {
		if koleksi[i].id == id {
			idx = i
		}
	}

	if idx == -1 {
		fmt.Println("\n  Buku tidak ditemukan!")
		fmt.Print("  Tekan Enter untuk kembali...")
		bacaInput()
		return
	}

	fmt.Printf("\n  Buku '%s' akan dihapus.\n", koleksi[idx].judul)
	fmt.Print("  Konfirmasi hapus? (y/n): ")
	if bacaInput() != "y" {
		fmt.Println("  Penghapusan dibatalkan!")
		fmt.Print("  Tekan Enter untuk kembali...")
		bacaInput()
		return
	}

	for i = idx; i < *jumlahBuku-1; i++ {
		koleksi[i] = koleksi[i+1]
	}
	*jumlahBuku--
	fmt.Println("\n  Buku berhasil dihapus!")
	fmt.Print("  Tekan Enter untuk kembali...")
	bacaInput()
}

func searchJudul(koleksi *tabBuku, jumlahBuku *int) {
	var i int
	var found bool
	var find string

	fmt.Print("\n  Masukkan Judul Buku: ")
	find = bacaInput()

	for i = 0; i < *jumlahBuku; i++ {
		if koleksi[i].judul == find {
			if !found {
				fmt.Println("\n  Buku Ditemukan:")
			}
			fmt.Printf("  - %s (ID: %s, Penulis: %s)\n", koleksi[i].judul, koleksi[i].id, koleksi[i].penulis)
			found = true
		}
	}

	if !found {
		fmt.Println("\n  Buku tidak ditemukan!")
	}
	fmt.Print("\n  Tekan Enter untuk kembali...")
	bacaInput()
}

func searchId(koleksi *tabBuku, jumlahBuku *int) {
	var find string
	var kiri, kanan, mid int
	var found bool
	var idxFound int

	fmt.Print("\n  Masukkan ID Buku (Binary Search): ")
	find = bacaInput()

	sortById(koleksi, jumlahBuku)

	kiri = 0
	kanan = *jumlahBuku - 1
	found = false

	for kiri <= kanan && !found {
		mid = (kiri + kanan) / 2
		if koleksi[mid].id == find {
			found = true
			idxFound = mid
		} else if find < koleksi[mid].id {
			kanan = mid - 1
		} else {
			kiri = mid + 1
		}
	}

	if found {
		fmt.Println("\n  Buku Ditemukan:")
		fmt.Printf("  ID      : %s\n", koleksi[idxFound].id)
		fmt.Printf("  Judul   : %s\n", koleksi[idxFound].judul)
		fmt.Printf("  Penulis : %s\n", koleksi[idxFound].penulis)
		fmt.Printf("  Tahun   : %d\n", koleksi[idxFound].tahunTerbit)
	} else {
		fmt.Println("\n  Buku tidak ditemukan!")
	}
	fmt.Print("\n  Tekan Enter untuk kembali...")
	bacaInput()
}

func sortById(koleksi *tabBuku, jumlahBuku *int) {
	var i, minIdx, N int
	var temp Buku

	for N = 0; N < *jumlahBuku-1; N++ {
		minIdx = N
		for i = N + 1; i < *jumlahBuku; i++ {
			if koleksi[i].id < koleksi[minIdx].id {
				minIdx = i
			}
		}
		if minIdx != N {
			temp = koleksi[minIdx]
			koleksi[minIdx] = koleksi[N]
			koleksi[N] = temp
		}
	}
}

func selectionSortTahunAsc(koleksi *tabBuku, jumlahBuku *int) {
	var i, minIdx, N int
	var temp Buku

	for N = 0; N < *jumlahBuku-1; N++ {
		minIdx = N
		for i = N + 1; i < *jumlahBuku; i++ {
			if koleksi[i].tahunTerbit < koleksi[minIdx].tahunTerbit {
				minIdx = i
			}
		}
		if minIdx != N {
			temp = koleksi[minIdx]
			koleksi[minIdx] = koleksi[N]
			koleksi[N] = temp
		}
	}
	fmt.Println("\n  [Selection Sort] Data diurutkan: Tahun Lama -> Baru")
	tampilkanSemuaBuku(koleksi, jumlahBuku)
}

func selectionSortTahunDesc(koleksi *tabBuku, jumlahBuku *int) {
	var i, maxIdx, N int
	var temp Buku

	for N = 0; N < *jumlahBuku-1; N++ {
		maxIdx = N
		for i = N + 1; i < *jumlahBuku; i++ {
			if koleksi[i].tahunTerbit > koleksi[maxIdx].tahunTerbit {
				maxIdx = i
			}
		}
		if maxIdx != N {
			temp = koleksi[maxIdx]
			koleksi[maxIdx] = koleksi[N]
			koleksi[N] = temp
		}
	}
	fmt.Println("\n  [Selection Sort] Data diurutkan: Tahun Baru -> Lama")
	tampilkanSemuaBuku(koleksi, jumlahBuku)
}

func insertionSortJudulAsc(koleksi *tabBuku, jumlahBuku *int) {
	var pass, i int
	var temp Buku

	pass = 1
	for pass < *jumlahBuku {
		i = pass
		temp = koleksi[pass]
		for i > 0 && temp.judul < koleksi[i-1].judul {
			koleksi[i] = koleksi[i-1]
			i = i - 1
		}
		koleksi[i] = temp
		pass = pass + 1
	}
	fmt.Println("\n  [Insertion Sort] Data diurutkan: Judul A -> Z")
	tampilkanSemuaBuku(koleksi, jumlahBuku)
}

func insertionSortJudulDesc(koleksi *tabBuku, jumlahBuku *int) {
	var pass, i int
	var temp Buku

	pass = 1
	for pass < *jumlahBuku {
		i = pass
		temp = koleksi[pass]
		for i > 0 && temp.judul > koleksi[i-1].judul {
			koleksi[i] = koleksi[i-1]
			i = i - 1
		}
		koleksi[i] = temp
		pass = pass + 1
	}
	fmt.Println("\n  [Insertion Sort] Data diurutkan: Judul Z -> A")
	tampilkanSemuaBuku(koleksi, jumlahBuku)
}

func statistik(koleksi *tabBuku, jumlahBuku *int) {
	var i, j int
	var totalTersedia int
	var kategoriList [NMAX]string
	var kategoriCount [NMAX]int
	var jumlahKategori int
	var foundKategori bool

	if *jumlahBuku == 0 {
		fmt.Println("\n  Buku tidak ditemukan!")
		fmt.Print("  Tekan Enter untuk kembali...")
		bacaInput()
		return
	}

	for i = 0; i < *jumlahBuku; i++ {
		if koleksi[i].status == "Tersedia" {
			totalTersedia++
		}
		
		foundKategori = false
		for j = 0; j < jumlahKategori && !foundKategori; j++ {
			if kategoriList[j] == koleksi[i].kategori {
				kategoriCount[j]++
				foundKategori = true
			}
		}
		
		if !foundKategori {
			kategoriList[jumlahKategori] = koleksi[i].kategori
			kategoriCount[jumlahKategori] = 1
			jumlahKategori++
		}
	}

	fmt.Println("\n  +----------------------------------+")
	fmt.Println("  |       Ringkasan Koleksi          |")
	fmt.Println("  +----------------------------------+")
	fmt.Printf("  | Total Buku        : %-12d |\n", *jumlahBuku)
	fmt.Printf("  | Tersedia          : %-12d |\n", totalTersedia)
	fmt.Printf("  | Tidak Tersedia    : %-12d |\n", *jumlahBuku-totalTersedia)
	fmt.Println("  +----------------------------------+")
	fmt.Println("  |      Buku per Kategori           |")
	fmt.Println("  +----------------------------------+")
	for i = 0; i < jumlahKategori; i++ {
		fmt.Printf("  | %-17s : %-12d |\n", kategoriList[i], kategoriCount[i])
	}
	fmt.Println("  +----------------------------------+")
	fmt.Print("\n  Tekan Enter untuk kembali...")
	bacaInput()
}