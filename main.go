package main

import (
	"fmt"
)

type Teman struct {
	Description []Mahasiswa
}

type Mahasiswa struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func Test() Teman {

	mhs := Mahasiswa{
		Nama:      "Ilham",
		Alamat:    "Jakarta Timur",
		Pekerjaan: "Freelance",
		Alasan:    "Ingin memperdalam ilmu Backend Terutama golang",
	}

	mhs2 := Mahasiswa{
		Nama:      "Rahmat",
		Alamat:    "Bekasi",
		Pekerjaan: "Mahasiswa",
		Alasan:    "Ingin mengasah ilmu Backend dan meraih sertifikat internasional",
	}

	Result := Teman{
		Description: []Mahasiswa{
			mhs,
			mhs2,
		},
	}

	return Result
}

func main() {
	Hasil := Test()

	for i, mhs := range Hasil.Description {
		fmt.Printf("Mahasiswa %d:\n", i+1)
		fmt.Printf("Nama: %s\n", mhs.Nama)
		fmt.Printf("Alamat: %s\n", mhs.Alamat)
		fmt.Printf("Pekerjaan: %s\n", mhs.Pekerjaan)
		fmt.Printf("Alasan: %s\n\n", mhs.Alasan)
	}
}
