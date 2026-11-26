package main

import "fmt"

func main() {
	fmt.Println("=== Konversi Bilangan ===")
	fmt.Println("0b = Biner, 0o = Oktal, 0x = Hex")
	fmt.Println("Tanpa prefix = Desimal")

	for {
		var input string

		fmt.Print("\nMasukkan angka (q untuk keluar): ")
		fmt.Scan(&input)

		if input == "q" || input == "Q" {
			break
		}

		angka, jenis, valid := toDecimal(input)

		if !valid {
			fmt.Println("Input tidak valid.")
			continue
		}

		fmt.Println("Jenis:", jenis)
		fmt.Printf("Desimal      : %d\n", angka)
		fmt.Printf("Biner        : %s\n", toBase(angka, 2))
		fmt.Printf("Oktal        : %s\n", toBase(angka, 8))
		fmt.Printf("Heksadesimal : %s\n", toBase(angka, 16))
	}
}

func toDecimal(input string) (int64, string, bool) {
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

	var hasil uint64

	for i := 0; i < len(input); i++ {
		digit := getDigit(input[i])

		if digit < 0 || digit >= basis {
			return 0, "", false
		}

		hasil = hasil*uint64(basis) + uint64(digit)
	}

	if negatif {
		if hasil > 9223372036854775808 {
			return 0, "", false
		}

		if hasil == 9223372036854775808 {
			return -9223372036854775808, jenis, true
		}

		return -int64(hasil), jenis, true
	}

	if hasil > 9223372036854775807 {
		return 0, "", false
	}

	return int64(hasil), jenis, true
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

func toBase(angka int64, basis int) string {
	if angka == 0 {
		return "0"
	}

	negatif := angka < 0
	var nilai uint64

	if negatif {
		nilai = uint64(-(angka + 1)) + 1
	} else {
		nilai = uint64(angka)
	}

	digit := "0123456789ABCDEF"
	hasil := ""

	for nilai > 0 {
		sisa := nilai % uint64(basis)
		hasil = string(digit[sisa]) + hasil
		nilai /= uint64(basis)
	}

	if negatif {
		hasil = "-" + hasil
	}

	return hasil
}
