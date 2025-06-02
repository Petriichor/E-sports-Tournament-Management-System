package main

import "fmt"

const maxTeam = 100

type Tim struct {
	Nama       string
	Menang     int
	Kalah      int
	Seri       int
	Poin       int
	SkorMenang int
	SkorKalah  int
	MatchPoint int
	JumlahMain int
}

var daftarTim [maxTeam]Tim
var jumlahTim int

func main() {
	var pilihan int
	for {
		menu()
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahTim()
		case 2:
			ubahNama()
		case 3:
			hapusTim()
		case 4:
			catatHasil()
		case 5:
			tampilKlasemen()
		case 6:
			cariTimSequential()
		case 7:
			cariTimBinary()
		case 8:
			urutkanTimSelection()
		case 9:
			urutkanTimInsertion()
		case 10:
			fmt.Println("Terima kasih telah menggunakan aplikasi!")
			return
		default:
			fmt.Print("Pilihan tidak valid.")
		}
	}
}

func menu() {
	fmt.Println("+===============================+")
	fmt.Println("|      APLIKASI E-SPORTS DATA   |")
	fmt.Println("+===============================+")
	fmt.Println("| 1. Tambah Tim                 |")
	fmt.Println("| 2. Ubah Nama Tim              |")
	fmt.Println("| 3. Hapus Tim                  |")
	fmt.Println("| 4. Catat Hasil Pertandingan   |")
	fmt.Println("| 5. Tampilkan Klasemen         |")
	fmt.Println("| 6. Cari Tim (Sequential)      |")
	fmt.Println("| 7. Cari Tim (Binary)          |")
	fmt.Println("| 8. Urutkan Tim (Selection)    |")
	fmt.Println("| 9. Urutkan Tim (Insertion)    |")
	fmt.Println("| 10. Keluar                    |")
	fmt.Println("+===============================+")
	fmt.Print("Pilih menu [1-10]: ")
}

func tambahTim() {
	var tim Tim
	fmt.Print("Masukkan nama tim: ")
	fmt.Scan(&tim.Nama)

	if jumlahTim >= maxTeam {
		fmt.Println("Maksimal tim sudah tercapai.")
	} else {
		daftarTim[jumlahTim] = tim
		fmt.Println("Tim berhasil ditambahkan!")
		jumlahTim++
	}
}

func ubahNama() {
	var namaAwal, namaAkhir string
	fmt.Print("Masukkan nama tim yang ingin diubah: ")
	fmt.Scan(&namaAwal)
	fmt.Print("Masukkan nama tim yang baru: ")
	fmt.Scan(&namaAkhir)

	var found bool = false
	var idx int

	// Mencari tim yang akan diubah namanya
	for i := 0; i < jumlahTim; i++ {
		if daftarTim[i].Nama == namaAwal {
			found = true
			idx = i
			break
		}
	}

	// Mengganti nama awal menjadi nama akhir
	if found {
		daftarTim[idx].Nama = namaAkhir
		fmt.Println("Nama tim berhasil diubah!")
	} else {
		fmt.Println("Tim tidak ditemukan.")
	}

}

func hapusTim() {
	var namaTim string
	fmt.Print("Masukkan nama tim yang ingin dihapus: ")
	fmt.Scan(&namaTim)

	var found bool = false
	var idx int

	// Mencari tim yang akan dihapus
	for i := 0; i < jumlahTim; i++ {
		if daftarTim[i].Nama == namaTim {
			found = true
			idx = i
			break
		}
	}

	if found {
		// Menggeser data ke kiri untuk menghapus tim
		for i := idx; i < jumlahTim-1; i++ {
			daftarTim[i] = daftarTim[i+1]
		}
		jumlahTim--
		fmt.Println("Tim berhasil dihapus!")
	} else {
		fmt.Println("Tim tidak ditemukan.")
	}
}

func catatHasil() {
	var tim1, tim2 string
	var skor1, skor2 int

	fmt.Print("Masukkan nama tim 1: ")
	fmt.Scan(&tim1)
	fmt.Print("Masukkan skor tim 1: ")
	fmt.Scan(&skor1)
	fmt.Print("Masukkan nama tim 2: ")
	fmt.Scan(&tim2)
	fmt.Print("Masukkan skor tim 2: ")
	fmt.Scan(&skor2)

	var found1, found2 bool = false, false
	var idx1, idx2 int

	// Mencari tim 1 dan tim 2
	for i := 0; i < jumlahTim; i++ {
		if daftarTim[i].Nama == tim1 {
			found1 = true
			idx1 = i
		}
		if daftarTim[i].Nama == tim2 {
			found2 = true
			idx2 = i
		}
	}

	if found1 && found2 {
		// Update statistik tim 1
		daftarTim[idx1].JumlahMain++
		daftarTim[idx1].MatchPoint += skor1

		// Update statistik tim 2
		daftarTim[idx2].JumlahMain++
		daftarTim[idx2].MatchPoint += skor2

		// Mencatat skor apabila tim 1 menang
		if skor1 > skor2 {
			daftarTim[idx1].Menang++
			daftarTim[idx2].Kalah++
			daftarTim[idx1].Poin += 3
			// Mencatat skor apabila tim 2 menang
		} else if skor2 > skor1 {
			daftarTim[idx2].Menang++
			daftarTim[idx1].Kalah++
			daftarTim[idx2].Poin += 3

			// Mencata skor draw
		} else {
			daftarTim[idx1].Seri++
			daftarTim[idx2].Seri++
			daftarTim[idx1].Poin++
			daftarTim[idx2].Poin++
		}

		fmt.Println("Hasil pertandingan berhasil dicatat!")
	} else {
		fmt.Println("Salah satu atau kedua tim tidak ditemukan.")
	}
}

// Fungsi untuk mengurutkan klasemen berdasarkan Poin (descending), jika poin sama urutkan berdasarkan MatchPoint (descending)
func urutkanKlasemen() {
	for i := 0; i < jumlahTim-1; i++ {
		maxIdx := i
		for j := i + 1; j < jumlahTim; j++ {
			if daftarTim[j].Poin > daftarTim[maxIdx].Poin ||
				(daftarTim[j].Poin == daftarTim[maxIdx].Poin && daftarTim[j].MatchPoint > daftarTim[maxIdx].MatchPoint) {
				maxIdx = j
			}
		}
		if maxIdx != i {
			daftarTim[i], daftarTim[maxIdx] = daftarTim[maxIdx], daftarTim[i]
		}
	}
}

func tampilKlasemen() {
	// Urutkan klasemen berdasarkan Poin (descending), jika poin sama urutkan berdasarkan MatchPoint (descending)
	urutkanKlasemen()
	fmt.Println("Klasemen Sementara (Terurut):")
	fmt.Printf("%-3s %-20s %-5s %-5s %-5s %-5s %-10s %-5s\n", "No", "Nama Tim", "Main", "Menang", "Seri", "Kalah", "MatchPoint", "Poin")
	for i := 0; i < jumlahTim; i++ {
		fmt.Printf("%-3d %-20s %-5d %-5d %-5d %-5d %-10d %-5d\n",
			i+1,
			daftarTim[i].Nama,
			daftarTim[i].JumlahMain,
			daftarTim[i].Menang,
			daftarTim[i].Seri,
			daftarTim[i].Kalah,
			daftarTim[i].MatchPoint,
			daftarTim[i].Poin)
	}
}

// Fungsi pencarian tim secara sequential
func cariTimSequential() {
	var namaCari string
	fmt.Print("Masukkan nama tim yang ingin dicari: ")
	fmt.Scan(&namaCari)
	var found bool = false
	for i := 0; i < jumlahTim; i++ {
		if daftarTim[i].Nama == namaCari {
			fmt.Println("Tim ditemukan pada indeks:", i)
			fmt.Printf("Nama: %s, Main: %d, Menang: %d, Seri: %d, Kalah: %d, MatchPoint: %d, Poin: %d\n",
				daftarTim[i].Nama, daftarTim[i].JumlahMain, daftarTim[i].Menang, daftarTim[i].Seri, daftarTim[i].Kalah, daftarTim[i].MatchPoint, daftarTim[i].Poin)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Tim tidak ditemukan.")
	}
}

// Fungsi pencarian tim secara binary search (pastikan sudah terurut)
func cariTimBinary() {
	var namaCari string
	fmt.Print("Masukkan nama tim yang ingin dicari: ")
	fmt.Scan(&namaCari)
	// Pastikan data sudah terurut berdasarkan nama
	urutkanTimSelection() // mengurutkan data berdasarkan nama tim apabila data teracak
	low := 0
	high := jumlahTim - 1
	found := false
	for low <= high {
		mid := (low + high) / 2
		if daftarTim[mid].Nama == namaCari {
			fmt.Println("Tim ditemukan pada indeks:", mid)
			fmt.Printf("Nama: %s, Main: %d, Menang: %d, Seri: %d, Kalah: %d, MatchPoint: %d, Poin: %d\n",
				daftarTim[mid].Nama, daftarTim[mid].JumlahMain, daftarTim[mid].Menang, daftarTim[mid].Seri, daftarTim[mid].Kalah, daftarTim[mid].MatchPoint, daftarTim[mid].Poin)
			found = true
			break
		} else if daftarTim[mid].Nama < namaCari {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if !found {
		fmt.Println("Tim tidak ditemukan.")
	}
}

// Fungsi pengurutan tim dengan selection sort berdasarkan nama tim (ascending)
func urutkanTimSelection() {
	for i := 0; i < jumlahTim-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlahTim; j++ {
			if daftarTim[j].Nama < daftarTim[minIdx].Nama {
				minIdx = j
			}
		}
		if minIdx != i {
			daftarTim[i], daftarTim[minIdx] = daftarTim[minIdx], daftarTim[i]
		}
	}
	fmt.Println("Tim berhasil diurutkan (Selection Sort, berdasarkan nama tim).")
}

// Fungsi pengurutan tim dengan insertion sort berdasarkan nama tim (ascending)
func urutkanTimInsertion() {
	for i := 1; i < jumlahTim; i++ {
		temp := daftarTim[i]
		j := i - 1
		for j >= 0 && daftarTim[j].Nama > temp.Nama {
			daftarTim[j+1] = daftarTim[j]
			j--
		}
		daftarTim[j+1] = temp
	}
	fmt.Println("Tim berhasil diurutkan (Insertion Sort, berdasarkan nama tim).")
}
