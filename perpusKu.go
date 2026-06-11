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

var reader = bufio.NewReader(os.Stdin)

func bacaInput() string {
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
	koleksi[2] = Buku{"B003", "Bumi Manusia", "Fiksi", "Pramoedya A.T.", 1980, "Tidak_Tersedia"}
	koleksi[3] = Buku{"B004", "Golang for Beginners", "Teknologi", "John Doe", 2022, "Tersedia"}
	koleksi[4] = Buku{"B005", "Sapiens", "Sejarah", "Yuval Noah Harari", 2011, "Tersedia"}
	koleksi[5] = Buku{"B006", "Atomic Habits", "Bisnis", "James Clear", 2018, "Tersedia"}
	koleksi[6] = Buku{"B007", "Kalkulus Edisi 9", "Pendidikan", "Purcell", 2003, "Tidak_Tersedia"}
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
		fmt.Println("╔══════════════════════════════════════════════╗")
		fmt.Println("║          📚 SiPerpus - Katalog Digital       ║")
		fmt.Println("╠══════════════════════════════════════════════╣")
		fmt.Println("║  1. ➕ Tambah Buku                           ║")
		fmt.Println("║  2. 📖 Tampilkan Semua Buku                  ║")
		fmt.Println("║  3. ✏️  Ubah Data Buku                        ║")
		fmt.Println("║  4. 🗑️  Hapus Buku                            ║")
		fmt.Println("║  5. 🔍 Cari Buku - Judul                     ║")
		fmt.Println("║  6. 🔑 Cari Buku - ID                        ║")
		fmt.Println("║  7. 📈 Urutkan Tahun: Lama -> Baru (Asc)     ║")
		fmt.Println("║  8. 📉 Urutkan Tahun: Baru -> Lama (Desc)    ║")
		fmt.Println("║  9. 📊 Statistik Perpustakaan                ║")
		fmt.Println("║ 10. 📁 Filter Buku per Kategori              ║")
		fmt.Println("║  0. 🚪 Keluar                                ║")
		fmt.Println("╚══════════════════════════════════════════════╝")
		fmt.Print("  👉 Pilih menu (0-10): ")
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
			selectionSortTahunTerbit(&koleksi, &jumlahBuku)
		case 8:
			insertionSortTahunTerbit(&koleksi, &jumlahBuku)
		case 9:
			statistik(&koleksi, &jumlahBuku)
		case 10:
			filterKategori(&koleksi, &jumlahBuku)
		case 0:
			clearScreen()
			fmt.Println("╔══════════════════════════════════════════════╗")
			fmt.Println("║  ✨ Terima kasih telah menggunakan SiPerpus! ║")
			fmt.Println("║           Sampai jumpa kembali 👋            ║")
			fmt.Println("╚══════════════════════════════════════════════╝")
			fmt.Println()
			return
		default:
			fmt.Println("\n  ❌ Pilihan tidak valid! Silakan coba lagi.")
			fmt.Print("  🔄 Tekan Enter untuk mengulang...")
			bacaInput()
		}
	}
}

func tambahBuku(koleksi *tabBuku, jumlahBuku *int) {
	if *jumlahBuku >= NMAX {
		fmt.Println("\n╔══════════════════════════════════════════════╗")
		fmt.Println("║  ⚠️  Kapasitas Perpustakaan Penuh!           ║")
		fmt.Println("╚══════════════════════════════════════════════╝")
		fmt.Print("  🔙 Tekan Enter untuk kembali ke menu...")
		bacaInput()
		return
	}
	var buku Buku
	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║              ➕ Tambah Buku Baru             ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Print("  🆔 ID Buku      : ")
	inputID := bacaInput()

	var idDuplicate bool = false
	for i := 0; i < *jumlahBuku && !idDuplicate; i++ {
		if strings.ToLower(koleksi[i].id) == strings.ToLower(inputID) {
			idDuplicate = true
		}
	}
	
	if idDuplicate {
		fmt.Println("\n  ❌ GAGAL: ID Buku tersebut sudah terdaftar di sistem!")
		fmt.Print("  🔙 Tekan Enter untuk kembali ke menu...")
		bacaInput()
		return
	}

	buku.id = inputID
	fmt.Print("  📕 Judul        : ")
	buku.judul = bacaInput()
	fmt.Print("  ✍️  Penulis      : ")
	buku.penulis = bacaInput()
	fmt.Print("  📁 Kategori     : ")
	buku.kategori = bacaInput()
	fmt.Print("  📅 Tahun Terbit : ")
	buku.tahunTerbit = bacaInt()
	buku.status = "Tersedia"

	koleksi[*jumlahBuku] = buku
	*jumlahBuku++
	fmt.Println("\n  ✅ Buku berhasil ditambahkan ke dalam sistem!")
	fmt.Print("  🔙 Tekan Enter untuk kembali ke menu...")
	bacaInput()
}

func tampilkanSemuaBuku(koleksi *tabBuku, jumlahBuku *int) {
	var i int
	var buku Buku

	if *jumlahBuku == 0 {
		fmt.Println("\n╔══════════════════════════════════════════════╗")
		fmt.Println("║             📭 Koleksi Masih Kosong          ║")
		fmt.Println("╚══════════════════════════════════════════════╝")
		fmt.Print("  🔙 Tekan Enter untuk kembali ke menu...")
		bacaInput()
		return
	}

	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║             📖 Daftar Koleksi Buku           ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Printf("  📌 Total: %d buku terdaftar\n\n", *jumlahBuku)

	for i = 0; i < *jumlahBuku; i++ {
		buku = koleksi[i]
		fmt.Println("  +--------------------------------------------+")
		fmt.Printf("  | 🔢 No      : %-27d |\n", i+1)
		fmt.Println("  +--------------------------------------------+")
		fmt.Printf("  | 🆔 ID      : %-27s |\n", buku.id)
		fmt.Printf("  | 📕 Judul   : %-27s |\n", buku.judul)
		fmt.Printf("  | ✍️  Penulis : %-27s |\n", buku.penulis)
		fmt.Printf("  | 📁 Kategori: %-27s |\n", buku.kategori)
		fmt.Printf("  | 📅 Tahun   : %-27d |\n", buku.tahunTerbit)
		fmt.Printf("  | 📌 Status  : %-27s |\n", buku.status)
		fmt.Println("  +--------------------------------------------+\n")
	}
	fmt.Print("  🔙 Tekan Enter untuk kembali ke menu...")
	bacaInput()
}


func ubahBuku(koleksi *tabBuku, jumlahBuku *int) {
	var id string
	var idx, i int
	var buku Buku
	idx = -1

	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║               ✏️  Ubah Data Buku              ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Print("  🔍 Masukkan ID Buku yang ingin diubah: ")
	id = bacaInput()

	for i = 0; i < *jumlahBuku && idx == -1; i++ {
		if koleksi[i].id == id {
			idx = i
		}
	}

	if idx == -1 {
		fmt.Println("\n  ❌ Buku dengan ID tersebut tidak ditemukan!")
		fmt.Print("  🔙 Tekan Enter untuk kembali ke menu...")
		bacaInput()
		return
	}

	fmt.Println("\n  +--------------------------------------------+")
	fmt.Println("  |              📜 Data Lama                  |")
	fmt.Println("  +--------------------------------------------+")
	fmt.Printf("  | 🆔 ID      : %-27s |\n", koleksi[idx].id)
	fmt.Printf("  | 📕 Judul   : %-27s |\n", koleksi[idx].judul)
	fmt.Printf("  | ✍️  Penulis : %-27s |\n", koleksi[idx].penulis)
	fmt.Printf("  | 📁 Kategori: %-27s |\n", koleksi[idx].kategori)
	fmt.Printf("  | 📅 Tahun   : %-27d |\n", koleksi[idx].tahunTerbit)
	fmt.Printf("  | 📌 Status  : %-27s |\n", koleksi[idx].status)
	fmt.Println("  +--------------------------------------------+\n")
	
	fmt.Println("  ✨ Masukkan Data Baru:")
	buku.id = id
	fmt.Print("  📕 Judul        : ")
	buku.judul = bacaInput()
	fmt.Print("  ✍️  Penulis      : ")
	buku.penulis = bacaInput()
	fmt.Print("  📁 Kategori     : ")
	buku.kategori = bacaInput()
	fmt.Print("  📅 Tahun Terbit : ")
	buku.tahunTerbit = bacaInt()
	fmt.Print("  📌 Status (Tersedia/Tidak_Tersedia): ")
	buku.status = bacaInput()

	koleksi[idx] = buku
	fmt.Println("\n  ✅ Data buku berhasil diperbarui!")
	fmt.Print("  🔙 Tekan Enter untuk kembali ke menu...")
	bacaInput()
}

// Sequential sEARCH
func hapusBuku(koleksi *tabBuku, jumlahBuku *int) {
	var id string
	var idx, i int
	idx = -1

	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║                 🗑️  Hapus Buku                ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Print("  🔍 Masukkan ID Buku yang ingin dihapus: ")
	id = bacaInput()

	for i = 0; i < *jumlahBuku && idx == -1; i++ {
		if koleksi[i].id == id {
			idx = i
		}
	}

	if idx == -1 {
		fmt.Println("\n  ❌ Buku dengan ID tersebut tidak ditemukan!")
		fmt.Print("  🔙 Tekan Enter untuk kembali ke menu...")
		bacaInput()
		return
	}

	fmt.Println("\n  +--------------------------------------------+")
	fmt.Println("  |         ⚠️  Buku yang akan dihapus          |")
	fmt.Println("  +--------------------------------------------+")
	fmt.Printf("  | 🆔 ID      : %-27s |\n", koleksi[idx].id)
	fmt.Printf("  | 📕 Judul   : %-27s |\n", koleksi[idx].judul)
	fmt.Printf("  | ✍️  Penulis : %-27s |\n", koleksi[idx].penulis)
	fmt.Printf("  | 📅 Tahun   : %-27d |\n", koleksi[idx].tahunTerbit)
	fmt.Println("  +--------------------------------------------+\n")
	
	fmt.Print("  ❓ Konfirmasi hapus permanen? (y/n): ")
	var konfirmasi string
	konfirmasi = bacaInput()

	if strings.ToLower(konfirmasi) != "y" {
		fmt.Println("\n  🛡️  Penghapusan dibatalkan!")
		fmt.Print("  🔙 Tekan Enter untuk kembali ke menu...")
		bacaInput()
		return
	}

	for i = idx; i < *jumlahBuku-1; i++ {
		koleksi[i] = koleksi[i+1]
	}
	*jumlahBuku--
	fmt.Println("\n  ✅ Buku berhasil dihapus dari sistem!")
	fmt.Print("  🔙 Tekan Enter untuk kembali ke menu...")
	bacaInput()
}


func searchJudul(koleksi *tabBuku, jumlahBuku *int) {
	var i int
	var headerSudahDiprint bool
	var find string

	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║              🔍 Cari Buku - Judul            ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Print("  ⌨️  Masukkan Judul Buku: ")
	find = bacaInput()

	for i = 0; i < *jumlahBuku; i++ {
		if strings.ToLower(koleksi[i].judul) == strings.ToLower(find) {
			if !headerSudahDiprint {
				fmt.Println("\n  +--------------------------------------------+")
				fmt.Println("  |             🎉 Buku Ditemukan!             |")
				fmt.Println("  +--------------------------------------------+")
			}
			fmt.Printf("  | 🆔 ID      : %-27s |\n", koleksi[i].id)
			fmt.Printf("  | 📕 Judul   : %-27s |\n", koleksi[i].judul)
			fmt.Printf("  | ✍️  Penulis : %-27s |\n", koleksi[i].penulis)
			fmt.Printf("  | 📁 Kategori: %-27s |\n", koleksi[i].kategori)
			fmt.Printf("  | 📅 Tahun   : %-27d |\n", koleksi[i].tahunTerbit)
			fmt.Printf("  | 📌 Status  : %-27s |\n", koleksi[i].status)
			fmt.Println("  +--------------------------------------------+\n")
			headerSudahDiprint = true
		}
	}

	if !headerSudahDiprint {
		fmt.Println("\n  ❌ Buku tidak ditemukan dalam sistem!")
	}
	fmt.Print("\n  🔙 Tekan Enter untuk kembali ke menu...")
	bacaInput()
}

func searchId(koleksi *tabBuku, jumlahBuku *int) {
	var find string
	var kiri, kanan, mid int
	var foundIdx int = -1

	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║               🔑 Cari Buku - ID              ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Print("  ⌨️  Masukkan ID Buku: ")
	find = bacaInput()

	sortById(koleksi, jumlahBuku)

	kiri = 0
	kanan = *jumlahBuku - 1

	for kiri <= kanan && foundIdx == -1 {
		mid = (kiri + kanan) / 2
		if strings.ToLower(koleksi[mid].id) == strings.ToLower(find) {
			foundIdx = mid
		} else if strings.ToLower(find) < strings.ToLower(koleksi[mid].id) {
			kanan = mid - 1
		} else {
			kiri = mid + 1
		}
	}

	if foundIdx != -1 {
		fmt.Println("\n  +--------------------------------------------+")
		fmt.Println("  |             🎉 Buku Ditemukan!             |")
		fmt.Println("  +--------------------------------------------+")
		fmt.Printf("  | 🆔 ID      : %-27s |\n", koleksi[foundIdx].id)
		fmt.Printf("  | 📕 Judul   : %-27s |\n", koleksi[foundIdx].judul)
		fmt.Printf("  | ✍️  Penulis : %-27s |\n", koleksi[foundIdx].penulis)
		fmt.Printf("  | 📁 Kategori: %-27s |\n", koleksi[foundIdx].kategori)
		fmt.Printf("  | 📅 Tahun   : %-27d |\n", koleksi[foundIdx].tahunTerbit)
		fmt.Printf("  | 📌 Status  : %-27s |\n", koleksi[foundIdx].status)
		fmt.Println("  +--------------------------------------------+\n")
	} else {
		fmt.Println("\n  ❌ Buku dengan ID tersebut tidak ditemukan!")
	}
	
	fmt.Print("  🔙 Tekan Enter untuk kembali ke menu...")
	bacaInput()
}

func sortById(koleksi *tabBuku, jumlahBuku *int) {
	var i, minIdx, N int
	var temp Buku

	for N = 0; N < *jumlahBuku-1; N++ {	
		minIdx = N
		for i = N + 1; i < *jumlahBuku; i++ {
			if strings.ToLower(koleksi[i].id) < strings.ToLower(koleksi[minIdx].id) {
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

func selectionSortTahunTerbit(koleksi *tabBuku, jumlahBuku *int) {
	var i, minIdx, N int
	var temp Buku

	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║     📈 Urutkan Tahun: Lama -> Baru (Asc)     ║")
	fmt.Println("╚══════════════════════════════════════════════╝")

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

	fmt.Println("\n  ✅ Data berhasil diurutkan secara Ascending!")
	tampilkanSemuaBuku(koleksi, jumlahBuku)
}

func insertionSortTahunTerbit(koleksi *tabBuku, jumlahBuku *int) {
	var pass, i int
	var temp Buku

	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║     📉 Urutkan Tahun: Baru -> Lama (Desc)    ║")
	fmt.Println("╚══════════════════════════════════════════════╝")

	pass = 1
	for pass < *jumlahBuku {
		i = pass
		temp = koleksi[pass]
		for i > 0 && temp.tahunTerbit > koleksi[i-1].tahunTerbit {
			koleksi[i] = koleksi[i-1]
			i = i - 1
		}
		koleksi[i] = temp
		pass = pass + 1
	}

	fmt.Println("\n  ✅ Data berhasil diurutkan secara Descending!")
	tampilkanSemuaBuku(koleksi, jumlahBuku)
}

func statistik(koleksi *tabBuku, jumlahBuku *int) {
	var i, j int
	var totalTersedia int
	var kategoriList [NMAX]string
	var kategoriCount [NMAX]int
	var jumlahKategori int
	var found bool

	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║             📊 Statistik SiPerpus            ║")
	fmt.Println("╚══════════════════════════════════════════════╝")

	if *jumlahBuku == 0 {
		fmt.Println("\n  📭 Belum ada data buku untuk dihitung!")
		fmt.Print("\n  🔙 Tekan Enter untuk kembali ke menu...")
		bacaInput()
		return
	}

	for i = 0; i < *jumlahBuku; i++ {
		if koleksi[i].status == "Tersedia" {
			totalTersedia++
		}
		found = false
		for j = 0; j < jumlahKategori && !found; j++ {
			if strings.ToLower(kategoriList[j]) == strings.ToLower(koleksi[i].kategori) {
				kategoriCount[j]++
				found = true
			}
		}
		if !found {
			kategoriList[jumlahKategori] = koleksi[i].kategori
			kategoriCount[jumlahKategori] = 1
			jumlahKategori++
		}
	}

	fmt.Println("\n  +--------------------------------------------+")
	fmt.Println("  |             📈 Ringkasan Koleksi           |")
	fmt.Println("  +--------------------------------------------+")
	fmt.Printf("  | 📚 Total Buku        : %-19d |\n", *jumlahBuku)
	fmt.Printf("  | ✅ Tersedia          : %-19d |\n", totalTersedia)
	fmt.Printf("  | ❌ Tidak Tersedia    : %-19d |\n", *jumlahBuku-totalTersedia)
	fmt.Println("  +--------------------------------------------+")
	fmt.Println("  |             📁 Buku per Kategori           |")
	fmt.Println("  +--------------------------------------------+")
	for i = 0; i < jumlahKategori; i++ {
		fmt.Printf("  | 🔹 %-17s : %-19d |\n", kategoriList[i], kategoriCount[i])
	}
	fmt.Println("  +--------------------------------------------+\n")
	fmt.Print("  🔙 Tekan Enter untuk kembali ke menu...")
	bacaInput()
}

func filterKategori(koleksi *tabBuku, jumlahBuku *int) {
	var count int = 0
	var kategoriList [NMAX]string
	var jumlahKategori int
	var found bool
	var pilihan int

	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║            📁 Filter Buku per Kategori       ║")
	fmt.Println("╚══════════════════════════════════════════════╝")

	if *jumlahBuku == 0 {
		fmt.Println("\n  📭 Belum ada data buku!")
		fmt.Print("\n  🔙 Tekan Enter untuk kembali ke menu...")
		bacaInput()
		return
	}

	for i := 0; i < *jumlahBuku; i++ {
		found = false
		for j := 0; j < jumlahKategori && !found; j++ {
			if strings.ToLower(kategoriList[j]) == strings.ToLower(koleksi[i].kategori) {
				found = true
			}
		}
		if !found {
			kategoriList[jumlahKategori] = koleksi[i].kategori
			jumlahKategori++
		}
	}

	fmt.Println("  📌 Pilih Kategori:")
	for i := 0; i < jumlahKategori; i++ {
		fmt.Printf("   %d. %s\n", i+1, kategoriList[i])
	}
	fmt.Println("   0. Batal / Kembali")
	fmt.Println("  ----------------------------------------------")

	fmt.Print("  👉 Pilih nomor kategori: ")
	pilihan = bacaInt()

	if pilihan == 0 {
		return 
	} else if pilihan > 0 && pilihan <= jumlahKategori {
		selectedKategori := kategoriList[pilihan-1]
		fmt.Printf("\n  🔍 Menampilkan hasil untuk kategori: '%s'\n\n", selectedKategori)
		
		for i := 0; i < *jumlahBuku; i++ {
			if strings.ToLower(koleksi[i].kategori) == strings.ToLower(selectedKategori) {
				fmt.Println("  +--------------------------------------------+")
				fmt.Printf("  | 🆔 ID      : %-27s |\n", koleksi[i].id)
				fmt.Printf("  | 📕 Judul   : %-27s |\n", koleksi[i].judul)
				fmt.Printf("  | ✍️  Penulis : %-27s |\n", koleksi[i].penulis)
				fmt.Printf("  | 📅 Tahun   : %-27d |\n", koleksi[i].tahunTerbit)
				fmt.Printf("  | 📌 Status  : %-27s |\n", koleksi[i].status)
				fmt.Println("  +--------------------------------------------+\n")
				count++
			}
		}
		fmt.Printf("  ✅ Total ditemukan: %d buku.\n", count)

	} else {
		fmt.Println("\n  ❌ Pilihan nomor tidak valid!")
	}
	fmt.Print("\n  🔙 Tekan Enter untuk kembali ke menu...")
	bacaInput()
}