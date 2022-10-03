package main

import (
	"fmt"
)

func main() {
	var angka int
	fmt.Println("Ketik angka : ")
	fmt.Scanln(&angka)

	fmt.Println(Terbilang(angka))
}

var arrSatuan = [...]string{"",
	"satu",
	"dua",
	"tiga",
	"empat",
	"lima",
	"enam",
	"tujuh",
	"delapan",
	"sembilan",
	"sepuluh",
	"sebelas",
}

func Terbilang(angka int) string {
	// var err error
	if angka <= 11 {
		return arrSatuan[angka]
	} else if angka > 11 && angka <= 19 {
		return Terbilang(angka%10) + " Belas "
	} else if angka >= 20 && angka <= 99 {
		return Terbilang(angka/10) + " Puluh " + Terbilang(angka%10)
	} else if angka >= 100 && angka <= 199 {
		return "Seratus" + Terbilang(angka%100)
	} else if angka >= 200 && angka <= 999 {
		return Terbilang(angka/100) + " Ratus " + Terbilang(angka%100)
	} else if angka >= 1000 && angka <= 1999 {
		return "Seribu" + Terbilang(angka%1000)
	} else if angka >= 2000 && angka <= 9999 {
		return Terbilang(angka/1000) + " Ribu " + Terbilang(angka%1000)
	} else if angka >= 10000 && angka <= 99999 {
		return Terbilang(angka/1000) + " Ribu " + Terbilang(angka%1000)
	} else if angka >= 100000 && angka <= 999999 {
		return Terbilang(angka/1000) + " Ribu " + Terbilang(angka%1000)
	} else {
		return "Angka melebihi dari 6 digit"
	}

}
