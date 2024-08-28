// Tubes Alpo
// Anggota Kelompok :
// 1. Frenwin (103052300054)
// 2. Adelina Vivian magdiel (103052300059)
//
// Who one to be Millionare
// LOGIN ADMIN
// Username : admin
// Password : admin
package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const Maxsoal int = 15
const NMAXPemain int = 10

type info struct {
	namaplayer string
	score      int
}
type Soal struct {
	index       int
	nPertanyaan string
	answer      string
	option      [4]string
}

const max int = 4

type arrayinfo [NMAXPemain]info

type question struct {
	nPertanyaan, answer string
	option              [max]string
	index               int
	benar               int
	salah               int
}
type arraykuis [Maxsoal]question

func main() {
	var DataPemain arrayinfo
	var nDataPemain int
	var Soal arraykuis
	var nSoal int
	DummyPemain(&DataPemain, &nDataPemain)
	DummySoal(&Soal, &nSoal)
	welcoming()
	redirect5()
	clearScreen()
	menu()
	login(&DataPemain, &nDataPemain, &Soal, &nSoal)
}
func welcoming() {
	fmt.Println("Selamat datang")
	fmt.Println("Semua pilihan yang ingin diinput berupa angka")
	fmt.Println("Namun ada beberapa bagian yang akan disi menggunakan byte/string")
	fmt.Println("Poin akhir akan diconvert menjadi dollar")
	fmt.Println("Semoga enjoy")

}
func menu() {
	fmt.Println("------------------------------------")
	fmt.Println("      Who one to be Millionare")
	fmt.Println("------------------------------------")
}

// DONE
func login(DataPemain *arrayinfo, nPemain *int, Soal *arraykuis, nSoal *int) {
	var numberlogin int
	fmt.Println("Pilih bagian:")
	fmt.Println("1. Admin")
	fmt.Println("2. Pemain")
	fmt.Println("3. Exit")
	fmt.Print("Pilih(1/2/3): ")
	fmt.Scan(&numberlogin)
	if numberlogin == 1 {
		fmt.Println()
		admin(DataPemain, nPemain, Soal, nSoal)
	} else if numberlogin == 2 {
		fmt.Println()
		player(DataPemain, nPemain, *Soal, *nSoal)
	} else if numberlogin == 3 {
		fmt.Println("Sampai jumpa :)")
	} else {
		fmt.Println()
		fmt.Println("Silahkan pilih sesuai bagian yang diberikan")
		login(DataPemain, nPemain, Soal, nSoal)
	}
	clearScreen()
}

// DONE
func admin(Data *arrayinfo, nData *int, Datasoal *arraykuis, nDataSoal *int) {
	var username string
	var passusername string
	fmt.Print("Silahkan masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&passusername)
	if username == "admin" && passusername == "admin" {
		loading()
		clearScreen()
		admininfo(Data, nData, Datasoal, nDataSoal)
	} else {
		fmt.Println()
		fmt.Println("Maaf, Login sebagai Admin gagal")
		redirect()
		clearScreen()
		login(Data, nData, Datasoal, nDataSoal)
	}
}

func player(Data *arrayinfo, nData *int, Datasoal arraykuis, nDataSoal int) {
	menu()
	fmt.Println("Apakah anda sudah pernah bermain ?")
	fmt.Println("1. Ya")
	fmt.Println("2. Tidak")
	fmt.Println("3. Balik ke login")
	fmt.Print("Pilih(1/2/3): ")
	var pilih int
	var p info
	fmt.Scan(&pilih)
	clearScreen()
	if pilih == 1 {
		Printinfopemain(Data, *nData, Datasoal, nDataSoal)
		fmt.Println("Apakah anda ingin memainkannya lagi?")
		fmt.Println("1. Ya")
		fmt.Println("2. Tidak")
		var pilihmainlagi int
		fmt.Print("Pilih(1/2): ")
		fmt.Scan(&pilihmainlagi)
		if pilihmainlagi == 1 {
			fmt.Print("Silahakan Masukkan nama Pemain:")
			var scannama string
			fmt.Scan(&scannama)
			p.namaplayer = scannama
			mainkanKuis(&p, Data, nData, Datasoal, nDataSoal)
		} else if pilihmainlagi == 2 {
			login(Data, nData, &Datasoal, &nDataSoal)
		} else {
			fmt.Println("Silahkan pilih bagian yang benar")
			player(Data, nData, Datasoal, nDataSoal)
		}
	} else if pilih == 2 {
		fmt.Print("Silahkan Masukkan nama pemain:")
		var pilihmainlagi2 string
		fmt.Scan(&pilihmainlagi2)
		p.namaplayer = pilihmainlagi2
		mainkanKuis(&p, Data, nData, Datasoal, nDataSoal)
	} else if pilih == 3 {
		login(Data, nData, &Datasoal, &nDataSoal)
	} else {
		fmt.Println("Silahkan pilih bagian yang benar")
		player(Data, nData, Datasoal, nDataSoal)
	}
}

func mainkanKuis(p *info, Data *arrayinfo, nData *int, Datasoal arraykuis, nDatasoal int) {
	var DatasoalDummy arraykuis
	var nSoalDummy int
	DummySoal(&DatasoalDummy, &nSoalDummy)
	rand.Seed(time.Now().UnixNano())
	var pilihmainlagi int
	bermainLagi := true
	for bermainLagi {
		fmt.Println("Selamat datang di kuis! Jawab pertanyaan berikut:")
		p.score = 0
		for i := 0; i < nSoalDummy; i++ {
			j := rand.Intn(i + 1)
			DatasoalDummy[i], DatasoalDummy[j] = DatasoalDummy[j], DatasoalDummy[i]
		}
		var cekjawaban bool
		for i := 0; i < 10 && !cekjawaban; i++ {
			soal := DatasoalDummy[i]
			fmt.Println("Pertanyaan:", soal.nPertanyaan)
			for j := 0; j < len(soal.option); j++ {
				fmt.Println(soal.option[j])
			}
			var jawaban string
			fmt.Print("Jawaban Anda: ")
			fmt.Scan(&jawaban)
			loading()
			if jawaban == soal.answer {
				fmt.Println("Benar!")
				Datasoal[i].benar++
				p.score += 100
			} else {
				cekjawaban = true
				fmt.Println("Salah.")
				fmt.Println("Jawaban yang benar adalah")
				fmt.Print("...")
				time.Sleep(2 * time.Second)
				fmt.Print("...")
				time.Sleep(1 * time.Second)
				time.Sleep(2 * time.Second)
				Datasoal[i].salah++
			}
			fmt.Println()
		}
		clearScreen()
		fmt.Printf("Kuis selesai! Skor Anda: %d\n", p.score)
		fmt.Printf("Total skor Anda: %d\n", p.score)
		fmt.Printf("Uang yang anda dapatkan: %d$\n", (p.score * 1000))
		fmt.Print("$")
		fmt.Print("Apakah Anda ingin bermain lagi? (1. Ya, 2. Tidak): ")
		fmt.Scan(&pilihmainlagi)
		clearScreen()
		if pilihmainlagi != 1 {
			bermainLagi = false
			var cek bool = false
			for i := 0; i < *nData && !cek; i++ {
				if p.namaplayer == Data[i].namaplayer {
					if p.score > Data[i].score {
						Data[i].score = p.score
						cek = true
						fmt.Println("Update score berhasil !!")
					} else {
						fmt.Println("Score lebih rendah dari sebelumnnya")
						fmt.Println("Update score gagal !!")
						cek = true
					}
				}
			}
			if !cek {
				fmt.Println("Pemain tidak ditemukan di dalam history")
				fmt.Println("Pemain akan ditambahakan ke daftar pemain baru")
				Pemainbaru(Data, nData, Datasoal, nDatasoal, *p)
			}
		}
	}
	redirect5()
	clearScreen()
	login(Data, nData, &Datasoal, &nDatasoal)
}

// DONE
func Pemainbaru(Data *arrayinfo, nData *int, Datasoal arraykuis, nDataSoal int, Pemainbaru info) {
	if *nData == NMAXPemain {
		fmt.Println("Maaf Pemain sudah penuh")
	} else {
		Data[9] = Pemainbaru
		*nData++
	}
}

// DONE
func Printinfopemain(DataPemain *arrayinfo, nDataPemain int, DataSoal arraykuis, nDatasoal int) {
	var maxscoreidx int
	maxscore(*DataPemain, nDataPemain, maxscoreidx)
	fmt.Println("Player Information:")
	for i := 0; i < *&nDataPemain; i++ {
		fmt.Printf("Player %d: %s - Score: %d\n", i+1, DataPemain[i].namaplayer, DataPemain[i].score)
	}
	fmt.Println("Apakah ingin melihat dari urutan skor??")
	fmt.Println("1. Ya")
	fmt.Println("2. Tidak")
	var pilihurut int
	fmt.Print("Pilih(1/2): ")
	fmt.Scan(&pilihurut)
	if pilihurut == 1 {
		fmt.Println("Pilih Urutan seperti apa?")
		fmt.Println("1. Terbesar ke terkecil")
		fmt.Println("2. Terkecil ke terbesar")
		fmt.Print("Pilih(1/2): ")
		var pilihurutfix int
		fmt.Scan(&pilihurutfix)
		if pilihurutfix == 2 {
			PrintPemainASC(DataPemain, nDataPemain)
			for i := 0; i < nDataPemain; i++ {
				fmt.Printf("Player %d: %s - Score: %d\n", i+1, DataPemain[i].namaplayer, DataPemain[i].score)
			}
		} else if pilihurutfix == 1 {
			PrintPemainDSC(DataPemain, nDataPemain)
			for i := 0; i < nDataPemain; i++ {
				fmt.Printf("Player %d: %s - Score: %d\n", i+1, DataPemain[i].namaplayer, DataPemain[i].score)
			}
		} else {
			fmt.Println("Pilihan anda tidak sesuai")
			player(DataPemain, &nDataPemain, DataSoal, nDatasoal)
		}
	} else if pilihurut == 2 {

	} else {
		fmt.Println("Maaf pilihan anda tidak sesuai")
		player(DataPemain, &nDataPemain, DataSoal, nDatasoal)
	}
}

// DONE Insertion sort
func PrintPemainASC(DataPemain *arrayinfo, nDataPemain int) {
	for i := 1; i < nDataPemain; i++ {
		key := DataPemain[i]
		j := i - 1
		for j >= 0 && DataPemain[j].score > key.score {
			DataPemain[j+1] = DataPemain[j]
			j--
		}
		DataPemain[j+1] = key
	}

}

// DONE Selection sort
func PrintPemainDSC(DataPemain *arrayinfo, nDataPemain int) {
	for i := 0; i < nDataPemain; i++ {
		for j := i + 1; j < nDataPemain; j++ {
			if DataPemain[i].score < DataPemain[j].score {
				DataPemain[i], DataPemain[j] = DataPemain[j], DataPemain[i]
			}
		}
	}

}

// DONE Sequential search
func maxscore(DataPemain arrayinfo, nDataPemain int, maxscore int) {
	fmt.Println("Player dengan highscore tertinggi:")

	var nceknama int
	for i := 0; i < nDataPemain; i++ {
		if DataPemain[i].score > maxscore {
			maxscore = DataPemain[i].score
			nceknama = i
		}
	}
	fmt.Println(DataPemain[nceknama].namaplayer, DataPemain[nceknama].score)
}

// Done
func admininfo(Data *arrayinfo, nData *int, Datasoal *arraykuis, nDataSoal *int) {
	var PilihOpsiAdmin int
	fmt.Println("Selamat datang Admin")
	fmt.Println("Pilih Opsi:")
	fmt.Println("1. Update")
	fmt.Println("2. Tampilkan Soal")
	fmt.Println("3. Balik ke login")
	fmt.Print("Silahkan dipilih: ")
	fmt.Scan(&PilihOpsiAdmin)
	clearScreen()
	if PilihOpsiAdmin == 1 {
		Update(Data, nData, Datasoal, nDataSoal)
	} else if PilihOpsiAdmin == 3 {
		login(Data, nData, Datasoal, nDataSoal)
	} else if PilihOpsiAdmin == 2 {
		PrintSoal(*Datasoal, *nDataSoal)
		printbenar(*Datasoal, *nDataSoal)
		printsalah(*Datasoal, *nDataSoal)
		admininfo(Data, nData, Datasoal, nDataSoal)
	} else {
		fmt.Println()
		fmt.Println("Silahkan pilih sesuai bagian yang diberikan")
		loading()
		clearScreen()
		admininfo(Data, nData, Datasoal, nDataSoal)
	}
}

// DONE
func Update(Data *arrayinfo, nData *int, Datasoal *arraykuis, nDataSoal *int) {
	var PilihUpdate int
	menu()
	fmt.Println("1. Edit Soal")
	fmt.Println("2. Balik ke login")
	fmt.Print("Pilih(1/2): ")
	fmt.Scan(&PilihUpdate)
	clearScreen()
	if PilihUpdate == 1 {
		Editsoal(Datasoal, nDataSoal, *Data, *nData)
	} else if PilihUpdate == 2 {
		clearScreen()
		login(Data, nData, Datasoal, nDataSoal)
	} else {
		clearScreen()
		fmt.Println()
		fmt.Println("Silahkan pilih sesuai bagian yang diberikan")
		loading()
		clearScreen()
		Update(Data, nData, Datasoal, nDataSoal)
	}
}

// DONE
func Editsoal(Datasoal *arraykuis, nDatasoal *int, DataPemain arrayinfo, nDataPemain int) {
	menu()
	fmt.Println("Silahakan Pilih bagian edit")
	fmt.Println("1. Tambah Soal")
	fmt.Println("2. Hapus Soal")
	fmt.Println("3. Ubah Soal")
	fmt.Print("Pilih(1/2/3): ")
	var PilihEdit int
	fmt.Scan(&PilihEdit)
	clearScreen()
	if PilihEdit == 1 {
		tambahsoal(Datasoal, nDatasoal, DataPemain, nDataPemain)
	} else if PilihEdit == 2 {
		hapussoal(Datasoal, nDatasoal, DataPemain, nDataPemain)
	} else if PilihEdit == 3 {
		ubahsoal(Datasoal, *nDatasoal, DataPemain, nDataPemain)
	} else {
		clearScreen()
		fmt.Println()
		fmt.Println("Silahkan pilih sesuai bagian yang diberikan")
		loading()
		clearScreen()
		Update(&DataPemain, &nDataPemain, Datasoal, nDatasoal)
	}
}

// DONE
func tambahsoal(DataSoal *arraykuis, nDatasoal *int, DataPemain arrayinfo, nDataPemain int) {
	menu()
	if *nDatasoal == Maxsoal {
		fmt.Println("Soal Sudah penuh")
		fmt.Println("Mohon untuk menghapus soal terlebih dahulu atau mengedit soal")
		redirect()
		clearScreen()
		Editsoal(DataSoal, nDatasoal, DataPemain, nDataPemain)
	} else {
		var temptambahsoal question
		var option byte
		var optionfix string
		var option1 byte
		var optionfix1 string
		var option2 byte
		var optionfix2 string
		var option3 byte
		var optionfix3 string
		var option4 byte
		var optionfix4 string
		fmt.Println("Silahkan Tambahkan soal")
		fmt.Print("Masukkan Pertanyaan: ")
		bacaoption(option, &optionfix)
		temptambahsoal.nPertanyaan = optionfix
		fmt.Print("Masukkan Opsi Pilihan 1: ")
		bacaoption2(option1, &optionfix1)
		temptambahsoal.option[0] = optionfix1
		fmt.Print("Masukkan Opsi Pilihan 2: ")
		bacaoption2(option2, &optionfix2)
		temptambahsoal.option[1] = optionfix2
		fmt.Print("Masukkan Opsi Pilihan 3: ")
		bacaoption2(option3, &optionfix3)
		temptambahsoal.option[2] = optionfix3
		fmt.Print("Masukkan Opsi Pilihan 4: ")
		bacaoption2(option4, &optionfix4)
		temptambahsoal.option[3] = optionfix4
		fmt.Print("Masukkan Jawaban (Tuliskan saja hanya satu huruf seperti 'A','B','C','D'): ")
		fmt.Scan(&temptambahsoal.answer)
		if temptambahsoal.answer == "A" || temptambahsoal.answer == "B" || temptambahsoal.answer == "C" || temptambahsoal.answer == "D" {
			fmt.Println("Update Berhasil")
			redirect()
			clearScreen()
		} else {
			fmt.Println("Silahakan isi ulang jawaban (Jika isi jawban tidak sesuai dengan persyaratan maka soal akan dihapus)")
			fmt.Print("Masukkan Jawaban (Tuliskan saja hanya satu huruf seperti 'A','B','C','D'): ")
			fmt.Scan(&temptambahsoal.answer)
			if temptambahsoal.answer == "A" || temptambahsoal.answer == "B" || temptambahsoal.answer == "C" || temptambahsoal.answer == "D" {
				fmt.Println("Update Berhasil")
			} else {
				fmt.Println("Tambah Soal Gagal")
				loading()
				clearScreen()
				tambahsoal(DataSoal, nDatasoal, DataPemain, nDataPemain)
			}

		}
		DataSoal[14] = temptambahsoal
		*nDatasoal++
		fmt.Println("Penambahan soal berhasil")
	}
	login(&DataPemain, &nDataPemain, DataSoal, nDatasoal)

}

// DONE
func hapussoal(DataSoal *arraykuis, nDatasoal *int, DataPemain arrayinfo, nDataPemain int) {
	menu()
	PrintSoal(*DataSoal, *nDatasoal)
	fmt.Println("1. Hapus Soal")
	fmt.Println("2. Balik ke login")
	var pilihhapussoal int
	var idx int
	fmt.Print("Pilih(1/2): ")
	fmt.Scan(&pilihhapussoal)
	if pilihhapussoal == 1 {
		var pilihhapussoalfix int
		fmt.Print("Pilih Soal yang ingin dihapus:")
		fmt.Scan(&pilihhapussoalfix)
		if pilihhapussoalfix >= 16 {
			fmt.Println("Silahkan pilih soal sesuai yang ingin dihapus")
			hapussoal(DataSoal, nDatasoal, DataPemain, nDataPemain)
		} else {
			idx = findsoal(*DataSoal, pilihhapussoalfix, *nDatasoal)
			for i := idx; i < *nDatasoal-1; i++ {
				DataSoal[i] = DataSoal[i+1]
			}
			fmt.Println("Hapus Soal telah berhasil")
			*nDatasoal--
			login(&DataPemain, &nDataPemain, DataSoal, nDatasoal)

		}
	} else if pilihhapussoal == 2 {
		login(&DataPemain, &nDataPemain, DataSoal, nDatasoal)
	} else {
		clearScreen()
		fmt.Println()
		fmt.Println("Silahkan pilih sesuai bagian yang diberikan")
		loading()
		clearScreen()
		hapussoal(DataSoal, nDatasoal, DataPemain, nDataPemain)
	}
	loading()
	clearScreen()

}

// Binary Search
func findsoal(Datasoal arraykuis, pilihhapus int, nDatasoal int) int {
	var left int = 0
	var right int = nDatasoal - 1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if Datasoal[mid].index == pilihhapus {
			return pilihhapus
		} else if pilihhapus < Datasoal[mid].index {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
func ubahsoal(DataSoal *arraykuis, nDatasoal int, DataPemain arrayinfo, nDataPemain int) {
	menu()
	PrintSoal(*DataSoal, nDatasoal)
	fmt.Println("1.Soal ingin diubah")
	fmt.Println("2. Balik ke login")
	fmt.Print("Pilih(1/2): ")
	var Pilih int
	fmt.Scan(&Pilih)
	if Pilih == 1 {
		var PilihUbahsoal int
		fmt.Print("Silahkan Pilih Soal yang ingin diubah: ")
		fmt.Scan(&PilihUbahsoal)
		clearScreen()
		if PilihUbahsoal > 15 {
			fmt.Println("Silahkan pilih soal sesuai yang ingin diubah")
			ubahsoal(DataSoal, nDatasoal, DataPemain, nDataPemain)
		} else {
			fmt.Println(DataSoal[PilihUbahsoal-1].nPertanyaan)
			fmt.Println(DataSoal[PilihUbahsoal-1].option[0])
			fmt.Println(DataSoal[PilihUbahsoal-1].option[1])
			fmt.Println(DataSoal[PilihUbahsoal-1].option[2])
			fmt.Println(DataSoal[PilihUbahsoal-1].option[3])
			fmt.Println(DataSoal[PilihUbahsoal-1].answer)
			fmt.Println("Pilih Bagian mana yang ingin diubah")
			fmt.Println("1. Soal/Pertanyaan")
			fmt.Println("2. Pilihan jawaban")
			fmt.Println("3. jawaban")
			fmt.Print("Pilih(1/2/3): ")
			var pilihbagian int
			fmt.Scan(&pilihbagian)
			if pilihbagian == 1 {
				fmt.Println("Masukkan Pertanyaan baru :")
				var pertanyaanbaru byte
				var pertanyaanbarufix string
				bacaoption(pertanyaanbaru, &pertanyaanbarufix)
				DataSoal[PilihUbahsoal-1].nPertanyaan = pertanyaanbarufix
				login(&DataPemain, &nDataPemain, DataSoal, &nDatasoal)
			} else if pilihbagian == 2 {
				fmt.Println("Masukkan Pilihan jawaban:")
				fmt.Print("Masukkan Pilihan pertama:")
				var pilihan1 byte
				var pilihanfix1 string
				bacaoption(pilihan1, &pilihanfix1)
				DataSoal[PilihUbahsoal-1].option[0] = pilihanfix1

				fmt.Print("Masukkan Pilihan kedua:")
				var pilihan2 byte
				var pilihanfix2 string
				bacaoption2(pilihan2, &pilihanfix2)
				DataSoal[PilihUbahsoal-1].option[1] = pilihanfix2
				fmt.Print("Masukkan Pilihan ketiga:")
				var pilihan3 byte
				var pilihanfix3 string
				bacaoption2(pilihan3, &pilihanfix3)
				DataSoal[PilihUbahsoal-1].option[2] = pilihanfix3

				fmt.Print("Masukkan Pilihan keempat:")
				var pilihan4 byte
				var pilihanfix4 string
				bacaoption2(pilihan4, &pilihanfix4)
				DataSoal[PilihUbahsoal-1].option[3] = pilihanfix4
				login(&DataPemain, &nDataPemain, DataSoal, &nDatasoal)
			} else if pilihbagian == 3 {
				fmt.Println("Masukkan Jawaban:")
				var jawaban string
				fmt.Scan(&jawaban)
				if jawaban == "A" || jawaban == "B" || jawaban == "C" || jawaban == "D" {
					fmt.Println("Masukkan Jawaban berhasil")
					DataSoal[PilihUbahsoal-1].answer = jawaban
					login(&DataPemain, &nDataPemain, DataSoal, &nDatasoal)
				} else {
					ubahsoal(DataSoal, nDatasoal, DataPemain, nDataPemain)
				}
			}

		}
	} else if Pilih == 2 {
		login(&DataPemain, &nDataPemain, DataSoal, &nDatasoal)
	} else {
		clearScreen()
		fmt.Println()
		fmt.Println("Silahkan pilih sesuai bagian yang diberikan")
		loading()
		clearScreen()
		ubahsoal(DataSoal, nDatasoal, DataPemain, nDataPemain)
	}

}

// DONE
func PrintSoal(Data arraykuis, nDatasoal int) {
	for i := 0; i < nDatasoal; i++ {
		fmt.Printf("%d %s\n", i+1, Data[i].nPertanyaan)
		fmt.Println(Data[i].option[0])
		fmt.Println(Data[i].option[1])
		fmt.Println(Data[i].option[2])
		fmt.Println(Data[i].option[3])
		fmt.Println(Data[i].answer)
		fmt.Println()
	}
}

// DONE
func bacaoption(pilihan byte, pilihanfix *string) {
	fmt.Scanf("\n%c", &pilihan)
	for pilihan != '\n' {
		*pilihanfix += string(pilihan)
		fmt.Scanf("%c", &pilihan)
	}
}

// DONE
func bacaoption2(pilihan byte, pilihanfix *string) {
	fmt.Scanf("%c", &pilihan)
	for pilihan != '\n' {
		*pilihanfix += string(pilihan)
		fmt.Scanf("%c", &pilihan)
	}
}

// DONE
func DummyPemain(Data *arrayinfo, nDatafunc *int) {
	Data[0].namaplayer = "Frenwin"
	Data[0].score = 1000

	Data[1].namaplayer = "Rasyid"
	Data[1].score = 100

	Data[2].namaplayer = "Dzikran"
	Data[2].score = 200

	Data[3].namaplayer = "Eza"
	Data[3].score = 300

	Data[4].namaplayer = "Adel"
	Data[4].score = 400

	Data[5].namaplayer = "Jeanie"
	Data[5].score = 400

	Data[6].namaplayer = "Felicia"
	Data[6].score = 500

	Data[7].namaplayer = "Carrisa"
	Data[7].score = 600

	Data[8].namaplayer = "Keisha"
	Data[8].score = 700

	*nDatafunc = 9
}
func printbenar(DataSoal arraykuis, nDataSoal int) {
	sortingSoalbenar(&DataSoal, &nDataSoal)
	var i int = 0
	for i != 5 {
		fmt.Println(i+1, DataSoal[i].nPertanyaan)
		fmt.Println("Banyaknya dijawab benar: ", DataSoal[i].benar)
		i++
	}
	fmt.Println()
}
func sortingSoalbenar(DataSoal *arraykuis, nDatasoal *int) {
	for i := 0; i < *nDatasoal; i++ {
		for j := i + 1; j < *nDatasoal; j++ {
			if DataSoal[i].benar < DataSoal[j].benar {
				DataSoal[i].benar, DataSoal[j].benar = DataSoal[j].benar, DataSoal[i].benar
			}
		}
	}

}
func printsalah(DataSoal arraykuis, nDataSoal int) {
	sortingSoalsalah(&DataSoal, &nDataSoal)
	var i int = 0
	for i != 5 {
		fmt.Println(i+1, DataSoal[i].nPertanyaan)
		fmt.Println("Banyaknya dijawab salah: ", DataSoal[i].salah)
		i++
	}
	fmt.Println()
}
func sortingSoalsalah(DataSoal *arraykuis, nDatasoal *int) {
	for i := 0; i < *nDatasoal; i++ {
		for j := i + 1; j < *nDatasoal; j++ {
			if DataSoal[i].salah < DataSoal[j].salah {
				DataSoal[i].salah, DataSoal[j].salah = DataSoal[j].salah, DataSoal[i].salah
			}
		}
	}

}

// DONE
func DummySoal(DatasoalDummy *arraykuis, nSoalDummy *int) {
	DatasoalDummy[0].index = 1
	DatasoalDummy[0].nPertanyaan = "siapa penulis buku “Laut Bercerita?"
	DatasoalDummy[0].answer = "C"
	DatasoalDummy[0].option[0] = "A. Andrea hirata"
	DatasoalDummy[0].option[1] = "B. Tere liye"
	DatasoalDummy[0].option[2] = "C. leila chudori"
	DatasoalDummy[0].option[3] = "D. Pidi baiq"
	DatasoalDummy[0].benar = 5
	DatasoalDummy[0].salah = 1

	DatasoalDummy[1].index = 2
	DatasoalDummy[1].nPertanyaan = "hormon apakah yang dihasilkan ketika seseorang merasa senang?"
	DatasoalDummy[1].answer = "A"
	DatasoalDummy[1].option[0] = "A. Dopamin"
	DatasoalDummy[1].option[1] = "B. Serotonin"
	DatasoalDummy[1].option[2] = "C. Okstitosin"
	DatasoalDummy[1].option[3] = "D. Endorfin"
	DatasoalDummy[1].benar = 7
	DatasoalDummy[1].salah = 1

	DatasoalDummy[2].index = 3
	DatasoalDummy[2].nPertanyaan = "ada berapa musim di Selandia Baru?"
	DatasoalDummy[2].answer = "B"
	DatasoalDummy[2].option[0] = "A. 2 Musim"
	DatasoalDummy[2].option[1] = "B. 3 Musim"
	DatasoalDummy[2].option[2] = "C. 4 Musim"
	DatasoalDummy[2].option[3] = "D. 6 Musim"
	DatasoalDummy[2].benar = 4
	DatasoalDummy[2].salah = 1

	DatasoalDummy[3].index = 4
	DatasoalDummy[3].nPertanyaan = "berapa kira-kira jarak bumi dari matahari?"
	DatasoalDummy[3].answer = "C"
	DatasoalDummy[3].option[0] = "A. 9,3 Juta Mile"
	DatasoalDummy[3].option[1] = "B. 39 Juta Mile"
	DatasoalDummy[3].option[2] = "C. 93 Juta Mile"
	DatasoalDummy[3].option[3] = "D. 193 Juta Mile"
	DatasoalDummy[3].benar = 4
	DatasoalDummy[3].salah = 1

	DatasoalDummy[4].index = 5
	DatasoalDummy[4].nPertanyaan = "tokoh kartun Lotso berasal dari film buatan Walt Disney yaitu..."
	DatasoalDummy[4].answer = "B"
	DatasoalDummy[4].option[0] = "A. Monster, Inc."
	DatasoalDummy[4].option[1] = "B. Toy Story"
	DatasoalDummy[4].option[2] = "C. Alice in Borderland"
	DatasoalDummy[4].option[3] = "D. Lilo & Stitch"
	DatasoalDummy[4].benar = 6
	DatasoalDummy[4].salah = 1

	DatasoalDummy[5].index = 6
	DatasoalDummy[5].nPertanyaan = "serangga apa yang menjadi awal mula istilah “bug” pada komputer?"
	DatasoalDummy[5].answer = "B"
	DatasoalDummy[5].option[0] = "A. Kecoa"
	DatasoalDummy[5].option[1] = "B. Ngengat"
	DatasoalDummy[5].option[2] = "C. Lalat"
	DatasoalDummy[5].option[3] = "D. Kumbang"
	DatasoalDummy[5].benar = 2
	DatasoalDummy[5].salah = 1

	DatasoalDummy[6].index = 7
	DatasoalDummy[6].nPertanyaan = "Mata uang manakah yang pertama kali menggunakan motto “In God We Trust”?"
	DatasoalDummy[6].answer = "D"
	DatasoalDummy[6].option[0] = "A. Nikel"
	DatasoalDummy[6].option[1] = "B. Five dollar bill"
	DatasoalDummy[6].option[2] = "C. One dollar bill"
	DatasoalDummy[6].option[3] = "D. Two-cent piece"
	DatasoalDummy[6].benar = 3
	DatasoalDummy[6].salah = 1

	DatasoalDummy[7].index = 8
	DatasoalDummy[7].nPertanyaan = "Berapa derajat celcius yang setara dengan 77 derajat Fahrenheit?"
	DatasoalDummy[7].answer = "C"
	DatasoalDummy[7].option[0] = "A. 15 Celcius"
	DatasoalDummy[7].option[1] = "B. 20 Celcius"
	DatasoalDummy[7].option[2] = "C. 25 Celcius"
	DatasoalDummy[7].option[3] = "D. 30 Celcius"
	DatasoalDummy[7].benar = 10
	DatasoalDummy[7].salah = 1

	DatasoalDummy[8].index = 9
	DatasoalDummy[8].nPertanyaan = "Manakah di antara nama berikut yang tidak ada dalam judul drama Shakespeare?"
	DatasoalDummy[8].answer = "D"
	DatasoalDummy[8].option[0] = "A. Hamlet"
	DatasoalDummy[8].option[1] = "B. Romeo"
	DatasoalDummy[8].option[2] = "C. Macbeth"
	DatasoalDummy[8].option[3] = "D. Darren"
	DatasoalDummy[8].benar = 1
	DatasoalDummy[8].salah = 1

	DatasoalDummy[9].index = 10
	DatasoalDummy[9].nPertanyaan = "Ikon AS “Paman Sam” selama perang tahun 1812 bekerja sebagai?"
	DatasoalDummy[9].answer = "A"
	DatasoalDummy[9].option[0] = "A. Pengawas daging"
	DatasoalDummy[9].option[1] = "B. Pengantar surat"
	DatasoalDummy[9].option[2] = "C. Sejarawan"
	DatasoalDummy[9].option[3] = "D. Mekanik senjata"
	DatasoalDummy[9].benar = 1
	DatasoalDummy[9].salah = 1

	DatasoalDummy[10].index = 11
	DatasoalDummy[10].nPertanyaan = "Dalam seri buku cerita anak-anak, Paddington Bear berasal dari?"
	DatasoalDummy[10].answer = "B"
	DatasoalDummy[10].option[0] = "A. India"
	DatasoalDummy[10].option[1] = "B. Peru"
	DatasoalDummy[10].option[2] = "C. Canada"
	DatasoalDummy[10].option[3] = "D. Iceland"
	DatasoalDummy[10].benar = 1
	DatasoalDummy[10].salah = 1

	DatasoalDummy[11].index = 12
	DatasoalDummy[11].nPertanyaan = "jenis bunga apakah “American Beauty”?"
	DatasoalDummy[11].answer = "B"
	DatasoalDummy[11].option[0] = "A. Melati"
	DatasoalDummy[11].option[1] = "B. Mawar"
	DatasoalDummy[11].option[2] = "C. Baby Breathe"
	DatasoalDummy[11].option[3] = "D. Sunflower"
	DatasoalDummy[11].benar = 1
	DatasoalDummy[11].salah = 1

	DatasoalDummy[12].index = 13
	DatasoalDummy[12].nPertanyaan = "Emmental adalah keju yang berasal dari negara?"
	DatasoalDummy[12].answer = "B"
	DatasoalDummy[12].option[0] = "A. Perancis"
	DatasoalDummy[12].option[1] = "B. Swiss"
	DatasoalDummy[12].option[2] = "C. Italia"
	DatasoalDummy[12].option[3] = "D. Belanda"
	DatasoalDummy[12].benar = 1
	DatasoalDummy[12].salah = 1

	DatasoalDummy[13].index = 14
	DatasoalDummy[13].nPertanyaan = "Apa itu Butterscotch?"
	DatasoalDummy[13].answer = "B"
	DatasoalDummy[13].option[0] = "A. Shortbeard"
	DatasoalDummy[13].option[1] = "B. Brittle toffee"
	DatasoalDummy[13].option[2] = "C. Garden Flower"
	DatasoalDummy[13].option[3] = "D. Pavement Game"
	DatasoalDummy[13].benar = 1
	DatasoalDummy[13].salah = 1

	DatasoalDummy[14].index = 15
	DatasoalDummy[14].nPertanyaan = "Lukisan terkenal apakah yang pernah dipotong untuk memperbesar doorway?"
	DatasoalDummy[14].answer = "C"
	DatasoalDummy[14].option[0] = "A. Mona Lisa"
	DatasoalDummy[14].option[1] = "B. The Starry Night"
	DatasoalDummy[14].option[2] = "C. The Last Supper"
	DatasoalDummy[14].option[3] = "D. The Birth of Venus"
	DatasoalDummy[14].benar = 1
	DatasoalDummy[14].salah = 1

	*nSoalDummy = 15
}

// Menghapus terminal jika masuk ke dalam suatu kondisi
func clearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// Menghapus hanya beberapa line yang kita mau
func ClearLines(lines int) {
	for i := 0; i < lines; i++ {
		fmt.Print("\033[A\033[K") // Move cursor up one line and clear the line
	}
}

// Fungsi tambahan untuk menghitung mundur saat masuk kedalam suatu kondisi
func redirect() {
	fmt.Println("")
	for i := 5; i > 0; i-- {
		fmt.Printf("\rRedirecting(%d)", i)
		time.Sleep(1 * time.Second)
	}
}
func redirect5() {
	fmt.Println("")
	for i := 7; i > 0; i-- {
		fmt.Printf("\rRedirecting(%d)", i)
		time.Sleep(1 * time.Second)
	}
}
func loading() {
	fmt.Println()
	fmt.Print("Loading")
	time.Sleep(1 * time.Second)
	fmt.Print("...")
	time.Sleep(1 * time.Second)
	fmt.Print("...")
}
