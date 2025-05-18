package main

import "fmt"

const maxTeam = 100

type Tim struct {
	Nama            string
	Menang          int
	Kalah           int
	Seri            int
	Poin            int
	SkorMenang      int
	SkorKalah       int
	JumlahGol       int
	JumlahKebobolan int
	JumlahMain      int
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
		daftarTim[idx1].JumlahGol += skor1
		daftarTim[idx1].JumlahKebobolan += skor2

		// Update statistik tim 2
		daftarTim[idx2].JumlahMain++
		daftarTim[idx2].JumlahGol += skor2
		daftarTim[idx2].JumlahKebobolan += skor1

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
