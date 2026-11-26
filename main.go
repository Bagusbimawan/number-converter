package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("=== Konversi Bilangan ===")
	fmt.Println("0b = Biner, 0o = Oktal, 0x = Hex")
	fmt.Println("Tanpa prefix = Desimal")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nMasukkan angka (q untuk keluar): ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(strings.ToLower(scanner.Text()))

		if input == "q" {
			break
		}

		negatif := ""
		if strings.HasPrefix(input, "-") {
			negatif = "-"
			input = input[1:]
		}

		basis := 10
		jenis := "Desimal"
		valid := "0123456789"

		switch {
		case strings.HasPrefix(input, "0b"):
			basis = 2
			jenis = "Biner"
			valid = "01"
			input = input[2:]

		case strings.HasPrefix(input, "0o"):
			basis = 8
			jenis = "Oktal"
			valid = "01234567"
			input = input[2:]

		case strings.HasPrefix(input, "0x"):
			basis = 16
			jenis = "Hex"
			valid = "0123456789abcdef"
			input = input[2:]
		}

		if input == "" {
			fmt.Println("Input tidak valid.")
			continue
		}

		validInput := true

		for _, c := range input {
			if !strings.ContainsRune(valid, c) {
				validInput = false
				break
			}
		}

		if !validInput {
			fmt.Println("Input tidak valid.")
			continue
		}

		angka, err := strconv.ParseInt(negatif+input, basis, 64)

		if err != nil {
			fmt.Println("Angka tidak valid atau terlalu besar.")
			continue
		}

		fmt.Println("Jenis:", jenis)
		fmt.Printf("Desimal      : %d\n", angka)
		fmt.Printf("Biner        : %b\n", angka)
		fmt.Printf("Oktal        : %o\n", angka)
		fmt.Printf("Heksadesimal : %X\n", angka)
	}
}
