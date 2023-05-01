package helpers

import (
	"fmt"
	"unicode"
)

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}

	return true
}

func Print(biodata []Biodata, absen int) {
	fmt.Println("Nama: ", biodata[absen].Nama)
	fmt.Println("Alamat: ", biodata[absen].Alamat)
	fmt.Println("Pekerjaan: ", biodata[absen].Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang: ", biodata[absen].Alasan)
}
