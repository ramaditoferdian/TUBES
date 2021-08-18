package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"

	"os/exec"
    "runtime"
    //"time"
)

/* UPDATE : 1. Dapat menampilkan Program Studi beserta jumlah pemilihnya [✓]
//			2. Menambahkan menu 5  [✓]
//			3. Program Sudah Dapat edit data mahasiswa  [✓]
//			4. Program Sudah Dapat edit data Program Studi [✓]
//			5. Error Handling untuk Seluruh Menu 1 - 5 [✓]
			6. Error Handling untuk menu dalem 1 - 3 [✓]

// BUG		: 1.Pada menu ke 2 saat menginputkan edit mahasiswa  [✓]
//		 	  2.Prodi yang sudah dipilih oleh mahasiswa tidak ikut berubah ketika Edit Prodi atau Delete Prodi 	[✓]
//			  3.Belum bisa menemukan Prodi Non Teknik [✓]
			  4.NILAI MASIH BISA DIINPUT STRING error handling [✓]
			  5.NILAI MASIH BISA IINPUT NEGATIF error handling [✓]
			  6.Delete Prodi Error Handling [✓]
			  7.Fungsi Edit masih Bisa Duplicate Nama Prodi [✓]

// Target	: 1. Buat Fungsi Delete untuk Data Mahasiswa [✓]
			  2. Program dapat mengurutkan berdasarkan Nilai TPA atau TBI , Nama Mahasiswa , Program Studi [✓]
			  3. Menambahkan Syarat Lulus Mahasiswa [✓]

// ================================================================== TIPE BENTUKAN & VARIABEL GLOBAL ==========================================================//
// ============================================================================================================================================================//
*/
type ProgramStudiTeknik struct{
	nama, akreditasi string
	
}

type ProgramStudiNonTeknik struct{
	nama, akreditasi string
	
}

type CalonMaba struct{
	nama, gender string
	nilaiTPA , nilaiTBI float64
	nilaiMat, nilaiIPA, nilaiIPS, nilaiInggris float64
	rataRaport , rataTesMasuk , rataSyarat float64
	numprodi int
	prodiTektek [3]ProgramStudiTeknik
	prodiNonTektek [3]ProgramStudiNonTeknik
}

type ProdiTeknik [100]ProgramStudiTeknik
type ProdiNonTeknik [100]ProgramStudiNonTeknik
type arrCalonMaba [100]CalonMaba

var (
	data arrCalonMaba
	prodiTek ProdiTeknik
	prodiNonTek ProdiNonTeknik
	numCamaba int
	numTek int
	numSos int
	numMabaTek int 
	numMabaSos int
	SyaratDiTerima float64
	
)

// UNTUK INPUTAN MENGGUNAKAN SPASI
var inputan = bufio.NewReader(os.Stdin)

var clear map[string]func() //create a map for storing clear funcs

// ===============================================================  MAIN PROGRAM ==========================================================//
// ==========================================================================================================================================//

func main(){
	var pil, jenis, namaProdi, akreditasiProdi, back , key ,konfirmasi string
	var indexCari int

	
	numCamaba = 0

	konfirmasi = "N"
	for konfirmasi != "Y" && konfirmasi != "y" {
		fmt.Println("|================================|")
		fmt.Println("| APLIKASI PENDAFTARAN MAHASISWA |")
		fmt.Println("|================================|")
		menuAwal()
		fmt.Print("Pilihan : ")
		fmt.Scan(&pil)
		CallClear() //Clear Screen
		
		if konfirmasi == "N" || konfirmasi == "n"{
			for pil != "0"{
				switch pil {
				// case 1 = input data
				case "1":
					menu1()
					fmt.Print("Pilihan : ")
					fmt.Scan(&pil)
					CallClear() //Clear Screen
					for pil != "0" {
						if pil == "1"{
							fmt.Print("Pilih [1] Program Studi Teknik atau [2] Program Studi Non Teknik : ")
							fmt.Scan(&jenis)
							if jenis == "1" {
								gakkepake := "m"
								fmt.Scanln(&gakkepake)
								
		
								InputDataProgramStudi(&namaProdi, &akreditasiProdi)
								indexCari = SearchNamaProdi(namaProdi)
		
								if indexCari == -1 {
		
									prodiTek[numTek].nama = namaProdi
									prodiTek[numTek].akreditasi = akreditasiProdi
									numTek++
		
									fmt.Println("|======================|")
									fmt.Println("| BERHASIL DITAMBAHKAN |")
									fmt.Println("|======================|")
		
									fmt.Scanln(&gakkepake)
									fmt.Print("Press 'Enter' to continue...")
									bufio.NewReader(os.Stdin).ReadBytes('\n') 
									CallClear() //Clear Screen 
								}else {
									fmt.Println()
									fmt.Println("GAGAL MENAMBAHKAN PROGRAM STUDI")
									fmt.Println()
									fmt.Println("PROGRAM STUDI YANG ANDA INPUTKAN TELAH TERDAFTAR")
									fmt.Scanln(&gakkepake)
									fmt.Print("Press 'Enter' to continue...")
									bufio.NewReader(os.Stdin).ReadBytes('\n') 
									CallClear() //Clear Screen 
								}
								
		
							}else if jenis == "2"{
								gakkepake := "m"
								fmt.Scanln(&gakkepake)
			
								InputDataProgramStudi(&namaProdi, &akreditasiProdi)
								indexCari = SearchNamaProdi(namaProdi)
		
								if indexCari == -1 {
									prodiNonTek[numSos].nama = namaProdi
									prodiNonTek[numSos].akreditasi = akreditasiProdi
									numSos++
		
									fmt.Println("|======================|")
									fmt.Println("| BERHASIL DITAMBAHKAN |")
									fmt.Println("|======================|")
		
									fmt.Scanln(&gakkepake)
									fmt.Print("Press 'Enter' to continue...")
									bufio.NewReader(os.Stdin).ReadBytes('\n') 
									CallClear() //Clear Screen 
								}else {
									fmt.Println()
									fmt.Println("GAGAL MENAMBAHKAN PROGRAM STUDI")
									fmt.Println()
									fmt.Println("PROGRAM STUDI YANG ANDA INPUTKAN TELAH TERDAFTAR")
									fmt.Scanln(&gakkepake)
									fmt.Print("Press 'Enter' to continue...")
									bufio.NewReader(os.Stdin).ReadBytes('\n') 
									CallClear() //Clear Screen 
								}
		
							}else {
								fmt.Println()
								fmt.Println("===== | INVALID | =====")
								fmt.Println()
								gakkepake := "m"
								fmt.Scanln(&gakkepake)
								fmt.Print("Press 'Enter' to continue...")
								bufio.NewReader(os.Stdin).ReadBytes('\n') 
								CallClear() //Clear Screen 
							}		
						
							fmt.Println()
							CallClear() //Clear Screen 
							
						}else if pil == "2"{
							fmt.Println("Pilih [1] Teknik atau [2] Non-Teknik : ")
							fmt.Print("Pilih : ")
							fmt.Scan(&jenis)
							if jenis == "1"{
								
								if numTek < 3 {
									fmt.Println()
									fmt.Println("Mohon Isi Program Studi Terlebih Dahulu Matur Nuwun")
									fmt.Println()
									for numTek < 3{
										gakkepake := "m"
										fmt.Scanln(&gakkepake)
		
										InputDataProgramStudi(&namaProdi, &akreditasiProdi)
										indexCari = SearchNamaProdi(namaProdi)
				
										if indexCari == -1 {
				
											prodiTek[numTek].nama = namaProdi
											prodiTek[numTek].akreditasi = akreditasiProdi
											numTek++
		
											fmt.Println("|======================|")
											fmt.Println("| BERHASIL DITAMBAHKAN |")
											fmt.Println("|======================|")
				
										}else {
											fmt.Println("GAGAL MENAMBAHKAN PROGRAM STUDI")
											fmt.Println()
											fmt.Println("PROGRAM STUDI YANG ANDA INPUTKAN TELAH TERDAFTAR")
											
										}
									}
		 
								}
								fmt.Print("Press 'Enter' to continue...")
								bufio.NewReader(os.Stdin).ReadBytes('\n') 
								CallClear() //Clear Screen 
		
								fmt.Println("|======================|")
								fmt.Println("| INPUT DATA MAHASISWA |")
								fmt.Println("|======================|")
								
			
								InputDataMahasiswa(numCamaba, jenis)
								InputPilihanProdi(numCamaba, jenis, numTek)
								numCamaba++
							}else if jenis == "2"{
								 if numSos < 3 {
									fmt.Println()
									 fmt.Println("Mohon Isi Program Studi Terlebih Dahulu Matur Nuwun")
									 fmt.Println() 
									 for numSos < 3{
										gakkepake := "m"
										fmt.Scanln(&gakkepake)
			
										 InputDataProgramStudi(&namaProdi, &akreditasiProdi)
										 indexCari = SearchNamaProdi(namaProdi)
		
										 if indexCari == -1 {
											 prodiNonTek[numSos].nama = namaProdi
											 prodiNonTek[numSos].akreditasi = akreditasiProdi
											 numSos++
				 
											 fmt.Println("|======================|")
											 fmt.Println("| BERHASIL DITAMBAHKAN |")
											 fmt.Println("|======================|")
		
										 }else {
											 fmt.Println()
											 fmt.Println("GAGAL MENAMBAHKAN PROGRAM STUDI")
											 fmt.Println()
											 fmt.Println("PROGRAM STUDI YANG ANDA INPUTKAN TELAH TERDAFTAR")
		
										 }
										 
									 }
		
								 }
								 fmt.Print("Press 'Enter' to continue...")
								 bufio.NewReader(os.Stdin).ReadBytes('\n') 
								 CallClear() //Clear Screen 
		
		
								fmt.Println("|======================|")
								fmt.Println("| INPUT DATA MAHASISWA |")
								fmt.Println("|======================|")
		
		
								InputDataMahasiswa(numCamaba, jenis)
								InputPilihanProdi(numCamaba, jenis, numSos)
								numCamaba++
							}
							
			
							gakkepake := "m"
							fmt.Scanln(&gakkepake)
							fmt.Print("Press 'Enter' to continue...")
							bufio.NewReader(os.Stdin).ReadBytes('\n') 
							CallClear() //Clear Screen 
						}else {
							fmt.Println()
							fmt.Println("===== | INVALID | =====")
							fmt.Println()
							gakkepake := "m"
							fmt.Scanln(&gakkepake)
							fmt.Print("Press 'Enter' to continue...")
							bufio.NewReader(os.Stdin).ReadBytes('\n') 
							CallClear() //Clear Screen 
						}
						menu1()
						fmt.Print("Pilihan : ")
						fmt.Scan(&pil)
						CallClear() //Clear Screen
					}
					
				case "2":
					menu2()
					fmt.Scan(&pil)
					CallClear() //Clear Screen
					for pil != "0" {
						if pil == "1" {
							fmt.Print("Pilih [1] Teknik atau [2] Non-Teknik : ")
							fmt.Scan(&jenis)
			
							if jenis == "1" {
								fmt.Println()
								fmt.Print("Masukan Nama Program Studi yang dicari : ")
								
								gakkepake := "m"
								fmt.Scanln(&gakkepake)
			
								key , _ = inputan.ReadString('\n')
								key = strings.TrimSpace(key)
								key = strings.ToUpper(key)
								//fmt.Scan(&key)
								//fmt.Println("Test : " ,key)
								indexCari = SearchNamaProdi(key)
			
								if indexCari == -1 {
									fmt.Println("|=================|")
									fmt.Println("| TIDAK DITEMUKAN |")
									fmt.Println("|=================|")
									fmt.Println()
								}else {
		
									ShowProdiTeknik(indexCari)
									fmt.Println()
									fmt.Print("Lanjut Edit Data Mahasiswa Tekan [Y] : ")
									fmt.Scan(&back)
									if back == "Y" || back == "y" {
										gakkepake := "m"
										fmt.Scanln(&gakkepake)
										
										InputDataProgramStudi(&namaProdi, &akreditasiProdi)
										cek := SearchNamaProdi(namaProdi)
										namaProdi = strings.ToUpper(namaProdi)
										for cek != -1 {
											fmt.Scanln(&gakkepake)
											
											fmt.Println()
											fmt.Println("GAGAL MENAMBAHKAN PROGRAM STUDI")
											fmt.Println()
											fmt.Println("PROGRAM STUDI YANG ANDA INPUTKAN TELAH TERDAFTAR")
											InputDataProgramStudi(&namaProdi, &akreditasiProdi)
											cek = SearchNamaProdi(namaProdi)
										}
										prodiTek[indexCari].nama = namaProdi
										prodiTek[indexCari].akreditasi = akreditasiProdi
										gantiprodimhsteknik(key, namaProdi)
									}
								}
							}else if jenis == "2" {
								fmt.Println()
								fmt.Print("Masukan Nama Program Studi yang dicari : ")
								
								gakkepake := "m"
								fmt.Scanln(&gakkepake)
			
								key , _ = inputan.ReadString('\n')
								key = strings.TrimSpace(key)
								key = strings.ToUpper(key)
								//fmt.Scan(&key)
								indexCari = SearchNamaProdi(key)
			
								if indexCari == -1 {
									fmt.Println("|=================|")
									fmt.Println("| TIDAK DITEMUKAN |")
									fmt.Println("|=================|")
									fmt.Println()
								}else {
									
									ShowProdiNonTeknik(indexCari)
									fmt.Println()
									fmt.Print("Lanjut Edit Data Mahasiswa Tekan [Y] : ")
									fmt.Scan(&back)
									if back == "Y" || back == "y" {
										gakkepake := "m"
										fmt.Scanln(&gakkepake)
										
										InputDataProgramStudi(&namaProdi, &akreditasiProdi)
										cek := SearchNamaProdi(namaProdi)
										namaProdi = strings.ToUpper(namaProdi)
		
										for cek != -1 {
											fmt.Scanln(&gakkepake)
											
											fmt.Println()
											fmt.Println("GAGAL MENAMBAHKAN PROGRAM STUDI")
											fmt.Println()
											fmt.Println("PROGRAM STUDI YANG ANDA INPUTKAN TELAH TERDAFTAR")
											InputDataProgramStudi(&namaProdi, &akreditasiProdi)
											cek = SearchNamaProdi(namaProdi)
										}
										prodiNonTek[indexCari].nama = namaProdi
										prodiNonTek[indexCari].akreditasi = akreditasiProdi
										gantiprodimhsnonteknik(key, namaProdi)
									}
								}	
							}
			
						}else if pil == "2" {
							fmt.Print("Pilih [1] Teknik atau [2] Non-Teknik : ")
							fmt.Scan(&jenis)
							if jenis == "1"{
								fmt.Println()
								fmt.Print("Masukan Nama Mahasiswa yang dicari : ")
			
								gakkepake := "m"
								fmt.Scanln(&gakkepake)
			
								key , _ = inputan.ReadString('\n')
								key = strings.ToUpper(key)
								//fmt.Scan(&key)
								indexCari = SearchNamaMhs(key)
			
								if indexCari == -1 {
									fmt.Println("|=================|")
									fmt.Println("| TIDAK DITEMUKAN |")
									fmt.Println("|=================|")
									fmt.Println()
								}else {
									ShowMhsTeknik(indexCari)
									ShowMhsNonTeknik(indexCari)
									fmt.Println()
									fmt.Print("Lanjut Edit Data Mahasiswa Tekan [Y] : ")
									fmt.Scan(&back)
									CallClear() //Clear Screen 
									if back == "Y" || back == "y" {
										gakkepake := "m"
										fmt.Scanln(&gakkepake)
										fmt.Println("|=====================|")
										fmt.Println("| EDIT DATA MAHASISWA |")
										fmt.Println("|=====================|")
			
										InputDataMahasiswa(indexCari, jenis)
																		
										EditPilihanProdi(indexCari, jenis, numTek)
										
										fmt.Println()
										gakkepake = "m"
										fmt.Scanln(&gakkepake)
										fmt.Print("Press 'Enter' to continue...")
										bufio.NewReader(os.Stdin).ReadBytes('\n') 
										CallClear() //Clear Screen 
									}
								}
								
							}else if jenis == "2"{
								fmt.Println()
								fmt.Print("Masukan Nama Mahasiswa yang dicari : ")
			
								gakkepake := "m"
								fmt.Scanln(&gakkepake)
								
								key , _ = inputan.ReadString('\n')
								key = strings.ToUpper(key)
								//fmt.Scan(&key)
								indexCari = SearchNamaMhs(key)
			
								if indexCari == -1 {
									fmt.Println("|=================|")
									fmt.Println("| TIDAK DITEMUKAN |")
									fmt.Println("|=================|")
									fmt.Println()
								}else {
									ShowMhsTeknik(indexCari)
									ShowMhsNonTeknik(indexCari)
									fmt.Println()
									fmt.Print("Lanjut Edit Data Mahasiswa Tekan [Y] : ")
									fmt.Scan(&back)
									CallClear() //Clear Screen 
									if back == "Y" || back == "y" {
			
										gakkepake := "m"
										fmt.Scanln(&gakkepake)
										fmt.Println("|=====================|")
										fmt.Println("| EDIT DATA MAHASISWA |")
										fmt.Println("|=====================|")
			
					
										InputDataMahasiswa(indexCari, jenis)
										EditPilihanProdi(indexCari, jenis, numSos)
										
										fmt.Println()
										gakkepake = "m"
										fmt.Scanln(&gakkepake)
										fmt.Print("Press 'Enter' to continue...")
										bufio.NewReader(os.Stdin).ReadBytes('\n') 
										CallClear() //Clear Screen 
									}
								}	
							}
							
						}else {
							fmt.Println()
							fmt.Println("===== | INVALID | =====")
							fmt.Println()
							gakkepake := "m"
							fmt.Scanln(&gakkepake)
							fmt.Print("Press 'Enter' to continue...")
							bufio.NewReader(os.Stdin).ReadBytes('\n') 
							CallClear() //Clear Screen 
						}
						menu2()
						fmt.Print("Pilihan : ")
						fmt.Scan(&pil)
						CallClear() //Clear Screen
					}
		
				case "3":
					menu3()
					fmt.Scan(&pil)
					CallClear() //Clear Screen
					for pil != "0" {
						if pil == "1" {
							fmt.Println()
							fmt.Print("Pilih [1] Program Studi Teknik atau [2] Program Studi Non Teknik : ")
							fmt.Scan(&jenis)
							if jenis == "1" || jenis == "2"{
								DeleteProdi(jenis)
								
							}else {
								fmt.Println()
								fmt.Println("===== | INVALID | =====")
								fmt.Println()
								gakkepake := "m"
								fmt.Scanln(&gakkepake)
								fmt.Print("Press 'Enter' to continue...")
								bufio.NewReader(os.Stdin).ReadBytes('\n') 
								CallClear() //Clear Screen 
							}
							
						} else if pil == "2" {
							fmt.Println()
							fmt.Print("Pilih [1] Teknik atau Non Teknik [2] : ")
							fmt.Scan(&jenis)
							if jenis == "1" || jenis == "2" {
								DeleteMhs(jenis)
		
							}else {
								fmt.Println()
								fmt.Println("===== | INVALID | =====")
								fmt.Println()
								gakkepake := "m"
								fmt.Scanln(&gakkepake)
								fmt.Print("Press 'Enter' to continue...")
								bufio.NewReader(os.Stdin).ReadBytes('\n') 
								CallClear() //Clear Screen 
							}
		
							
						}else {
							fmt.Println()
							fmt.Println("===== | INVALID | =====")
							fmt.Println()
							gakkepake := "m"
							fmt.Scanln(&gakkepake)
							fmt.Print("Press 'Enter' to continue...")
							bufio.NewReader(os.Stdin).ReadBytes('\n') 
							CallClear() //Clear Screen 
						}
						menu3()
						fmt.Print("Pilihan : ")
						fmt.Scan(&pil)
						CallClear() //Clear Screen
					}
					
		
		
				case "4":
					menu4()
					fmt.Print("Pilihan : ")
					fmt.Scan(&pil)
					CallClear() //Clear Screen
					for pil != "0" {
						if pil == "1" {
							fmt.Print("Pilih [1] Teknik atau [2] Non-Teknik : ")
							fmt.Scan(&jenis)
			
							if jenis == "1" {
								fmt.Println()
								fmt.Print("Masukan Nama Program Studi yang dicari : ")
								
								gakkepake := "m"
								fmt.Scanln(&gakkepake)
			
								key , _ = inputan.ReadString('\n')
								key = strings.TrimSpace(key)
								key = strings.ToUpper(key)
								//fmt.Scan(&key)
								//fmt.Println("Test : " ,key)
								indexCari = SearchNamaProdi(key)
			
								if indexCari == -1 {
									fmt.Println("|=================|")
									fmt.Println("| TIDAK DITEMUKAN |")
									fmt.Println("|=================|")
									fmt.Println()
								}else {
									ShowProdiTeknik(indexCari)
									fmt.Println()
								}
							}else if jenis == "2" {
								fmt.Println()
								fmt.Print("Masukan Nama Program Studi yang dicari : ")
								
								gakkepake := "m"
								fmt.Scanln(&gakkepake)
			
								key , _ = inputan.ReadString('\n')
								key = strings.TrimSpace(key)
								key = strings.ToUpper(key)
								//fmt.Scan(&key)
								//fmt.Println("Test : " ,key)
								indexCari = SearchNamaProdi(key)
			
								if indexCari == -1 {
									fmt.Println("|=================|")
									fmt.Println("| TIDAK DITEMUKAN |")
									fmt.Println("|=================|")
									fmt.Println()
								}else {
									ShowProdiNonTeknik(indexCari)
									fmt.Println()
								}
			
			
							}
						
						} else if pil == "2" {
							fmt.Println()
							fmt.Print("Nama Mahasiswa yang dicari : ")
							gakkepake := "m"
							fmt.Scanln(&gakkepake)
							key , _ = inputan.ReadString('\n')
							key = strings.ToUpper(key)
							//fmt.Scan(&key)
							indexCari = SearchNamaMhs(key)
			
							if indexCari == -1 {
								fmt.Println("|=================|")
								fmt.Println("| TIDAK DITEMUKAN |")
								fmt.Println("|=================|")
								fmt.Println()
							}else {
								ShowMhsTeknik(indexCari)
								ShowMhsNonTeknik(indexCari)
								fmt.Println()
							}
						}else {
							fmt.Println()
							fmt.Println("===== | INVALID | =====")
							fmt.Println()
							gakkepake := "m"
							fmt.Scanln(&gakkepake)
							fmt.Print("Press 'Enter' to continue...")
							bufio.NewReader(os.Stdin).ReadBytes('\n') 
							CallClear() //Clear Screen 
						}
						menu4()
						fmt.Print("Pilihan : ")
						fmt.Scan(&pil)
						CallClear() //Clear Screen
					
					}
					
		
				case "5":
					menu5()
					fmt.Print("Pilihan : ")
					fmt.Scan(&pil)
					CallClear() //Clear Screen
					for pil != "0" {
						if pil == "1" {
						
							if numTek == 0 {
								fmt.Println("== | Belum Ada Program Studi Teknik | ==")
							}else {
								fmt.Println("====== Program Studi Teknik ======")
								for i := 0 ; i < numTek ; i++ {
								ShowProdiTeknik(i)
								fmt.Println(PemilihProdTeknik(i)," Pemilih")
								fmt.Println()
								}
							fmt.Println("Total Program Studi Teknik :", numTek)
							}
							
							if numSos == 0 {
								fmt.Println("== | Belum Ada Program Studi Non Teknik | ==")
							}else {
								fmt.Println("====== Program Studi Non Teknik ======")
								for i := 0 ; i < numSos ; i++ {
								ShowProdiNonTeknik(i)
								fmt.Println(PemilihProdNonTeknik(i)," Pemilih")
								fmt.Println()
								}
							fmt.Println("Total Program Studi Non Teknik :", numSos)
							}
							
			
							fmt.Println()
			
						}else if pil == "2" {
							menu5_2()
							fmt.Print("Pilihan : ")
							fmt.Scan(&pil)
							CallClear() //Clear Screen
							for pil != "0" && numCamaba != 0 {
								if pil == "1" {
									SortRataTotal(numCamaba)
									SortNilaiTPA(numCamaba)
									fmt.Println("|=====================================================|")
									fmt.Println("| DATA TERURUT BERDASARKAN NILAI TES POTENSI AKADEMIK |")
									fmt.Println("|=====================================================|")
								}else if pil == "2" {
									SortRataTotal(numCamaba)
									SortNilaiTBI(numCamaba)
									fmt.Println("|===================================================|")
									fmt.Println("| DATA TERURUT BERDASARKAN NILAI TES BAHASA INGGRIS |")
									fmt.Println("|===================================================|")
								}else if pil == "3" {
									SortRataTotal(numCamaba)
									fmt.Println("|================================================|")
									fmt.Println("| DATA TERURUT BERDASARKAN NILAI RATA-RATA TOTAL |")
									fmt.Println("|================================================|")
								}else if pil == "4" {
									SortingString(numCamaba)
									fmt.Println("|=========================================|")
									fmt.Println("| DATA TERURUT BERDASARKAN NAMA MAHASISWA |")
									fmt.Println("|=========================================|")
								}else if pil == "5" {
		
									if SyaratDiTerima == 0 {
										fmt.Println("MASUKAN TERLEBIH DAHULU BATAS NILAI MINIMUM")
										fmt.Print(" Masukan Batas Nilai Minumm Untuk Diterima [1 ... 100] : ")
										fmt.Scan(&SyaratDiTerima)
									}else {
										SortRataTotal(numCamaba)
										fmt.Println("|==============================|=============================|")
										fmt.Printf("| PENDAFTAR YANG LULUS SELEKSI | BATAS NILAI MINIMUM : %.2f  |\n" , SyaratDiTerima)
										fmt.Println("|==============================|=============================|")
										fmt.Println()
										
										ShowMhsLulus(numCamaba)
		
										fmt.Println()
		
										gakkepake := "m"
										fmt.Scanln(&gakkepake)
										fmt.Print("Press 'Enter' to continue...")
										bufio.NewReader(os.Stdin).ReadBytes('\n') 
										CallClear() //Clear Screen 
									}
		
		
		
								}else if pil == "6" {
		
									if SyaratDiTerima == 0 {
										fmt.Println("MASUKAN TERLEBIH DAHULU BATAS NILAI MINIMUM")
										fmt.Print(" Masukan Batas Nilai Minumm Untuk Diterima [1 ... 100] : ")
										fmt.Scan(&SyaratDiTerima)
									}else {
										SortRataTotal(numCamaba)
										fmt.Println("|====================================|=============================|")
										fmt.Printf("| PENDAFTAR YANG TIDAK LULUS SELEKSI | BATAS NILAI MINIMUM : %.2f  |\n" , SyaratDiTerima)
										fmt.Println("|====================================|=============================|")
										fmt.Println()
			
										ShowMhsTidakLulus(numCamaba)
										
										fmt.Println()
		
										gakkepake := "m"
										fmt.Scanln(&gakkepake)
										fmt.Print("Press 'Enter' to continue...")
										bufio.NewReader(os.Stdin).ReadBytes('\n') 
										CallClear() //Clear Screen 
									}
		
		
								
								}else {
									fmt.Println()
									fmt.Println("===== | INVALID | =====")
									fmt.Println()
									gakkepake := "m"
									fmt.Scanln(&gakkepake)
									fmt.Print("Press 'Enter' to continue...")
									bufio.NewReader(os.Stdin).ReadBytes('\n') 
									CallClear() //Clear Screen 
								}
			
								if pil == "1" || pil == "2" || pil == "3" || pil == "4"  {
									fmt.Println("====== Mahasiswa Teknik ======")
									for i := 0 ; i < numCamaba ; i++ {
										// fmt.Printf("[%d] \n" , i)
										ShowMhsTeknik(i)
										fmt.Println()
									}
									fmt.Println("Total Mahasiswa Teknik :", numMabaTek)
					
									fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
					
									fmt.Println("====== Mahasiswa Non Teknik ======")
									for i := 0 ; i < numCamaba ; i++ {
										// fmt.Printf("[%d] \n" , i)
										ShowMhsNonTeknik(i)
										fmt.Println()
									}
									fmt.Println("Total Mahasiswa Non Teknik :", numMabaSos)
					
									fmt.Println()
									fmt.Println()
									fmt.Println("Jumlah Pendaftar : ",numCamaba)
									fmt.Println()
									fmt.Println()
		
									gakkepake := "m"
									fmt.Scanln(&gakkepake)
									fmt.Print("Press 'Enter' to continue...")
									bufio.NewReader(os.Stdin).ReadBytes('\n') 
									CallClear() //Clear Screen 
								}
								menu5_2()
								fmt.Print("Pilihan : ")
								fmt.Scan(&pil)
								CallClear() //Clear Screen
							}
							if numCamaba == 0 {
								CallClear() //Clear Screen
								fmt.Println()
								fmt.Println("| BELUM ADA PENDAFTAR |")
								fmt.Println()
								gakkepake := "m"
								fmt.Scanln(&gakkepake)
								fmt.Print("Press 'Enter' to continue...")
								bufio.NewReader(os.Stdin).ReadBytes('\n') 
								CallClear() //Clear Screen 
							}
							
						}else {
							fmt.Println()
							fmt.Println("===== | INVALID | =====")
							fmt.Println()
							gakkepake := "m"
							fmt.Scanln(&gakkepake)
							fmt.Print("Press 'Enter' to continue...")
							bufio.NewReader(os.Stdin).ReadBytes('\n') 
							CallClear() //Clear Screen 
						}
						menu5()
						fmt.Print("Pilihan : ")
						fmt.Scan(&pil)
						CallClear() //Clear Screen
					}
					
				case "6":
					fmt.Print(" Masukan Batas Nilai Minumm Untuk Diterima [1 ... 100] : ")
					fmt.Scan(&SyaratDiTerima)
					for SyaratDiTerima < 1 || SyaratDiTerima > 100 {
						fmt.Println("Invalid")
						fmt.Print(" Masukan Batas Nilai Minumm Untuk Diterima [1 ... 100] : ")
						fmt.Scan(&SyaratDiTerima)
					}
					
				
				default:
					fmt.Println()
					fmt.Println("===== | INVALID | =====")
					fmt.Println()
					gakkepake := "m"
					fmt.Scanln(&gakkepake)
					fmt.Print("Press 'Enter' to continue...")
					bufio.NewReader(os.Stdin).ReadBytes('\n') 
					CallClear() //Clear Screen 
				}
				menuAwal()
				fmt.Print("Pilihan : ")
				fmt.Scan(&pil)
				CallClear() //Clear Screen
			}
		}else {
			fmt.Println()
			fmt.Println("===== | INVALID | =====")
			fmt.Println()
			gakkepake := "m"
			fmt.Scanln(&gakkepake)
			fmt.Print("Press 'Enter' to continue...")
			bufio.NewReader(os.Stdin).ReadBytes('\n') 
			CallClear() //Clear Screen 
		}

		fmt.Print("Yakin Ingin Keluar Program [Y/N] : ")
		fmt.Scan(&konfirmasi)
		CallClear() //Clear Screen
	}

	fmt.Println()
	fmt.Println("=== | PROGRAM SELESAI | ===")
	
}

// ============================================================  KUMPULAN FUNGSI DAN PROSEDUR ================================================================//
// ===========================================================================================================================================================//

func menuAwal(){
	fmt.Println("|======== Menu Utama ======|")
	fmt.Println("| 1. Tambah Data           |")
	fmt.Println("| 2. Edit Data             |")
	fmt.Println("| 3. Hapus Data            |")
	fmt.Println("| 4. Cari Data             |")
	fmt.Println("| 5. Tampilkan Data        |")
	fmt.Println("| 6. Input Syarat Diterima |")
	fmt.Println("| 0. Keluar                |")
	fmt.Println("|==========================|")
}

func menu1(){
	fmt.Println("|==============================|")
	fmt.Println("| 1. Tambah Data Program Studi |")
	fmt.Println("| 2. Tambah Data Mahasiswa     |")
	fmt.Println("| 0. Kembali                   |")
	fmt.Println("|==============================|")
}

func menu2(){
	fmt.Println("|============================|")
	fmt.Println("| 1. Edit Data Program Studi |")
	fmt.Println("| 2. Edit Data Mahasiswa     |")
	fmt.Println("| 0. Kembali                 |")
	fmt.Println("|============================|")
}

func menu3(){
	fmt.Println("|=============================|")
	fmt.Println("| 1. Hapus Data Program Studi |")
	fmt.Println("| 2. Hapus Data Mahasiswa     |")
	fmt.Println("| 0. Kembali                  |")
	fmt.Println("|=============================|")
}

func menu4(){
	fmt.Println("|=============================|")
	fmt.Println("| 1. Cari Program Studi       |")
	fmt.Println("| 2. Cari Data Mahasiswa      |")
	fmt.Println("| 0. Kembali                  |")
	fmt.Println("|=============================|")

}
func menu5(){
	fmt.Println("|=================================|")
	fmt.Println("| 1. Tampilkan Data Program Studi |")
	fmt.Println("| 2. Tampilkan Data Mahasiswa     |")
	fmt.Println("| 0. Kembali                      |")
	fmt.Println("|=================================|")
}
func menu5_2(){
	fmt.Println("|================================================|")
	fmt.Println("| 1. Terurut Berdasarkan Nilai TPA               |")
	fmt.Println("| 2. Terurut Berdasarkan Nilai TBI               |")
	fmt.Println("| 3. Terurut Berdasarkan Nilai Rata-rata Total   |")
	fmt.Println("| 4. Terurut Berdasarkan Nama Mahasiswa          |")
	fmt.Println("| 5. Pendaftar yang Lulus                        |")
	fmt.Println("| 6. Pendaftar yang Tidak Lulus                  |")
	fmt.Println("| 0. Kembali                                     |")
	fmt.Println("|================================================|")
}



// Tampilkan Data

func ShowProdiTeknik(prodiKe int) {
	fmt.Printf("[%d] %s Akreditasi: %s ", prodiKe+1, prodiTek[prodiKe].nama, prodiTek[prodiKe].akreditasi)
}

func ShowProdiNonTeknik(prodiKe int){
	fmt.Printf("[%d] %s Akreditasi: %s ", prodiKe+1, prodiNonTek[prodiKe].nama, prodiNonTek[prodiKe].akreditasi)
}

func ShowMhsTeknik(mhsKe int) {
	if data[mhsKe].prodiTektek[0].nama != "" {
		// MAHASISWA TEKNIK
		fmt.Print("Nama                    :  " , strings.ToUpper(data[mhsKe].nama))
		fmt.Print("Prodi			:  ")
		if data[mhsKe].numprodi > 0 {
			i := 0
			fmt.Print("[",i+1,"].", data[mhsKe].prodiTektek[i].nama )
			for i=1; i<data[mhsKe].numprodi; i++ {
				fmt.Print(" , [",i+1,"].", data[mhsKe].prodiTektek[i].nama )
			}
		}
		fmt.Println()
		fmt.Println("Jenis Kelamin		: " , strings.ToUpper(data[mhsKe].gender))
		fmt.Println("Nilai Matematika	: " , data[mhsKe].nilaiMat)
		fmt.Println("Nilai IPA		: " , data[mhsKe].nilaiIPA)
		fmt.Println("Nilai Bahasa Inggris	: " , data[mhsKe].nilaiInggris)
		fmt.Println("Nilai TPA		: " , data[mhsKe].nilaiTPA)
		fmt.Println("Nilai TBI		: " , data[mhsKe].nilaiTBI)
		fmt.Println("========================================")
		fmt.Printf("| Rata-rata Nilai Raport      :  %.2f |\n" , data[mhsKe].rataRaport)
		fmt.Println("========================================")
		fmt.Printf("| Rata-rata Nilai TPA dan TBI :  %.2f |\n" , data[mhsKe].rataTesMasuk)
		fmt.Println("========================================")
		fmt.Printf("| Rata-rata Nilai Total       :  %.2f |\n" , data[mhsKe].rataSyarat)
		fmt.Println("=======================================")
	}
}
func ShowMhsNonTeknik(mhsKe int ) {
	if data[mhsKe].prodiNonTektek[0].nama != "" {
		// MAHASISWA NON TEKNIK
		fmt.Print("Nama                    :  " , strings.ToUpper(data[mhsKe].nama))
		fmt.Print("Prodi			:  ")
		if data[mhsKe].numprodi > 0 {
			i := 0
			fmt.Print("[",i+1,"].", data[mhsKe].prodiNonTektek[i].nama )
			for i=1; i<data[mhsKe].numprodi; i++ {
				fmt.Print(" , [",i+1,"].", data[mhsKe].prodiNonTektek[i].nama )
			}
		}
		fmt.Println()
		fmt.Println("Jenis Kelamin		: " , strings.ToUpper(data[mhsKe].gender))
		fmt.Println("Nilai Matematika	: " , data[mhsKe].nilaiMat)
		fmt.Println("Nilai IPS		: " , data[mhsKe].nilaiIPS)
		fmt.Println("Nilai Bahasa Inggris	: " , data[mhsKe].nilaiInggris)
		fmt.Println("Nilai TPA		: " , data[mhsKe].nilaiTPA)
		fmt.Println("Nilai TBI		: " , data[mhsKe].nilaiTBI)
		fmt.Println("|======================================|")
		fmt.Printf("| Rata-rata Nilai Raport      :  %.2f |\n" , data[mhsKe].rataRaport)
		fmt.Println("|======================================|")
		fmt.Printf("| Rata-rata Nilai TPA dan TBI :  %.2f |\n" , data[mhsKe].rataTesMasuk)
		fmt.Println("|======================================|")
		fmt.Printf("| Rata-rata Nilai Total       :  %.2f |\n" , data[mhsKe].rataSyarat)
		fmt.Println("|=====================================|")
	}
}

func ShowMhsLulus(numCamaba int){
	fmt.Println("====== Mahasiswa Teknik ======")
	for i := 0 ; i < numCamaba ; i++ {
		// fmt.Printf("[%d] \n" , i)
		if data[i].prodiTektek[0].nama != "" && data[i].rataSyarat >= SyaratDiTerima{
			
			ShowMhsTeknik(i)
			fmt.Println()
			
		}

	}

	fmt.Println("====== Mahasiswa Non Teknik ======")
	for i := 0 ; i < numCamaba ; i++ {
		// fmt.Printf("[%d] \n" , i)
		if data[i].prodiNonTektek[0].nama != "" && data[i].rataSyarat >= SyaratDiTerima {
			
			ShowMhsNonTeknik(i)
			fmt.Println()
			
		}
	}
}

func ShowMhsTidakLulus(numCamaba int){
	fmt.Println("====== Mahasiswa Teknik ======")
	for i := 0 ; i < numCamaba ; i++ {
		// fmt.Printf("[%d] \n" , i)
		if data[i].prodiTektek[0].nama != "" && data[i].rataSyarat < SyaratDiTerima{
			
			ShowMhsTeknik(i)
			fmt.Println()
			
		}

	}

	fmt.Println("====== Mahasiswa Non Teknik ======")
	for i := 0 ; i < numCamaba ; i++ {
		// fmt.Printf("[%d] \n" , i)
		if data[i].prodiNonTektek[0].nama != "" && data[i].rataSyarat < SyaratDiTerima {
			
			ShowMhsNonTeknik(i)
			fmt.Println()
			
		}


	}
}

// UNTUK MENAMBAHKAN PROGRAM STUDI
func InputDataProgramStudi(nama, akreditasi *string){
	
	
	fmt.Print("Nama Program Studi : ")
	*nama, _ = inputan.ReadString('\n')
	*nama = strings.TrimSpace(*nama)
	*nama = strings.ToUpper(*nama)

	
	
	//fmt.Scan(&*nama)
	fmt.Print("Akreditasi ", *nama ," : ")
	fmt.Scan(&*akreditasi)
	*akreditasi = strings.ToUpper(*akreditasi)
}
// input jenis dulu buat milih teknik / nontek
func InputDataMahasiswa(numCamaba int ,  jenis string){

	fmt.Print("Nama : ")
	data[numCamaba].nama, _ = inputan.ReadString('\n')
	data[numCamaba].nama = strings.ToUpper(data[numCamaba].nama)
	//fmt.Scan(&data[numCamaba].nama)
	fmt.Print("Jenis Kelamin [L] / [P]	: ")
	fmt.Scan(&data[numCamaba].gender)
	data[numCamaba].gender = strings.ToUpper(data[numCamaba].gender) 
	for data[numCamaba].gender != "L"  && data[numCamaba].gender != "P"{
		fmt.Println("Invalid")
		fmt.Print("Jenis Kelamin [L] / [P]	: ")
		fmt.Scan(&data[numCamaba].gender)
	} 
	fmt.Print("Nilai Matematika : ")
	fmt.Scan(&data[numCamaba].nilaiMat)
	for data[numCamaba].nilaiMat < 0 || data[numCamaba].nilaiMat > 100{
		fmt.Println("Invalid")
		fmt.Print("Nilai Matematika : ")
		fmt.Scan(&data[numCamaba].nilaiMat)
	}
	// jenis 1 = teknik | jenis 2 = nontek
	if jenis == "1" {
		fmt.Print("Nilai IPA : ")
		fmt.Scan(&data[numCamaba].nilaiIPA)
		for data[numCamaba].nilaiIPA < 0 || data[numCamaba].nilaiIPA > 100{
			fmt.Println("Invalid")
			fmt.Print("Nilai IPA : ")
			fmt.Scan(&data[numCamaba].nilaiIPA)
		}
	} else if jenis == "2" {
		fmt.Print("Nilai IPS : ")
		fmt.Scan(&data[numCamaba].nilaiIPS)
		for data[numCamaba].nilaiIPS < 0 || data[numCamaba].nilaiIPS > 100 {
			fmt.Println("Invalid")
			fmt.Print("Nilai IPS : ")
			fmt.Scan(&data[numCamaba].nilaiIPS)
		} 
	}
	fmt.Print("Nilai Bahasa Inggris : ")
	fmt.Scan(&data[numCamaba].nilaiInggris)
	for data[numCamaba].nilaiInggris < 0 || data[numCamaba].nilaiInggris > 100 {
		fmt.Println("Invalid")
		fmt.Print("Nilai Bahasa Inggris : ")
		fmt.Scan(&data[numCamaba].nilaiInggris)
	}
	fmt.Print("Nilai TPA : ")
	fmt.Scan(&data[numCamaba].nilaiTPA)
	for data[numCamaba].nilaiTPA < 0 || data[numCamaba].nilaiTPA > 100 {
		fmt.Println("Invalid")
		fmt.Print("Nilai TPA : ")
		fmt.Scan(&data[numCamaba].nilaiTPA)
	}
	fmt.Print("Nilai TBI : ")
	fmt.Scan(&data[numCamaba].nilaiTBI)
	for data[numCamaba].nilaiTBI < 0 || data[numCamaba].nilaiTBI > 100 {
		fmt.Println("Invalid")
		fmt.Print("Nilai TBI : ")
		fmt.Scan(&data[numCamaba].nilaiTBI)
	}
	if jenis == "1" {
		data[numCamaba].rataRaport = (data[numCamaba].nilaiIPA + data[numCamaba].nilaiMat + data[numCamaba].nilaiInggris)/3
	} else if jenis == "2" {
		data[numCamaba].rataRaport = (data[numCamaba].nilaiIPS + data[numCamaba].nilaiMat + data[numCamaba].nilaiInggris)/3
	}
	data[numCamaba].rataTesMasuk = (data[numCamaba].nilaiTPA + data[numCamaba].nilaiTBI)/2 

	data[numCamaba].rataSyarat = (data[numCamaba].rataRaport + data[numCamaba].rataTesMasuk)/2
}

// UNTUK MENAMBAHKAN DAFTAR PRODI KE DALLAM ARRAY PILIHAN PRODI MAHASISWA YG BERJUMLAH 3
func InputPilihanProdi(numCamaba int , jenis string, idxprodi int){
	//output prodi
	//error handling ganti ke string jgn lupa
	var pil int
	awal := data[numCamaba].numprodi
	if jenis == "1" {
		fmt.Println("_______________________________")
		fmt.Println("| DAFTAR PROGRAM STUDI TEKNIK |")
		fmt.Println("|=============================|")
		for i:=0; i<idxprodi; i++{
			fmt.Println("[", i+1, "] ", prodiTek[i].nama, prodiTek[i].akreditasi)
		}
		fmt.Println("Pilih Program Studi Teknik [Cukup inputkan indeks Prodi Saja]")
		fmt.Println()
		for i:=awal; i<awal+3; i++{
			fmt.Printf("Program Studi Pilihan %d : " , i+1)
			fmt.Scan(&pil)
			data[numCamaba].numprodi++
			data[numCamaba].prodiTektek[i].nama = prodiTek[pil-1].nama
			data[numCamaba].prodiTektek[i].akreditasi = prodiTek[pil-1].akreditasi
		}
		numMabaTek++
	 }else if jenis == "2" {
		fmt.Println("___________________________________")
		fmt.Println("| DAFTAR PROGRAM STUDI NON TEKNIK |")
		fmt.Println("|=================================|")
		for i:=0; i<idxprodi; i++{
			fmt.Println("[", i+1, "] ", prodiNonTek[i].nama, prodiNonTek[i].akreditasi)
		}
		fmt.Println("Pilih Program Studi Non Teknik [Cukup inputkan indeks Prodi Saja]")
		fmt.Println()
		for i:=awal; i<awal+3; i++{
			fmt.Printf("Program Studi Pilihan %d : " , i+1)
			fmt.Scan(&pil)
			data[numCamaba].numprodi++
			data[numCamaba].prodiNonTektek[i].nama = prodiNonTek[pil-1].nama
			data[numCamaba].prodiNonTektek[i].akreditasi = prodiNonTek[pil-1].akreditasi
		}
		numMabaSos++
	}
}

// DIPAKAI PADA MENU EDIT MAHASISWA
func EditPilihanProdi(numCamaba int , jenis string, idxprodi int){
	//output prodi
	//error handling ganti ke string jgn lupa
	var pil int
	//awal := data[numCamaba].numprodi
	if jenis == "1" {
		fmt.Println("_______________________________")
		fmt.Println("| DAFTAR PROGRAM STUDI TEKNIK |")
		fmt.Println("|=============================|")
		for i:=0; i<idxprodi; i++{
			fmt.Println("[", i+1, "] ", prodiTek[i].nama, prodiTek[i].akreditasi)
		}
		fmt.Println("Pilih Program Studi Teknik [Cukup inputkan indeks Prodi Saja]")
		fmt.Println()
		for i:=0; i<3; i++{
			fmt.Printf("Program Studi Pilihan %d : " , i+1)
			fmt.Scan(&pil)
			
			data[numCamaba].prodiTektek[i].nama = prodiTek[pil-1].nama
			data[numCamaba].prodiTektek[i].akreditasi = prodiTek[pil-1].akreditasi
		}
		
	 }else if jenis == "2" {
		fmt.Println("___________________________________")
		fmt.Println("| DAFTAR PROGRAM STUDI NON TEKNIK |")
		fmt.Println("|=================================|")
		for i:=0; i<idxprodi; i++{
			fmt.Println("[", i+1, "] ", prodiNonTek[i].nama, prodiNonTek[i].akreditasi)
		}
		fmt.Println("Pilih Program Studi Non Teknik [Cukup inputkan indeks Prodi Saja]")
		fmt.Println()
		for i:=0; i<3; i++{
			fmt.Printf("Program Studi Pilihan %d : " , i+1)
			fmt.Scan(&pil)
			
			data[numCamaba].prodiNonTektek[i].nama = prodiNonTek[pil-1].nama
			data[numCamaba].prodiNonTektek[i].akreditasi = prodiNonTek[pil-1].akreditasi
		}
		
	}
}


// UNTUK MENGHAPUS PROGRAM STUDI
func DeleteProdi(jenis string){
	var pil int
	if jenis == "1"{
		if numTek == 0{
			fmt.Println("== | Belum Ada Program Studi Teknik | ==")
		}else{
			for i:=0; i<numTek; i++{
				fmt.Println("[", i+1, "] ", prodiTek[i].nama, prodiTek[i].akreditasi)
			}
			fmt.Println("Pilih Program Studi Yang Akan di Hapus")
			fmt.Print("Pilihan : ")
			fmt.Scan(&pil)
			if pil > numTek || pil <= 0 {
				fmt.Println()
				fmt.Println("===== | INVALID | =====")
				fmt.Println()
				gakkepake := "m"
				fmt.Scanln(&gakkepake)
				fmt.Print("Press 'Enter' to continue...")
				bufio.NewReader(os.Stdin).ReadBytes('\n') 
				CallClear() //Clear Screen 
			} else {
				DeleteProdiMhs(jenis, prodiTek[pil-1].nama)
				for pil-2 < numTek{
					prodiTek[pil-1] = prodiTek[pil] 
					pil++
				}
				fmt.Println("|BERHASIL DIHAPUS|")
				numTek--
			}
			
		}
	}else if jenis == "2"{
		if numSos == 0{
			fmt.Println("== | Belum Ada Program Studi Non Teknik | ==")
		}else{
			for i:=0; i<numSos; i++{
				fmt.Println("[", i+1, "] ", prodiNonTek[i].nama, prodiNonTek[i].akreditasi)
			}
			fmt.Println("Pilih Program Studi Yang Akan di Hapus")
			fmt.Print("Pilihan : ")
			fmt.Scan(&pil)
			if pil > numSos || pil <= 0 {
				fmt.Println()
				fmt.Println("===== | INVALID | =====")
				fmt.Println()
				gakkepake := "m"
				fmt.Scanln(&gakkepake)
				fmt.Print("Press 'Enter' to continue...")
				bufio.NewReader(os.Stdin).ReadBytes('\n') 
				CallClear() //Clear Screen 
			}else {
				DeleteProdiMhs(jenis, prodiTek[pil-1].nama)
				for pil-2 < numSos{
					prodiNonTek[pil-1] = prodiNonTek[pil] 
					pil++
				}
				fmt.Println("|BERHASIL DIHAPUS|")
				numSos--
			}

		}
	}
}

// UNTUK MENGHAPUS PROGRAM STUDI PADA MAHASISWA
func DeleteProdiMhs(pil, key string) {

	if pil == "1" {
		for i:=0; i<numMabaTek; i++ {
			for j:=0; j<3; j++ {
				if data[i].prodiTektek[j].nama == key {
					for k:=j; k<2; k++ {
						data[i].prodiTektek[k] = data[i].prodiTektek[k+1]
					}
					data[i].prodiTektek[2].nama = ""
					data[i].prodiTektek[2].akreditasi = ""
					data[i].numprodi--
				}
			}
		}
	} else if pil == "2" {
		for i:=0; i<numMabaSos; i++ {
			for j:=0; j<2; j++ {
				if data[i].prodiNonTektek[j].nama == key {
					for k:=j; k<2; k++ {
						data[i].prodiNonTektek[k] = data[i].prodiNonTektek[k+1]
					}
					data[i].prodiNonTektek[2].nama = ""
					data[i].prodiNonTektek[2].akreditasi = ""
					data[i].numprodi--
				}
			}
		}
	}
}

// MENGHAPUS SELURUH DATA MAHASISWA yang DIPILIH
func DeleteMhs(jenis string){
	var pil, i ,no int

	

	if jenis == "1" {
		i=0
		if numMabaTek == 0 {
			fmt.Println("| BELUM ADA MAHASISWA TEKNIK |")
		}else{
			
			for i<numMabaTek {
				if data[i].prodiTektek[0].nama != "" {
					fmt.Printf("[ %d ]. %s\n", no+1, data[i].nama)
					no++
				}
				i++
			}
			fmt.Println("Pilih Mahasiswa Teknik Yang Akan di Hapus")
			fmt.Print("Pilihan : ")
			fmt.Scan(&pil)
			if pil > numMabaTek || pil <= 0 {
				fmt.Println()
				fmt.Println("===== | INVALID | =====")
				fmt.Println()
				gakkepake := "m"
				fmt.Scanln(&gakkepake)
				fmt.Print("Press 'Enter' to continue...")
				bufio.NewReader(os.Stdin).ReadBytes('\n') 
				CallClear() //Clear Screen 
			}else {
				for pil-2 < numMabaTek{
					data[pil-1] = data[pil]
					pil++
				}
				fmt.Println("|BERHASIL DIHAPUS|")
				numMabaTek--
				numCamaba--
				
			}
	
			
		}


	
	} else if jenis == "2" {
		i=0
		if numMabaSos == 0 {
			fmt.Println("| BELUM ADA MAHASISWA NON TEKNIK |")
		}else{
			
			for i<numMabaSos {
				if data[i].prodiNonTektek[0].nama != "" {
					fmt.Printf("[ %d ]. %s\n", no+1, data[i].nama)
					no++
				}
				i++
			}
			fmt.Println("Pilih Mahasiswa Non Teknik Yang Akan di Hapus")
			fmt.Print("Pilihan : ")
			fmt.Scan(&pil)
			if pil > numMabaSos || pil <= 0 {
				fmt.Println()
				fmt.Println("===== | INVALID | =====")
				fmt.Println()
				gakkepake := "m"
				fmt.Scanln(&gakkepake)
				fmt.Print("Press 'Enter' to continue...")
				bufio.NewReader(os.Stdin).ReadBytes('\n') 
				CallClear() //Clear Screen 
			}else {
				for pil-2 < numMabaSos{
					data[pil-1] = data[pil]
					pil++
				}
				fmt.Println("|BERHASIL DIHAPUS|")
				numMabaSos--
				numCamaba--
				
			}
	
			
		}

	}
	
}

// MENGEMBALIKAN JUMLAH PEMILIH PADA SETIAP PRODI TEKNIK
func PemilihProdTeknik(i int) int{
	var jum int

	for j:=0; j<numMabaTek; j++ {
		for k:=0; k<3 ; k++ {
			if prodiTek[i].nama == data[j].prodiTektek[k].nama{
				jum++
			}	
		}	
	}
	return jum
}

// MENGEMBALIKAN JUMLAH PEMILIH PADA SETIAP PRODI NON TEKNIK 
func PemilihProdNonTeknik(i int) int{
	var jum int

	for j:=0; j<numMabaSos; j++ {
		for k:=0; k<3 ; k++ {
			if prodiNonTek[i].nama == data[j].prodiNonTektek[k].nama{
				jum++
			}
			
		}
		
	}
	return jum
}

// Digunakan pada Menu 5 
func SortRataTotal(BanyaknyaMahasiswa int){
	var max int
	
	for i:=0; i<BanyaknyaMahasiswa-1; i++{
		max = i
		
		for j := i+1; j<BanyaknyaMahasiswa; j++{
			if data[j].rataSyarat > data[max].rataSyarat{
				max = j
			}
		}
		// SWAP
		data[max], data[i] = data[i], data[max]
	}
}

func SortNilaiTPA(BanyaknyaMahasiswa int){
	var max int
	
	for i:=0; i<BanyaknyaMahasiswa-1; i++{
		max = i
		
		for j := i+1; j<BanyaknyaMahasiswa; j++{
			if data[j].nilaiTPA > data[max].nilaiTPA{
				max = j
			}
		}
		// SWAP
		data[max], data[i] = data[i], data[max]
	}
}

func SortNilaiTBI(BanyaknyaMahasiswa int){
	var max int
	
	for i:=0; i<BanyaknyaMahasiswa-1; i++{
		max = i
		
		for j := i+1; j<BanyaknyaMahasiswa; j++{
			if data[j].nilaiTBI > data[max].nilaiTBI{
				max = j
			}
		}
		// SWAP
		data[max], data[i] = data[i], data[max]
	}
}


// UNTUK MENGURUTKAN BERDASARKAN NAMA MAHASISWA
func SortingString(BanyaknyaMahasiswa int) {
	// Sorting Nama Mahasiswa
	var min int
	for i := 0; i < BanyaknyaMahasiswa; i++ {
		min = i
		for j := i + 1; j < BanyaknyaMahasiswa; j++ {
			if data[j].nama < data[min].nama {
				min = j
			}
		}
		// SWAP
		data[min], data[i] = data[i], data[min]
	}

}



func SearchNamaMhs(key string)int {
	i := 0
	for (data[i].nama != key) && (i<numCamaba) {
		i++
	}
	if data[i].nama == key {
		return i
	}else {
		return -1
	}
}


func SearchNamaProdi(key string)int {
	
	i := 0
	for (prodiTek[i].nama != key && prodiNonTek[i].nama != key) && (i<numSos || i<numTek) {
		i++
	}
	if prodiTek[i].nama == key || prodiNonTek[i].nama == key {
		return i
	}else {
		return -1
	}
	
}

// UNTUK MENGGANTI PRODI PADA MAHASISWA SETELAH DI EDIT
func gantiprodimhsteknik(key, namaProdi string) {
	
	for i:=0; i<numCamaba; i++ {
		for j:=0; j<3; j++ {
			if data[i].prodiTektek[j].nama == key {
				data[i].prodiTektek[j].nama = namaProdi
			}
		}
	}
	
}

func gantiprodimhsnonteknik(key, namaProdi string) {
	
	for i:=0; i<numCamaba; i++ {
		for j:=0; j<3; j++ {
			if data[i].prodiNonTektek[j].nama == key {
				data[i].prodiNonTektek[j].nama = namaProdi
			}
		}
	}
	
}
/*
///////////////////////////////////////  FUNGSI TAMBAHAN  //////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////
*/

func init() {
    clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() { 
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}



func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}
