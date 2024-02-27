package main

import "fmt"

type Mahasiswa struct {
	ID        int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func CheckID(id int) {
	var Res = []Mahasiswa{
		{ID: 1, Nama: "Ilham Ramadhan", Alamat: "Jakarta", Pekerjaan: "Freelance", Alasan: "Ingin memperdalam ilmu Backend Terutama golang"},
		{ID: 2, Nama: "Dimas", Alamat: "Bekasi", Pekerjaan: "UI/UX", Alasan: "Ingin mengikuti pelatihan terkait Backend"},
		{ID: 3, Nama: "Luthfi", Alamat: "Bogor", Pekerjaan: "Fresh Graduate", Alasan: "Ingin menjadi seorang QA Engineer"},
	}

	for _, v := range Res {
		if v.ID == id {
			fmt.Printf("Absen %d\n", v.ID)
			fmt.Printf("Nama: %s\n", v.Nama)
			fmt.Printf("Alamat: %s\n", v.Alamat)
			fmt.Printf("Pekerjaan: %s\n", v.Pekerjaan)
			return
		}
	}

	fmt.Println(id)
}

func main() {
	var exm int
	fmt.Printf("Masukkann Input Absen : ")
	fmt.Scan(&exm)
	CheckID(exm)
}
