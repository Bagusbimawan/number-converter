package main

import "fmt"

func main() {
	fmt.Println("=== Konversi Bilangan ===")
	fmt.Println("0b = Biner, 0o = Oktal, 0x = Hex")
	fmt.Println("Tanpa prefix = Desimal")
	fmt.Println("Bilangan negatif integer menggunakan 8-bit Two's Complement")
	fmt.Println("Masukkan q untuk keluar")

	for {
		var input string

		fmt.Print("\nMasukkan angka: ")
		fmt.Scan(&input)

		if input == "q" || input == "Q" {
			break
		}

		angka, jenis, valid := parseInput(input)

		if !valid {
			fmt.Println("Input tidak valid atau di luar range.")
			continue
		}

		fmt.Println("Jenis:", jenis)
		fmt.Printf("Desimal      : %g\n", angka)
		fmt.Printf("Biner        : %s\n", toBinary(angka))
		fmt.Printf("Oktal        : %s\n", toBaseInteger(angka, 8))
		fmt.Printf("Heksadesimal : %s\n", toBaseInteger(angka, 16))
	}
}

func parseInput(input string) (float64, string, bool) {
	negatif := false

	if len(input) == 0 {
		return 0, "", false
	}

	if input[0] == '-' {
		negatif = true
		input = input[1:]
	}

	if len(input) == 0 {
		return 0, "", false
	}

	basis := 10
	jenis := "Desimal"

	switch {
	case len(input) >= 2 && input[0:2] == "0b":
		basis = 2
		jenis = "Biner"
		input = input[2:]

	case len(input) >= 2 && input[0:2] == "0o":
		basis = 8
		jenis = "Oktal"
		input = input[2:]

	case len(input) >= 2 && input[0:2] == "0x":
		basis = 16
		jenis = "Hex"
		input = input[2:]
	}

	if len(input) == 0 {
		return 0, "", false
	}

	hasil := 0.0
	faktor := 1.0
	koma := false
	jumlahKoma := 0

	for i := 0; i < len(input); i++ {
		if input[i] == '.' {
			if koma || basis != 10 {
				return 0, "", false
			}

			koma = true
			continue
		}

		digit := getDigit(input[i])

		if digit < 0 || digit >= basis {
			return 0, "", false
		}

		if !koma {
			hasil = hasil*float64(basis) + float64(digit)
		} else {
			faktor /= float64(basis)
			hasil += float64(digit) * faktor
			jumlahKoma++
		}
	}

	if jumlahKoma > 10 {
		return 0, "", false
	}

	if negatif {
		hasil = -hasil
	}

	if !koma && (hasil < -128 || hasil > 127) {
		return 0, "", false
	}

	return hasil, jenis, true
}

func getDigit(c byte) int {
	switch {
	case c >= '0' && c <= '9':
		return int(c - '0')

	case c >= 'a' && c <= 'f':
		return int(c-'a') + 10

	case c >= 'A' && c <= 'F':
		return int(c-'A') + 10

	default:
		return -1
	}
}

func toBinary(angka float64) string {

	if angka == float64(int64(angka)) {
		nilai := int64(angka)

		if nilai < -128 || nilai > 127 {
			return "Di luar range 8-bit"
		}

		bits := uint8(nilai)
		hasil := ""

		for i := 7; i >= 0; i-- {
			if bits&(1<<i) != 0 {
				hasil += "1"
			} else {
				hasil += "0"
			}
		}

		return hasil
	}

	negatif := angka < 0

	if negatif {
		angka = -angka
	}

	bagianInteger := int64(angka)
	bagianPecahan := angka - float64(bagianInteger)

	hasil := toPositiveBinary(bagianInteger)
	hasilPecahan := ""

	for i := 0; i < 10 && bagianPecahan > 0; i++ {
		bagianPecahan *= 2

		if bagianPecahan >= 1 {
			hasilPecahan += "1"
			bagianPecahan -= 1
		} else {
			hasilPecahan += "0"
		}
	}

	if hasilPecahan != "" {
		hasil += "." + hasilPecahan
	}

	if negatif {
		hasil = "-" + hasil
	}

	return hasil
}

func toPositiveBinary(angka int64) string {
	if angka == 0 {
		return "0"
	}

	hasil := ""

	for angka > 0 {
		sisa := angka % 2
		hasil = string('0'+sisa) + hasil
		angka /= 2
	}

	return hasil
}

func toBaseInteger(angka float64, basis int) string {

	if angka != float64(int64(angka)) {
		return "Integer saja"
	}

	nilai := int64(angka)

	if nilai < -128 || nilai > 127 {
		return "Di luar range 8-bit"
	}

	bits := uint8(nilai)
	digit := "0123456789ABCDEF"
	hasil := ""

	for bits > 0 {
		sisa := bits % uint8(basis)
		hasil = string(digit[sisa]) + hasil
		bits /= uint8(basis)
	}

	if hasil == "" {
		hasil = "0"
	}

	if basis == 8 {
		for len(hasil) < 3 {
			hasil = "0" + hasil
		}
	}

	if basis == 16 {
		for len(hasil) < 2 {
			hasil = "0" + hasil
		}
	}

	return hasil
}
