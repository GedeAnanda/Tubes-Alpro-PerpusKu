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

func main() {
	var koleksi tabBuku
	var jumlahBuku int = 0
	var pilih int

	for {
		clearScreen()
		fmt.Println("╔══════════════════════════════════╗")
		fmt.Println("║  SiPerpus - Katalog Digital      ║")
		fmt.Println("╠══════════════════════════════════╣")
		fmt.Println("║  1. Tambah Buku                  ║")
		fmt.Println("║  2. Tampilkan Semua Buku         ║")
		fmt.Println("║  3. Ubah Data Buku               ║")
		fmt.Println("║  4. Hapus Buku                   ║")
		fmt.Println("║  5. Cari Buku - Judul            ║")
		fmt.Println("║  6. Cari Buku - ID               ║")
		fmt.Println("║  7. Urutkan Tahun: 2000 -> 2024  ║")
		fmt.Println("║  8. Urutkan Tahun: 2024 -> 2000  ║")
		fmt.Println("║  9. Statistik                    ║")
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
			selectionSortTahunTerbit(&koleksi, &jumlahBuku)
		case 8:
			insertionSortTahunTerbit(&koleksi, &jumlahBuku)
		case 9:
			statistik(&koleksi, &jumlahBuku)
		case 0:
			clearScreen()
			fmt.Println("╔══════════════════════════════════╗")
			fmt.Println("║    Terima kasih - SiPerpus!      ║")
			fmt.Println("╚══════════════════════════════════╝")
			fmt.Println()
			return
		default:
			fmt.Println()
			fmt.Println("  Pilihan tidak valid!")
			fmt.Println()
		}
	}
}

func tambahBuku(koleksi *tabBuku, jumlahBuku *int) {
	if *jumlahBuku >= NMAX {
		fmt.Println("╔══════════════════════════════════╗")
		fmt.Println("║   Kapasitas Perpustakaan Penuh   ║")
		fmt.Println("╚══════════════════════════════════╝")
		fmt.Print("  Tekan Enter untuk kembali ke menu...")
		bacaInput()
		return
	}
	var buku Buku
	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║        Tambah Buku Baru          ║")
	fmt.Println("╚══════════════════════════════════╝")
	fmt.Print("  ID Buku      : ")
	buku.id = bacaInput()
	fmt.Print("  Judul        : ")
	buku.judul = bacaInput()
	fmt.Print("  Penulis      : ")
	buku.penulis = bacaInput()
	fmt.Print("  Kategori     : ")
	buku.kategori = bacaInput()
	fmt.Print("  Tahun Terbit : ")
	buku.tahunTerbit = bacaInt()
	buku.status = "Tersedia"

	koleksi[*jumlahBuku] = buku
	*jumlahBuku++
	fmt.Println()
	fmt.Println("  Buku berhasil ditambahkan!")
	fmt.Println()
	fmt.Print("  Tekan Enter untuk kembali ke menu...")
	bacaInput()
}

func tampilkanSemuaBuku(koleksi *tabBuku, jumlahBuku *int) {
	var i int
	var buku Buku

	if *jumlahBuku == 0 {
		fmt.Println("╔══════════════════════════════════╗")
		fmt.Println("║       Koleksi Masih Kosong       ║")
		fmt.Println("╚══════════════════════════════════╝")
		fmt.Print("  Tekan Enter untuk kembali ke menu...")
		bacaInput()
		return
	}

	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║       Daftar Koleksi Buku        ║")
	fmt.Println("╚══════════════════════════════════╝")
	fmt.Printf("  Total: %d buku\n\n", *jumlahBuku)

	for i = 0; i < *jumlahBuku; i++ {
		buku = koleksi[i]
		fmt.Println("  +----------------------------------+")
		fmt.Printf("  | No      : %-22d |\n", i+1)
		fmt.Println("  +----------------------------------+")
		fmt.Printf("  | ID      : %-22s |\n", buku.id)
		fmt.Printf("  | Judul   : %-22s |\n", buku.judul)
		fmt.Printf("  | Penulis : %-22s |\n", buku.penulis)
		fmt.Printf("  | Kategori: %-22s |\n", buku.kategori)
		fmt.Printf("  | Tahun   : %-22d |\n", buku.tahunTerbit)
		fmt.Printf("  | Status  : %-22s |\n", buku.status)
		fmt.Println("  +----------------------------------+")
		fmt.Println()
	}
	fmt.Print("  Tekan Enter untuk kembali ke menu...")
	bacaInput()
}

func ubahBuku(koleksi *tabBuku, jumlahBuku *int) {
	var id string
	var idx, i int
	var buku Buku
	idx = -1

	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║         Ubah Data Buku           ║")
	fmt.Println("╚══════════════════════════════════╝")
	fmt.Print("  Masukkan ID Buku: ")
	id = bacaInput()

	for i = 0; i < *jumlahBuku; i++ {
		if koleksi[i].id == id {
			idx = i
			break
		}
	}

	if idx == -1 {
		fmt.Println()
		fmt.Println("  Buku tidak ditemukan!")
		fmt.Println()
		fmt.Print("  Tekan Enter untuk kembali ke menu...")
		bacaInput()
		return
	}

	fmt.Println()
	fmt.Println("  +----------------------------------+")
	fmt.Println("  |          Data Lama               |")
	fmt.Println("  +----------------------------------+")
	fmt.Printf("  | ID      : %-22s |\n", koleksi[idx].id)
	fmt.Printf("  | Judul   : %-22s |\n", koleksi[idx].judul)
	fmt.Printf("  | Penulis : %-22s |\n", koleksi[idx].penulis)
	fmt.Printf("  | Kategori: %-22s |\n", koleksi[idx].kategori)
	fmt.Printf("  | Tahun   : %-22d |\n", koleksi[idx].tahunTerbit)
	fmt.Printf("  | Status  : %-22s |\n", koleksi[idx].status)
	fmt.Println("  +----------------------------------+")
	fmt.Println()
	fmt.Println("  Masukkan Data Baru:")

	buku.id = id
	fmt.Print("  Judul        : ")
	buku.judul = bacaInput()
	fmt.Print("  Penulis      : ")
	buku.penulis = bacaInput()
	fmt.Print("  Kategori     : ")
	buku.kategori = bacaInput()
	fmt.Print("  Tahun Terbit : ")
	buku.tahunTerbit = bacaInt()
	fmt.Print("  Status (Tersedia/Tidak_Tersedia): ")
	buku.status = bacaInput()

	koleksi[idx] = buku
	fmt.Println()
	fmt.Println("  Buku berhasil diubah!")
	fmt.Println()
	fmt.Print("  Tekan Enter untuk kembali ke menu...")
	bacaInput()
}

func hapusBuku(koleksi *tabBuku, jumlahBuku *int) {
	var id string
	var idx, i int
	idx = -1

	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║           Hapus Buku             ║")
	fmt.Println("╚══════════════════════════════════╝")
	fmt.Print("  Masukkan ID Buku: ")
	id = bacaInput()

	for i = 0; i < *jumlahBuku; i++ {
		if koleksi[i].id == id {
			idx = i
			break
		}
	}

	if idx == -1 {
		fmt.Println()
		fmt.Println("  Buku tidak ditemukan!")
		fmt.Println()
		fmt.Print("  Tekan Enter untuk kembali ke menu...")
		bacaInput()
		return
	}

	fmt.Println()
	fmt.Println("  +----------------------------------+")
	fmt.Println("  |      Buku yang akan dihapus      |")
	fmt.Println("  +----------------------------------+")
	fmt.Printf("  | ID      : %-22s |\n", koleksi[idx].id)
	fmt.Printf("  | Judul   : %-22s |\n", koleksi[idx].judul)
	fmt.Printf("  | Penulis : %-22s |\n", koleksi[idx].penulis)
	fmt.Printf("  | Tahun   : %-22d |\n", koleksi[idx].tahunTerbit)
	fmt.Println("  +----------------------------------+")
	fmt.Println()
	fmt.Print("  Konfirmasi hapus? (y/n): ")
	var konfirmasi string
	konfirmasi = bacaInput()

	if konfirmasi != "y" {
		fmt.Println()
		fmt.Println("  Penghapusan dibatalkan!")
		fmt.Println()
		fmt.Print("  Tekan Enter untuk kembali ke menu...")
		bacaInput()
		return
	}

	for i = idx; i < *jumlahBuku-1; i++ {
		koleksi[i] = koleksi[i+1]
	}
	*jumlahBuku--
	fmt.Println()
	fmt.Println("  Buku berhasil dihapus!")
	fmt.Println()
	fmt.Print("  Tekan Enter untuk kembali ke menu...")
	bacaInput()
}

func searchJudul(koleksi *tabBuku, jumlahBuku *int) {
	var i int
	var found bool
	var find string

	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║       Cari Buku - Judul          ║")
	fmt.Println("╚══════════════════════════════════╝")
	fmt.Print("  Masukkan Judul Buku: ")
	find = bacaInput()

	for i = 0; i < *jumlahBuku; i++ {
		if koleksi[i].judul == find {
			if !found {
				fmt.Println()
				fmt.Println("  +----------------------------------+")
				fmt.Println("  |        Buku Ditemukan!           |")
				fmt.Println("  +----------------------------------+")
			}
			fmt.Printf("  | ID      : %-22s |\n", koleksi[i].id)
			fmt.Printf("  | Judul   : %-22s |\n", koleksi[i].judul)
			fmt.Printf("  | Penulis : %-22s |\n", koleksi[i].penulis)
			fmt.Printf("  | Kategori: %-22s |\n", koleksi[i].kategori)
			fmt.Printf("  | Tahun   : %-22d |\n", koleksi[i].tahunTerbit)
			fmt.Printf("  | Status  : %-22s |\n", koleksi[i].status)
			fmt.Println("  +----------------------------------+")
			fmt.Println()
			found = true
		}
	}

	if !found {
		fmt.Println()
		fmt.Println("  Buku tidak ditemukan!")
		fmt.Println()
	}
	fmt.Print("  Tekan Enter untuk kembali ke menu...")
	bacaInput()
}

func searchId(koleksi *tabBuku, jumlahBuku *int) {
	var find string
	var kiri, kanan, mid int

	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║        Cari Buku - ID            ║")
	fmt.Println("╚══════════════════════════════════╝")
	fmt.Print("  Masukkan ID Buku: ")
	find = bacaInput()

	sortById(koleksi, jumlahBuku)

	kiri = 0
	kanan = *jumlahBuku - 1

	for kiri <= kanan {
		mid = (kiri + kanan) / 2
		if koleksi[mid].id == find {
			fmt.Println()
			fmt.Println("  +----------------------------------+")
			fmt.Println("  |        Buku Ditemukan!           |")
			fmt.Println("  +----------------------------------+")
			fmt.Printf("  | ID      : %-22s |\n", koleksi[mid].id)
			fmt.Printf("  | Judul   : %-22s |\n", koleksi[mid].judul)
			fmt.Printf("  | Penulis : %-22s |\n", koleksi[mid].penulis)
			fmt.Printf("  | Kategori: %-22s |\n", koleksi[mid].kategori)
			fmt.Printf("  | Tahun   : %-22d |\n", koleksi[mid].tahunTerbit)
			fmt.Printf("  | Status  : %-22s |\n", koleksi[mid].status)
			fmt.Println("  +----------------------------------+")
			fmt.Println()
			fmt.Print("  Tekan Enter untuk kembali ke menu...")
			bacaInput()
			return
		} else if find < koleksi[mid].id {
			kanan = mid - 1
		} else {
			kiri = mid + 1
		}
	}

	fmt.Println()
	fmt.Println("  Buku tidak ditemukan!")
	fmt.Println()
	fmt.Print("  Tekan Enter untuk kembali ke menu...")
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

func selectionSortTahunTerbit(koleksi *tabBuku, jumlahBuku *int) {
	var i, minIdx, N int
	var temp Buku

	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║  Urutkan Tahun: 2000 -> 2024     ║")
	fmt.Println("╚══════════════════════════════════╝")

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

	fmt.Println()
	fmt.Println("  Data berhasil diurutkan Ascending!")
	fmt.Println()
	tampilkanSemuaBuku(koleksi, jumlahBuku)
}

func insertionSortTahunTerbit(koleksi *tabBuku, jumlahBuku *int) {
	var pass, i int
	var temp Buku

	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║  Urutkan Tahun: 2024 -> 2000     ║")
	fmt.Println("╚══════════════════════════════════╝")

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

	fmt.Println()
	fmt.Println("  Data berhasil diurutkan Descending!")
	fmt.Println()
	tampilkanSemuaBuku(koleksi, jumlahBuku)
}

func statistik(koleksi *tabBuku, jumlahBuku *int) {
	var i, j int
	var totalTersedia int
	var kategoriList [NMAX]string
	var kategoriCount [NMAX]int
	var jumlahKategori int
	var found bool

	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║       Statistik SiPerpus         ║")
	fmt.Println("╚══════════════════════════════════╝")

	if *jumlahBuku == 0 {
		fmt.Println()
		fmt.Println("  Buku tidak ditemukan!")
		fmt.Println()
		fmt.Print("  Tekan Enter untuk kembali ke menu...")
		bacaInput()
		return
	}

	for i = 0; i < *jumlahBuku; i++ {
		if koleksi[i].status == "Tersedia" {
			totalTersedia++
		}
		found = false
		for j = 0; j < jumlahKategori; j++ {
			if kategoriList[j] == koleksi[i].kategori {
				kategoriCount[j]++
				found = true
				break
			}
		}
		if !found {
			kategoriList[jumlahKategori] = koleksi[i].kategori
			kategoriCount[jumlahKategori] = 1
			jumlahKategori++
		}
	}

	fmt.Println()
	fmt.Println("  +----------------------------------+")
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
	fmt.Println()
	fmt.Print("  Tekan Enter untuk kembali ke menu...")
	bacaInput()
}